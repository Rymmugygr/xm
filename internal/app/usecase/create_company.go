package usecase

import (
	"context"
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
		Type              entity.CompanyType
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
		input.Type,
	)
	if err != nil {
		return CreateCompanyOutput{}, err
	}

	err = uc.repository.Insert(ctx, company)
	if err != nil {
		return CreateCompanyOutput{}, err
	}

	return CreateCompanyOutput{
		Company: company,
	}, nil
}
