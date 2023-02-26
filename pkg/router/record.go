package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/middleware"
)

func RegisterRecord(
	router fiber.Router,
	wrapper schema.ServerInterfaceWrapper,
) {
	records := router.Group("/domains/:domain/records", middleware.Domain)
	record := records.Group("/:record")
	records.Get("", wrapper.ListRecordsInDomain)
	record.Get("", wrapper.GetRecord)
}
