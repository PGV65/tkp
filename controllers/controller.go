package controllers

import (
	"encoding/json"
	"log"
	"tkp/services"

	"github.com/syndicatedb/vodka"
)

type Controller interface {
	GetTKP(ctx *vodka.Context) (interface{}, error)
}
type controller struct {
	service services.Service
}

// New - controller constructor
func New(s services.Service) Controller {
	return &controller{
		service: s,
	}
}

func (c *controller) GetTKP(ctx *vodka.Context) (interface{}, error) {
	data := make(map[string]string)
	log.Printf("%+v", ctx)
	ctx.Request.FormFile("path")
	err := json.Unmarshal(ctx.Raw.Body, &data)
	if err != nil {
		return nil, err
	}

	return c.service.GetTKP(data["path"])
}
