-- TABLE ingredient category:
insert into ingredient_category (category_name,created_at) values('',NOW());
insert into ingredient_category (category_name,created_at) values('chmiel',NOW());
insert into ingredient_category (category_name,created_at) values('drożdż',NOW());
-- TABLE ingredient
insert into ingredient (ingredient_name,unit,quantity,created_at,ingredient_category_id) values('Chmiel cytrusowy','kg',125,NOW(),2);
insert into ingredient (ingredient_name,unit,quantity,created_at,ingredient_category_id) values('Drożdże','kg',125,NOW(),3);
-- TABLE recipe category:
insert into recipe_category(name,created_at) values('IPA',NOW());
-- TABLE recipe:
insert into recipe(name,created_at,recipe_category_id) values('IPA #1',NOW(),1);
-- TABLE recipe ingredient list:
insert into recipe_ingredient_list (quantity,unit,created_at,recipe_id,ingredient_id) values(5,'kg',NOW(),1,2);
insert into recipe_ingredient_list (quantity,unit,created_at,recipe_id,ingredient_id) values(5,'kg',NOW(),1,3);
--TABLE mash tun:
insert into mash_stage(temperature,created_at,recipe_id,pump_work,stage_time) values(60,NOW(),1,true,3600000);
insert into mash_stage(temperature,created_at,recipe_id,pump_work,stage_time) values(60,NOW(),1,true,600000);