package api

import (
	"api/extract_data_from_pdf/service"
	"github.com/gofiber/fiber/v2"
)

var serviceExtract *service.ExtractService

func Extract(ctx *fiber.Ctx) error {
	body := ctx.Body()
	bodyString := string(body)
	bodyString = "diario1.pdf"

	err := serviceExtract.ExtractDataFromPDF(bodyString)
	if err != nil {
		return err
	}

	return nil
}

func Handlers(app *fiber.App, extractService *service.ExtractService) {
	serviceExtract = extractService
	app.Get("/process/pdf", Extract)
}
