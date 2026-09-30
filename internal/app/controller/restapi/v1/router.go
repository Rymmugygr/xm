package v1

import (
	"github.com/Rymmugygr/xm/internal/app/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewRoutes(apiV1Group fiber.Router, logger *logrus.Logger, useCases usecase.UseCase) {
	r := &V1{
		uc: useCases,
		l:  logger,
		v:  validator.New(validator.WithRequiredStructEnabled()),
	}

	companyGroup := apiV1Group.Group("/company")
	{
		companyGroup.Post("/", r.createCompany)
		companyGroup.Get("/:id", r.getCompany)
		companyGroup.Patch("/:id", r.pathCompany)
		companyGroup.Delete("/:id", r.deleteCompany)
	}
}
