package usecase

import (
	"context"
	"uuid"

	"github.com/Rymmugygr/xm/internal/app/entity"
	"github.com/Rymmugygr/xm/internal/app/repository"
)

type ModifyCompanyUseCase interface {
	Execute(ctx context.Context, input ModifyCompanyInput) (ModifyCompanyOutput, error)
}

type (
	ModifyCompanyInput struct {
		ID                uuid.UUID
		Name              *string
		Description       *string
		AmountOfEmployees *int
		Registered        *bool
		Type              *string
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
	company := entity.PartialCompany{
		ID:                entity.CompanyId(input.ID),
		Name:              input.Name,
		Description:       input.Description,
		AmountOfEmployees: input.AmountOfEmployees,
		Registered:        input.Registered,
	}
	if input.Type != nil {
		company.Type = new(entity.CompanyTypeFromString(*input.Type))
	}

	updated, err := uc.repository.Update(ctx, &company)
	if err != nil {
		return ModifyCompanyOutput{}, err
	}

	return ModifyCompanyOutput{
		Company: updated,
	}, nil
}
