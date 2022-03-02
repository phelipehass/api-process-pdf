package main

import (
	"api/config"
	"api/extract_data_from_pdf/delivery/api"
	"api/extract_data_from_pdf/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//TODO adicionar middleware
	config.LoadEnv()
	app := fiber.New()
	extractService := service.NewService()
	api.Handlers(app, extractService)
	app.Listen(":" + config.ApiPort())
}
