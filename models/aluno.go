package models

import (
	"gopkg.in/validator.v2" // pacote para validar as entradas da api
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`                // nao permite valor vazio
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`   // nao permite que tamanho seja diferente de 9 e somente sao permitidos numeros
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"` // nao permite que tamanho seja diferente de 11 e somente sao permitidos numeros
}

func ValidadaDadosDeAlunos(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
