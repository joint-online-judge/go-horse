package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

func RegisterRecord(
	router fiber.Router,
	wrapper schemas.ServerInterfaceWrapper,
) {
	records := router.Group("/domains/:domain/records")
	record := records.Group("/:record")
	records.Get("", wrapper.ListRecordsInDomain)
	record.Get("", wrapper.GetRecord)
}
