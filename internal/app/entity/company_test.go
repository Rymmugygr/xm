package entity

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"testing"
	"uuid"
)

func TestCompany_Validate(t *testing.T) {
	tests := map[string]struct {
		expErr  error
		company Company
	}{
		"Empty Name": {
			ErrEmptyName,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "",
				AmountOfEmployees: 1,
				Type:              CompanyTypeNonProfit,
			},
		},
		"Long Name": {
			ErrLongName,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "abcdefghijklmnop",
				AmountOfEmployees: 1,
				Type:              CompanyTypeNonProfit,
			},
		},
		"Long Name (multibyte)": {
			ErrLongName,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "+" + strings.Repeat("🌍", 15),
				AmountOfEmployees: 1,
				Type:              CompanyTypeNonProfit,
			},
		},
		"Correct Name": {
			nil,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "abcdef",
				AmountOfEmployees: 1,
				Type:              CompanyTypeNonProfit,
			},
		},
		"Correct Name (multibyte)": {
			nil,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              strings.Repeat("🌍", 15),
				AmountOfEmployees: 1,
				Type:              CompanyTypeNonProfit,
			},
		},

		"Long Description": {
			ErrLongDescription,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				Description:       "+" + strings.Repeat("+", 3000),
				AmountOfEmployees: 1,
				Type:              CompanyTypeNonProfit,
			},
		},
		"Long Description (multibyte)": {
			ErrLongDescription,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				Description:       "+" + strings.Repeat("🌍", 3000),
				AmountOfEmployees: 1,
				Type:              CompanyTypeNonProfit,
			},
		},
		"Correct Description": {
			nil,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				Description:       "B",
				AmountOfEmployees: 1,
				Type:              CompanyTypeNonProfit,
			},
		},
		"Correct Description (empty)": {
			nil,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				Description:       "",
				AmountOfEmployees: 1,
				Type:              CompanyTypeNonProfit,
			},
		},
		"Correct Description (multibyte)": {
			nil,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				Description:       strings.Repeat("🌍", 3000),
				AmountOfEmployees: 1,
				Type:              CompanyTypeNonProfit,
			},
		},

		"Negative Employees": {
			ErrInvalidAmountOfEmployees,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				AmountOfEmployees: -1,
				Type:              CompanyTypeNonProfit,
			},
		},
		"Zero Employees": {
			ErrInvalidAmountOfEmployees,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				AmountOfEmployees: 0,
				Type:              CompanyTypeNonProfit,
			},
		},
		"Correct Employees": {
			nil,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				AmountOfEmployees: 1234,
				Type:              CompanyTypeNonProfit,
			},
		},
		"Correct Employees (max)": {
			nil,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				AmountOfEmployees: math.MaxInt,
				Type:              CompanyTypeNonProfit,
			},
		},

		"Empty CompanyType": {
			ErrInvalidType,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				AmountOfEmployees: 1,
			},
		},
		"Invalid CompanyType": {
			ErrInvalidType,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				AmountOfEmployees: 1,
				Type:              CompanyType(-1),
			},
		},
		"Correct CompanyType": {
			nil,
			Company{
				ID:                CompanyId(uuid.NewV7()),
				Name:              "A",
				AmountOfEmployees: 1,
				Type:              CompanyTypeCorporations,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if err := tt.company.Validate(); errors.Is(err, tt.expErr) == false {
				t.Error(fmt.Errorf("testcase '%s' got (%w), expected (%w)", name, err, tt.expErr))
			}
		})
	}
}
