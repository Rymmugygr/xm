package v1

import (
	"github.com/Rymmugygr/xm/config"
	"github.com/Rymmugygr/xm/internal/app/usecase"
	"github.com/go-playground/validator/v10"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func errorHandler(c fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized or invalid token",
	})
}

func NewRoutes(app *fiber.App, apiV1Group fiber.Router, logger *logrus.Logger, cfg *config.Config, useCases usecase.UseCase) {
	r := &V1{
		uc: useCases,
		l:  logger,
		v:  validator.New(validator.WithRequiredStructEnabled()),
	}

	// Public
	authGroup := apiV1Group.Group("/auth")
	{
		authGroup.Get("/token", r.getToken)
	}

	companyGroup := apiV1Group.Group("/company")
	{
		companyGroup.Get("/:id", r.getCompany)
	}

	// TODO
	// The REAL authorization process should be here
	// This is a fake, for demonstration only
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: cfg.Jwt.Secret},
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
	}))

	// Private
	companyGroupPrivate := apiV1Group.Group("/company")
	{
		companyGroupPrivate.Post("/", r.createCompany)
		companyGroupPrivate.Patch("/:id", r.pathCompany)
		companyGroupPrivate.Delete("/:id", r.deleteCompany)
	}
}
