package usecase

import (
	"context"

	"github.com/Rymmugygr/xm/internal/app/entity"
	"github.com/Rymmugygr/xm/internal/app/repository"
)

type DeleteCompanyUseCase interface {
	Execute(ctx context.Context, input DeleteCompanyInput) (DeleteCompanyOutput, error)
}

type (
	DeleteCompanyInput struct {
		ID entity.CompanyId
	}

	DeleteCompanyOutput struct {
		Success bool
	}

	DeleteCompanyUseCaseImpl struct {
		repository repository.CompanyRepository
	}
)

func NewDeleteCompanyUseCaseImpl(repository repository.CompanyRepository) DeleteCompanyUseCaseImpl {
	return DeleteCompanyUseCaseImpl{
		repository: repository,
	}
}

func (uc DeleteCompanyUseCaseImpl) Execute(ctx context.Context, input DeleteCompanyInput) (DeleteCompanyOutput, error) {
	err := uc.repository.Delete(ctx, input.ID)
	if err != nil {
		return DeleteCompanyOutput{}, err
	}

	return DeleteCompanyOutput{
		Success: true,
	}, nil
}
