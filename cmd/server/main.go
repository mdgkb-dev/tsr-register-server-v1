package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mdgkb/tsr-tegister-server-v1/migrations"

	_ "github.com/go-pg/pg/v10/orm"
	"github.com/pro-assistance/pro-assister/config"
	helperPack "github.com/pro-assistance/pro-assister/helper"
	"log"
	"mdgkb/tsr-tegister-server-v1/routing"
	"net/http"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	router := gin.Default()
	helper := helperPack.NewHelper(*conf)
	db := helper.DB.InitDB()
	defer db.Close()
	helper.DB.DB = db
	helper.Init(migrations.Migrations)
	if helper.MigrateMode() {
		return
	}
	routing.Init(router, db, helper)
	err = http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), router)
	if err != nil {
		log.Fatalln(err)
	}
}
