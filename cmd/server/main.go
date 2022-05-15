package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/go-pg/pg/v10/orm"
	"github.com/pro-assistance/pro-assister/config"
	helperPack "github.com/pro-assistance/pro-assister/helper"
	"log"
	"mdgkb/tsr-tegister-server-v1/database/connect"
	"mdgkb/tsr-tegister-server-v1/routing"
	"net/http"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	router := gin.Default()

	db := connect.InitDB(conf.DB)
	defer db.Close()
	redis := connect.InitRedis(conf)
	helper := helperPack.NewHelper(*conf)
	routing.Init(router, db, redis, helper)

	err = http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), router)
	if err != nil {
		log.Fatalln(err)
	}
}
