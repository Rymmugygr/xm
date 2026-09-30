package db

import (
	"context"
	"errors"
	"fmt"
	"time"
	"uuid"

	"github.com/Rymmugygr/xm/internal/app/entity"
	"github.com/Rymmugygr/xm/internal/app/repository"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

type Company struct {
	Id                uuid.UUID `gorm:"column:id;primaryKey;not null;<-:false"`
	Name              string    `gorm:"column:name;unique;not null"`
	Description       string    `gorm:"column:description;not null;default:''"`
	AmountOfEmployees int       `gorm:"column:employees;not null"`
	Registered        bool      `gorm:"column:registered;not null"`
	Type              string    `gorm:"column:type;not null"`

	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (c *CompanyRepository) TableName() string {
	return "companies"
}

var companyTypeMap = map[entity.CompanyType]string{
	entity.CompanyTypeCorporations:       "corporations",
	entity.CompanyTypeNonProfit:          "non_profit",
	entity.CompanyTypeCooperative:        "cooperative",
	entity.CompanyTypeSoleProprietorship: "sole_proprietorship",
}

func companyTypeToModel(companyType entity.CompanyType) string {
	if v, ok := companyTypeMap[companyType]; ok {
		return v
	}
	return "unknown"
}

func companyTypeToEntity(companyType string) entity.CompanyType {
	for k, v := range companyTypeMap {
		if v == companyType {
			return k
		}
	}
	return entity.CompanyTypeWrong
}

func NewCompanyRepository(gorm *gorm.DB) *CompanyRepository {
	return &CompanyRepository{
		db: gorm,
	}
}

func (c *CompanyRepository) FindByID(ctx context.Context, id entity.CompanyId) (*entity.Company, error) {
	var company Company

	err := c.db.WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", id).
		Take(&company).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, fmt.Errorf("company repository (postgres): could not find Company: %w", err)
	}

	return &entity.Company{
		ID:                entity.CompanyId(company.Id),
		Name:              company.Name,
		Description:       company.Description,
		AmountOfEmployees: company.AmountOfEmployees,
		Registered:        company.Registered,
		Type:              companyTypeToEntity(company.Type),
	}, nil
}

func (c *CompanyRepository) Insert(ctx context.Context, company *entity.Company) error {
	row := Company{
		Name:              company.Name,
		Description:       company.Description,
		AmountOfEmployees: company.AmountOfEmployees,
		Registered:        company.Registered,
		Type:              companyTypeToModel(company.Type),
	}

	err := c.db.WithContext(ctx).
		Create(&row).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return repository.ErrConflict
		}
		return fmt.Errorf("company repository (postgres): could not insert Company: %w", err)
	}

	return nil
}

func (c *CompanyRepository) Update(ctx context.Context, partial *entity.PartialCompany) (updated *entity.Company, err error) {
	err = c.db.Transaction(func(tx *gorm.DB) error {
		company, err := c.FindByID(ctx, partial.ID)
		if err != nil {
			return err
		}

		fields := map[string]any{
			"name":        "",
			"description": "",
			"employees":   0,
			"registered":  false,
			"type":        "",
		}

		if partial.Name != nil {
			fields["name"] = *partial.Name
		} else {
			fields["name"] = company.Name
		}
		if partial.Description != nil {
			fields["description"] = *partial.Description
		} else {
			fields["description"] = company.Description
		}
		if partial.AmountOfEmployees != nil {
			fields["employees"] = *partial.AmountOfEmployees
		} else {
			fields["employees"] = company.AmountOfEmployees
		}
		if partial.Registered != nil {
			fields["registered"] = *partial.Registered
		} else {
			fields["registered"] = company.Registered
		}
		if partial.Type != nil {
			fields["type"] = companyTypeToModel(*partial.Type)
		} else {
			fields["type"] = companyTypeToModel(company.Type)
		}

		fmt.Println(fields)
		err = c.db.WithContext(ctx).Model(&Company{}).
			Debug().
			Where("id = ? AND deleted_at IS NULL", company.ID).
			Updates(fields).
			Error

		updated, _ = c.FindByID(ctx, partial.ID)

		return err
	})

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, repository.ErrConflict
		}
		return nil, fmt.Errorf("company repository (postgres): could not update Company with ID %v: %w", partial.ID, err)
	}
	return updated, nil
}

func (c *CompanyRepository) Delete(ctx context.Context, id entity.CompanyId) error {
	err := c.db.WithContext(ctx).Model(&Company{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", gorm.Expr("now()")).
		Error

	return err
}
