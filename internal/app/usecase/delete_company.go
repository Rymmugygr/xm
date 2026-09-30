package usecase

import (
	"context"
	"uuid"

	"github.com/Rymmugygr/xm/internal/app/entity"
	"github.com/Rymmugygr/xm/internal/app/repository"
)

type DeleteCompanyUseCase interface {
	Execute(ctx context.Context, input DeleteCompanyInput) (DeleteCompanyOutput, error)
}

type (
	DeleteCompanyInput struct {
		ID uuid.UUID
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
	err := uc.repository.Delete(ctx, entity.CompanyId(input.ID))
	if err != nil {
		return DeleteCompanyOutput{}, err
	}

	return DeleteCompanyOutput{
		Success: true,
	}, nil
}
