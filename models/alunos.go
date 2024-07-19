package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero,min=3,max=100"`
	CPF  string `json:"cpf"  validate:"max=15,min=11, regexp=^[0-9]{3}\\.?[0-9]{3}\\.?[0-9]{3}\\-?[0-9]{2}$"`
	RG   string `json:"rg"   validate:"max=9"`
}	

func Validator(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
