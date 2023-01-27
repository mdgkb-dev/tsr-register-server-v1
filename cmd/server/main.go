package main

import (
	"log"
	"mdgkb/tsr-tegister-server-v1/migrations"
	"mdgkb/tsr-tegister-server-v1/routing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/pro-assistance/pro-assister/config"
	helperPack "github.com/pro-assistance/pro-assister/helper"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	helper := helperPack.NewHelper(*conf)

	router := gin.New()
	router.Use(gin.Recovery())
	routing.Init(router, helper)
	helper.Run(migrations.Init(), router)
}
