package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/Rymmugygr/xm/internal/app/entity"
	"github.com/Rymmugygr/xm/internal/app/repository"
)

type CreateCompanyUseCase interface {
	Execute(ctx context.Context, input CreateCompanyInput) (CreateCompanyOutput, error)
}

type (
	CreateCompanyInput struct {
		Name              string
		Description       string
		AmountOfEmployees int
		Registered        bool
		Type              string
	}

	CreateCompanyOutput struct {
		Company *entity.Company
	}

	CreateCompanyUseCaseImpl struct {
		repository repository.CompanyRepository
	}
)

func NewCreateCompanyUseCaseImpl(repository repository.CompanyRepository) CreateCompanyUseCaseImpl {
	return CreateCompanyUseCaseImpl{
		repository: repository,
	}
}

func (uc CreateCompanyUseCaseImpl) Execute(ctx context.Context, input CreateCompanyInput) (CreateCompanyOutput, error) {
	company, err := entity.NewCompany(
		strings.TrimSpace(input.Name),
		strings.TrimSpace(input.Description),
		input.AmountOfEmployees,
		input.Registered,
		entity.CompanyTypeFromString(input.Type),
	)
	if err != nil {
		return CreateCompanyOutput{}, err
	}

	err = uc.repository.Insert(ctx, company)
	if err != nil {
		if errors.Is(err, repository.ErrConflict) {
			return CreateCompanyOutput{}, err
		}
		return CreateCompanyOutput{}, err
	}

	return CreateCompanyOutput{
		Company: company,
	}, nil
}
