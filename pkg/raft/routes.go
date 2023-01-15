package raft

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func AppendEntryRoutes(app fiber.Router, service Service) {
	app.Post("/append", func(ctx *fiber.Ctx) error {
		var entry AppendEntry
		err := ctx.BodyParser(&entry)

		service.Acknowledge(entry)

		if err != nil {
			log.Error(err)
			ctx.Status(http.StatusBadRequest)
			return ctx.SendString("Ooops")
		}

		ctx.Status(http.StatusOK)
		return ctx.SendString("Acknowledged.")
	})
}
