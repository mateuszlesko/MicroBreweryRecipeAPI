package microbreweryrecipeapi

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gorilla_hander "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/handlers"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/helpers"
)

func main() {
	ch := gorilla_hander.CORS(gorilla_hander.AllowedOrigins([]string{"http://localhost:3000"}))
	l := log.New(os.Stdout, "MicroBreweryRecipeAPI", log.LstdFlags)
	var err error
	err, pgsql := helpers.OpenConnection()
	if err != nil {
		log.Panic(err)
	}
	defer pgsql.Close()
	err = pgsql.Ping()
	if err != nil {
		panic(err)
	}
	rch := handlers.NewRecipeCategory(l)
	rh := handlers.NewRecipe(l)
	smux := mux.NewRouter()

	getRecipe := smux.Methods(http.MethodGet).Subrouter()
	putRecipe := smux.Methods(http.MethodPut).Subrouter()
	deleteRecipe := smux.Methods(http.MethodDelete).Subrouter()
	postRecipe := smux.Methods(http.MethodPost).Subrouter()

	getRecipeCategory := smux.Methods(http.MethodGet).Subrouter()
	putRecipeCategory := smux.Methods(http.MethodPut).Subrouter()
	deleteRecipeCategory := smux.Methods(http.MethodDelete).Subrouter()
	postRecipeCategory := smux.Methods(http.MethodPost).Subrouter()

	getRecipe.HandleFunc("/recipes", rh.GetRecipes)
	postRecipe.HandleFunc("/recipes", rh.PostRecipe)
	deleteRecipe.HandleFunc("/recipes", rh.DeleteRecipe)
	putRecipe.HandleFunc("/recipes", rh.PutRecipe)

	getRecipeCategory.HandleFunc("/recipes", rch.GetRecipeCategories)
	postRecipeCategory.HandleFunc("/recipes", rch.PostRecipeCategory)
	deleteRecipeCategory.HandleFunc("/recipes", rch.DeleteRecipeCategory)
	putRecipeCategory.HandleFunc("/recipes", rch.PutRecipeCategory)

	s := &http.Server{
		Addr:              ":9990",
		Handler:           ch(smux),
		TLSConfig:         &tls.Config{},
		ReadTimeout:       2 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    16,
		TLSNextProto:      map[string]func(*http.Server, *tls.Conn, http.Handler){},
		ConnState: func(net.Conn, http.ConnState) {
		},
		ErrorLog: l,
	}
	go func() {
		err := s.ListenAndServe()
		if err == nil {
			l.Fatal(err)
		}
	}()
	//serwer waits unitl everyone finish and does not take any new request, after that it will peacfully shut down.
	sChan := make(chan os.Signal, 1)
	signal.Notify(sChan, os.Interrupt)
	signal.Notify(sChan, syscall.SIGTERM)
	sig := <-sChan
	l.Println("Recieved terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 2*time.Second)
	s.Shutdown(tc)
}
