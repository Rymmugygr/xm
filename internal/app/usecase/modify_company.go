package usecase

import (
	"context"

	"github.com/Rymmugygr/xm/internal/app/entity"
	"github.com/Rymmugygr/xm/internal/app/repository"
)

type ModifyCompanyUseCase interface {
	Execute(ctx context.Context, input ModifyCompanyInput) (ModifyCompanyOutput, error)
}

type (
	ModifyCompanyInput struct {
		ID                entity.CompanyId
		Name              string
		Description       string
		AmountOfEmployees int
		Registered        bool
		Type              entity.CompanyType
	}

	ModifyCompanyOutput struct {
		Company *entity.Company
	}

	ModifyCompanyUseCaseImpl struct {
		repository repository.CompanyRepository
	}
)

func NewModifyCompanyUseCaseImpl(repository repository.CompanyRepository) ModifyCompanyUseCaseImpl {
	return ModifyCompanyUseCaseImpl{
		repository: repository,
	}
}

func (uc ModifyCompanyUseCaseImpl) Execute(ctx context.Context, input ModifyCompanyInput) (ModifyCompanyOutput, error) {
	company := entity.Company{
		ID:                input.ID,
		Name:              input.Name,
		Description:       input.Description,
		AmountOfEmployees: input.AmountOfEmployees,
		Registered:        input.Registered,
		Type:              input.Type,
	}

	err := uc.repository.Update(ctx, &company)
	if err != nil {
		return ModifyCompanyOutput{}, err
	}

	return ModifyCompanyOutput{
		Company: &company,
	}, nil
}
