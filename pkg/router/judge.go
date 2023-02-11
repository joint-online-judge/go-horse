package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
)

func RegisterJudge(
	router fiber.Router,
	wrapper schema.ServerInterfaceWrapper,
) {
	record := router.Group("/domains/:domain/records/:record")
	record.Put("/cases/:index/judge", wrapper.SubmitCaseByJudger)
	record.Put("/judge", wrapper.SubmitRecordByJudger)
	record.Post("/judge/claim", wrapper.ClaimRecordByJudger)
}
