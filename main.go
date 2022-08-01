package main

import (
	"api/config"
	"api/extract_data_from_pdf/delivery/api"
	"api/extract_data_from_pdf/repository"
	"api/extract_data_from_pdf/service"
	"database/sql"
	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//TODO adicionar middleware
	config.LoadEnv()
	db := config.InitPostgres()
	redisClient := config.InitRedis()
	var name = "sss"

	defer db.Close()
	defer redisClient.Close()

	app := fiber.New()
	extractService := configService(db)
	log.Infof(name)

	api.Handlers(app, extractService)
	app.Listen(":" + config.ApiPort())
}

func configService(db *sql.DB) (extractService *service.ExtractService) {
	indicationRepo := repository.NewPostgresRepository(db)
	extractService = service.NewService(indicationRepo)
	return
}
