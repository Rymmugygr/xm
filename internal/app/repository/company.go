package repository

import (
	"context"
	"errors"

	"github.com/Rymmugygr/xm/internal/app/entity"
)

var ErrNotFound = errors.New("company not found")
var ErrConflict = errors.New("company already exists")

type CompanyRepository interface {
	FindByID(ctx context.Context, id entity.CompanyId) (*entity.Company, error)
	Insert(ctx context.Context, company *entity.Company) error
	Update(ctx context.Context, company *entity.PartialCompany) (updated *entity.Company, err error)
	Delete(ctx context.Context, id entity.CompanyId) error
}
