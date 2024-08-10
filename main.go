package main

import (
	"github.com/JoaoVicentim/api-go-gin/database"
	"github.com/JoaoVicentim/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	// models.Alunos = []models.Aluno{
	// 	{Nome: "Joao Vicentim", CPF: "00000000000", RG: "202020202"},
	// 	{Nome: "Giovanna Dias", CPF: "11111111111", RG: "101010101"},
	// }
	routes.HandleRequests()
}
