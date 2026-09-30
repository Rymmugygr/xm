package v1

import (
	"github.com/Rymmugygr/xm/internal/app/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type V1 struct {
	uc usecase.UseCase
	l  *logrus.Logger
	v  *validator.Validate
}
