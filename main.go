package main

import (
	"api/config"
	"api/extractDataFromPdf/delivery/api"
	"api/extractDataFromPdf/repository"
	"api/extractDataFromPdf/service"
	diaryRepository "api/getDiariesFromCouncil/repository"
	diaryService "api/getDiariesFromCouncil/service"
	"api/job"
	"database/sql"
	"github.com/bamzi/jobrunner"
	"github.com/gofiber/fiber/v2"
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

	//TODOS OS DIAS 23h
	jobrunner.Schedule("TZ=America/Sao_Paulo 0 23 * * *", &job.GetDiaries{
		GetDiariesUrls: diariesService,
		Redis:          redisClient,
	})

	extractData := api.NewExtractData()

	extractData.AddServiceExtract(extractService)
	extractData.AddRedisClient(redisClient)
	extractData.AddDiariesService(diariesService)
	extractData.Handlers(app)
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
