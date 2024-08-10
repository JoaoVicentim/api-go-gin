package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	//valores serao devolvidos em formato json
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len:11, regexp:^[0-9]*$"`
	RG   string `json:"rg" validate:"len:9, regexp:^[0-9]*$"`
}

// Lista de Alunos referente à struct de Aluno
// var Alunos []Aluno

func ValidaDadosDeAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
