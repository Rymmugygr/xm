package response

import "github.com/Rymmugygr/xm/internal/app/entity"

type Company struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	AmountOfEmployees int    `json:"amountOfEmployees"`
	Registered        bool   `json:"registered"`
	Type              string `json:"type"`
}

func (c *Company) FromEntity(input *entity.Company) Company {
	return Company{
		Id:                input.ID.String(),
		Name:              input.Name,
		Description:       input.Description,
		AmountOfEmployees: input.AmountOfEmployees,
		Registered:        input.Registered,
		Type:              input.Type.String(),
	}
}
