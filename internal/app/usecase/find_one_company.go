package usecase

import (
	"context"
	"uuid"

	"github.com/Rymmugygr/xm/internal/app/entity"
	"github.com/Rymmugygr/xm/internal/app/repository"
)

type FindOneCompanyUseCase interface {
	Execute(ctx context.Context, input FindOneCompanyInput) (FindOneCompanyOutput, error)
}

type (
	FindOneCompanyInput struct {
		ID uuid.UUID
	}

	FindOneCompanyOutput struct {
		Company *entity.Company
	}

	FindOneCompanyUseCaseImpl struct {
		repository repository.CompanyRepository
	}
)

func NewFindOneCompanyUseCaseImpl(repository repository.CompanyRepository) FindOneCompanyUseCaseImpl {
	return FindOneCompanyUseCaseImpl{
		repository: repository,
	}
}

func (uc FindOneCompanyUseCaseImpl) Execute(ctx context.Context, input FindOneCompanyInput) (FindOneCompanyOutput, error) {
	var company *entity.Company

	company, err := uc.repository.FindByID(ctx, entity.CompanyId(input.ID))
	if err != nil {
		return FindOneCompanyOutput{}, err
	}

	return FindOneCompanyOutput{
		Company: company,
	}, nil

}
