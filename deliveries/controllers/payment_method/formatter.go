package payment_method

import (
	PM "group-project1/entities/payment_method"

	"gorm.io/gorm"
)

// =================== Create Payment Method =======================
type CreatePMRequestFormat struct {
	Name string `json:"name" form:"name"`
}

type CreatePMResponseFormat struct {
	Name string `json:"name"`
}

func ToCreatePMResponseFormat(PMResponse PM.PaymentMethods) CreatePMResponseFormat {
	return CreatePMResponseFormat{
		Name: PMResponse.Name,
	}
}

// =================== Update PM =======================
type UpdatePMRequestFormat struct {
	Name string `json:"name" form:"name"`
}

func (UURF UpdatePMRequestFormat) ToUpdatePMRequestFormat(ID uint) PM.PaymentMethods {
	return PM.PaymentMethods{
		Model: gorm.Model{ID: ID},
		Name:  UURF.Name,
	}
}

type UpdatePMResponseFormat struct {
	Name string `json:"name"`
}

func ToUpdatePMResponseFormat(PMResponse PM.PaymentMethods) UpdatePMResponseFormat {
	return UpdatePMResponseFormat{
		Name: PMResponse.Name,
	}
}

type PMGetResponseFormat struct {
	Name string `json:"name"`
}

func ToPMGetResponseFormat(PMResponses []PM.PaymentMethods) []PMGetResponseFormat {
	PMGetResponses := make([]PMGetResponseFormat, len(PMResponses))
	for i := 0; i < len(PMResponses); i++ {
		PMGetResponses[i].Name = PMResponses[i].Name
	}
	return PMGetResponses
}