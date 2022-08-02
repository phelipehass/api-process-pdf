package main

import (
	"api/config"
	"api/extract_data_from_pdf/delivery/api"
	"api/extract_data_from_pdf/repository"
	"api/extract_data_from_pdf/service"
	diaryRepository "api/getDiariesFromCouncil/repository"
	diaryService "api/getDiariesFromCouncil/service"
	"api/job"
	"database/sql"
	"github.com/bamzi/jobrunner"
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	//TODO adicionar middleware
	config.LoadEnv()
	db := config.InitPostgres()
	redisClient := config.InitRedis()

	defer db.Close()
	defer redisClient.Close()

	app := fiber.New()
	extractService, diariesService := configService(db)

	jobrunner.Start()
	jobrunner.Now(&job.GetDiaries{
		GetDiariesUrls: diariesService,
		Redis:          redisClient,
	})
	jobrunner.Every(10*time.Hour, &job.GetDiaries{
		GetDiariesUrls: diariesService,
		Redis:          redisClient,
	})

	api.Handlers(app, extractService)
	app.Listen(":" + config.ApiPort())

	jobrunner.Stop()
}

func configService(db *sql.DB) (extractService *service.ExtractService, diariesService *diaryService.GetDiariesFromCouncilService) {
	restRepo := diaryRepository.NewRepositoryRest(config.GetURLBaseConsult(), config.GetCookie())
	indicationRepo := repository.NewPostgresRepository(db)
	diaryRepo := diaryRepository.NewPostgresRepository(db)
	extractService = service.NewService(indicationRepo)
	diariesService = diaryService.NewService(diaryRepo, restRepo)
	return
}
