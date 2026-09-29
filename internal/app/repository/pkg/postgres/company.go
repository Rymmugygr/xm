package postgres

import (
	"context"
	"fmt"
	"time"
	"uuid"

	"github.com/Rymmugygr/xm/internal/app/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (c *CompanyRepository) FindByID(ctx context.Context, id entity.CompanyId) (*entity.Company, error) {
	var company Company

	err := c.db.WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", id).
		Take(&company).
		Error

	if err != nil {
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
		return fmt.Errorf("company repository (postgres): could not create Company: %w", err)
	}
	return nil
}

func (c *CompanyRepository) Update(ctx context.Context, company *entity.Company) error {
	row := Company{
		Name:              company.Name,
		Description:       company.Description,
		AmountOfEmployees: company.AmountOfEmployees,
		Registered:        company.Registered,
		Type:              companyTypeToModel(company.Type),
	}

	err := c.db.WithContext(ctx).Model(&Company{}).
		Clauses(clause.Returning{}).
		Where("id = ? AND deleted_at IS NULL", company.ID).
		Updates(&row).
		Error

	if err != nil {
		return fmt.Errorf("company repository (postgres): could not update Company with ID %v: %w", company.ID, err)
	}
	return nil
}

func (c *CompanyRepository) Delete(ctx context.Context, id entity.CompanyId) error {
	rows := c.db.WithContext(ctx).Model(&Company{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Update("deleted_at", gorm.Expr("now()")).
		RowsAffected

	if rows == 0 {
		return fmt.Errorf("company repository (postgres): could not delete Company with ID %v", id)
	}
	return nil
}
