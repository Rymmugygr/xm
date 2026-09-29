package repository

import (
	"context"

	"github.com/Rymmugygr/xm/internal/app/entity"
)

type CompanyRepository interface {
	FindByID(ctx context.Context, id entity.CompanyId) (*entity.Company, error)
	Insert(ctx context.Context, company *entity.Company) error
	Update(ctx context.Context, company *entity.Company) error
	Delete(ctx context.Context, id entity.CompanyId) error
}
