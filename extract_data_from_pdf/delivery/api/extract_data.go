package api

import (
	"api/extract_data_from_pdf/service"
	"github.com/gofiber/fiber/v2"
)

var serviceExtract *service.ExtractService

func Extract(ctx *fiber.Ctx) error {
	fileBuffer, err := ctx.FormFile("document")
	if err != nil {
		return err
	}

	err = serviceExtract.ExtractDataFromPDF(fileBuffer)

	return err
}

func Handlers(app *fiber.App, extractService *service.ExtractService) {
	serviceExtract = extractService
	app.Post("/process/pdf", Extract)
}
