package request

type CreateCompany struct {
	Name              string `json:"name"              validate:"required,min=1,max=15"`
	Description       string `json:"description"       validate:"max=3_000"`
	AmountOfEmployees int    `json:"amountOfEmployees" validate:"required,max=2_000_000_000"`
	Registered        bool   `json:"registered"        validate:"required"`
	Type              string `json:"type"              validate:"required"`
}

type PatchCompany struct {
	Name              *string `json:"name"              validate:"omitempty,required,min=1,max=15"`
	Description       *string `json:"description"       validate:"omitempty,max=3_000"`
	AmountOfEmployees *int    `json:"amountOfEmployees" validate:"omitempty,required,max=2_000_000_000"`
	Registered        *bool   `json:"registered"        validate:"omitempty,required"`
	Type              *string `json:"type"              validate:"omitempty,required"`
}
