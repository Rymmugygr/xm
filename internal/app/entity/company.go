package entity

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
	"uuid"
)

var (
	ErrEmptyName                = errors.New("company entity: 'Name' is empty")
	ErrLongName                 = errors.New("company entity: 'Name' is too long")
	ErrLongDescription          = errors.New("company entity: 'Description' is too long")
	ErrInvalidAmountOfEmployees = errors.New("company entity: 'Amount of Employees' is negative or zero")
	ErrInvalidType              = errors.New("company entity: 'CompanyType' is invalid")
)

const (
	NameMaxLength        = 15
	DescriptionMaxLength = 3000
)

type CompanyType int

const (
	CompanyTypeWrong CompanyType = iota
	CompanyTypeCorporations
	CompanyTypeNonProfit
	CompanyTypeCooperative
	CompanyTypeSoleProprietorship
)

var mapCompanyTypes = map[CompanyType]string{
	CompanyTypeCorporations:       "Corporations",
	CompanyTypeNonProfit:          "Non Profit",
	CompanyTypeCooperative:        "Cooperative",
	CompanyTypeSoleProprietorship: "Sole Proprietorship",
}

func (t CompanyType) Validate() error {
	switch t {
	case CompanyTypeCorporations, CompanyTypeNonProfit, CompanyTypeCooperative, CompanyTypeSoleProprietorship:
		return nil
	default:
		return ErrInvalidType
	}
}

func (t CompanyType) String() string {
	if val, ok := mapCompanyTypes[t]; ok {
		return val
	}
	return "Unknown"
}

func CompanyTypeFromString(companyType string) CompanyType {
	for k, v := range mapCompanyTypes {
		if v == companyType {
			return k
		}
	}
	return CompanyTypeWrong
}

type CompanyId uuid.UUID

func (id CompanyId) String() string {
	return uuid.UUID(id).String()
}

type Company struct {
	ID                CompanyId
	Type              CompanyType
	AmountOfEmployees int
	Registered        bool
	Name              string
	Description       string
}

type PartialCompany struct {
	ID                CompanyId
	Type              *CompanyType
	AmountOfEmployees *int
	Registered        *bool
	Name              *string
	Description       *string
}

func NewCompany(name, description string, amountOfEmployees int, registered bool, companyType CompanyType) (*Company, error) {
	company := &Company{
		ID:                CompanyId(uuid.NewV7()),
		Type:              companyType,
		AmountOfEmployees: amountOfEmployees,
		Registered:        registered,
		Name:              name,
		Description:       description,
	}

	err := company.Validate()
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (c Company) Validate() error {
	var errs []error

	if c.Name == "" {
		errs = append(errs, ErrEmptyName)
	}
	if utf8.RuneCountInString(c.Name) > NameMaxLength {
		errs = append(errs, ErrLongName)
	}

	if utf8.RuneCountInString(c.Description) > DescriptionMaxLength {
		errs = append(errs, ErrLongDescription)
	}

	if c.AmountOfEmployees <= 0 {
		errs = append(errs, ErrInvalidAmountOfEmployees)
	}

	if err := c.Type.Validate(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (c Company) String() string {
	var s strings.Builder

	s.WriteString("ID: ")
	s.WriteString(c.ID.String())
	s.WriteString("\nName: ")
	s.WriteString(c.Name)
	s.WriteString("\nDescription: ")
	s.WriteString(c.Description)
	s.WriteString("\nAmount of Employees: ")
	s.WriteString(fmt.Sprintf("%d", c.AmountOfEmployees))
	s.WriteString("\nRegistered: ")
	if c.Registered {
		s.WriteString("yes")
	} else {
		s.WriteString("no")
	}
	s.WriteString("\nType: ")
	s.WriteString(c.Type.String())

	return s.String()
}
