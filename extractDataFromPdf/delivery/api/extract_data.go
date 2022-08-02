package api

import (
	"api/extractDataFromPdf/service"
	diaryService "api/getDiariesFromCouncil/service"
	"api/job"
	"github.com/bamzi/jobrunner"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type ExtractData struct {
	serviceExtract *service.ExtractService
	diariesService *diaryService.GetDiariesFromCouncilService
	redisClient    *redis.Client
}

func NewExtractData() (extractData *ExtractData) {
	return &ExtractData{}
}

func (es *ExtractData) AddServiceExtract(serviceExtract *service.ExtractService) {
	es.serviceExtract = serviceExtract
}

func (es *ExtractData) AddDiariesService(diariesService *diaryService.GetDiariesFromCouncilService) {
	es.diariesService = diariesService
}

func (es *ExtractData) AddRedisClient(redisClient *redis.Client) {
	es.redisClient = redisClient
}

func (es *ExtractData) Extract(ctx *fiber.Ctx) error {
	fileBuffer, err := ctx.FormFile("document")
	if err != nil {
		return err
	}

	err = es.serviceExtract.ExtractDataFromPDF(fileBuffer)

	return err
}

func (es *ExtractData) ProcessMoreDiaries(ctx *fiber.Ctx) error {
	initialDate := ctx.Params("initialDate")
	finalDate := ctx.Params("finalDate")

	jobrunner.Now(&job.GetDiaries{
		GetDiariesUrls:    es.diariesService,
		Redis:             es.redisClient,
		InitialDateFilter: initialDate,
		FinalDateFilter:   finalDate,
	})

	return nil
}

func (es *ExtractData) Handlers(app *fiber.App) {
	app.Post("/process/pdf", es.Extract)
	app.Post("/process/diaries/:initialDate/:finalDate", es.ProcessMoreDiaries)
}
