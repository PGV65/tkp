package main

import (
	"log"
	"tkp/controllers"
	"tkp/modules/sibur"
	"tkp/services"

	"github.com/syndicatedb/vodka"
)

var config Config

func init() {
	var err error
	config, err = NewConfig("./config.json")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// excelFileName := "./tkp.xlsx"
}
func main() {
	engine := vodka.New()
	engine.Server(config.HTTPConfig)
	engine.Validation("./validation.json")

	sib := sibur.New()
	srv := services.New(sib)
	ctrl := controllers.New(srv)

	engine.Router.POST("/sibur", ctrl.GetTKP)

	for {
		engine.Start()
	}
}
