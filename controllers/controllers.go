package controllers

import (
	"net/http"

	"github.com/JoaoVicentim/api-go-gin/database"
	"github.com/JoaoVicentim/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) { //esse parametro é padrão
	//retorna o status code e uma mensagem,
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

func CriaNovoAluno(c *gin.Context) {
	//precisamos do endereço de memória de um aluno, para ele conseguir vincular o corpo da requisição com a struct
	var aluno models.Aluno
	//funcao do gin que pega todo o corpo da requisição e empacota com base na struct &aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	//se nao der erro, salvar as infos no banco de dados
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
