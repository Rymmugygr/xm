package restapi

import (
	"net/http"

	"github.com/Rymmugygr/xm/config"
	v1 "github.com/Rymmugygr/xm/internal/app/controller/restapi/v1"
	"github.com/Rymmugygr/xm/internal/app/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewRouter(logger *logrus.Logger, app *fiber.App, cfg *config.Config, useCases usecase.UseCase) {
	app.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })

	apiV1Group := app.Group("/v1")
	{
		v1.NewRoutes(app, apiV1Group, logger, cfg, useCases)
	}
}
