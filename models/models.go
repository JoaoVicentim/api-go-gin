package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	//valores serao devolvidos em formato json
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

// Lista de Alunos referente à struct de Aluno
var Alunos []Aluno
