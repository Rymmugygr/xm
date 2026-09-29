package postgres

import (
	"fmt"
	"testing"

	"github.com/Rymmugygr/xm/internal/app/entity"
)

func TestCompanyType_Convert1(t *testing.T) {
	tests := map[string]struct {
		exp         string
		companyType entity.CompanyType
	}{
		"Invalid Company Type (entity)":                       {"unknown", entity.CompanyTypeWrong},
		"Correct Company Type (entity) (corporations)":        {"corporations", entity.CompanyTypeCorporations},
		"Correct Company Type (entity) (non profit)":          {"non_profit", entity.CompanyTypeNonProfit},
		"Correct Company Type (entity) (cooperative)":         {"cooperative", entity.CompanyTypeCooperative},
		"Correct Company Type (entity) (sole proprietorship)": {"sole_proprietorship", entity.CompanyTypeSoleProprietorship},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := companyTypeToModel(tt.companyType); got != tt.exp {
				t.Error(fmt.Errorf("testcase '%s' got (%s), expected (%s)", name, got, tt.exp))
			}
		})
	}
}

func TestCompanyType_Convert2(t *testing.T) {
	tests := map[string]struct {
		exp         entity.CompanyType
		companyType string
	}{
		"Invalid Company Type (model) (empty)":               {entity.CompanyTypeWrong, ""},
		"Invalid Company Type (model) (not existed)":         {entity.CompanyTypeWrong, "abcd"},
		"Correct Company Type (model) (corporations)":        {entity.CompanyTypeCorporations, "corporations"},
		"Correct Company Type (model) (non profit)":          {entity.CompanyTypeNonProfit, "non_profit"},
		"Correct Company Type (model) (cooperative)":         {entity.CompanyTypeCooperative, "cooperative"},
		"Correct Company Type (model) (sole proprietorship)": {entity.CompanyTypeSoleProprietorship, "sole_proprietorship"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := companyTypeToEntity(tt.companyType); got != tt.exp {
				t.Error(fmt.Errorf("testcase '%s' got (%s), expected (%s)", name, got, tt.exp))
			}
		})
	}
}
