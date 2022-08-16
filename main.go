package main

import (
	"context"
	"crypto/tls"
	"fmt"
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
		pgsql.Close()
		log.Panic(err)
	}
	err = pgsql.Ping()
	if err != nil {
		pgsql.Close()
		panic(err)
	}
	pgsql.Close()

	rch := handlers.NewRecipeCategory(l)
	rh := handlers.NewRecipe(l)
	msh := handlers.NewMashStage(l)

	smux := mux.NewRouter()
	smux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("XD"))
	})
	getRecipe := smux.Methods(http.MethodGet).Subrouter()
	putRecipe := smux.Methods(http.MethodPut).Subrouter()
	deleteRecipe := smux.Methods(http.MethodDelete).Subrouter()
	postRecipe := smux.Methods(http.MethodPost).Subrouter()

	getRecipeCategory := smux.Methods(http.MethodGet).Subrouter()
	putRecipeCategory := smux.Methods(http.MethodPut).Subrouter()
	deleteRecipeCategory := smux.Methods(http.MethodDelete).Subrouter()
	postRecipeCategory := smux.Methods(http.MethodPost).Subrouter()

	getMashStage := smux.Methods(http.MethodGet).Subrouter()
	postMashStage := smux.Methods(http.MethodPost).Subrouter()
	putMashStage := smux.Methods(http.MethodPut).Subrouter()
	deleteMashStage := smux.Methods(http.MethodDelete).Subrouter()

	getRecipe.HandleFunc("/recipes/", rh.GetRecipes)
	getRecipe.HandleFunc("/recipes/details/", rh.GetRecipeById)
	postRecipe.HandleFunc("/recipes/", rh.PostRecipe)
	deleteRecipe.HandleFunc("/recipes/", rh.DeleteRecipe)
	putRecipe.HandleFunc("/recipes/", rh.PutRecipe)

	getRecipeCategory.HandleFunc("/recipecategories/", rch.GetRecipeCategories)
	postRecipeCategory.HandleFunc("/recipecategories/", rch.PostRecipeCategory)
	deleteRecipeCategory.HandleFunc("/recipecategories/{id:[0-9]+}", rch.DeleteRecipeCategory)
	putRecipeCategory.HandleFunc("/recipecategories/{id:[0-9]+}", rch.PutRecipeCategory)

	getMashStage.HandleFunc("/mashstages/", msh.GetMashStageByRecipeId)
	postMashStage.HandleFunc("/mashstages/", msh.PostMashStage)
	putMashStage.HandleFunc("/mashstages/{id:[0:9]+}", msh.UpdateMashStage)
	deleteMashStage.HandleFunc("/mashstages/{id:[0-9]+}", msh.DeleteMashStage)
	fmt.Println("Server is listening on :9990")
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
