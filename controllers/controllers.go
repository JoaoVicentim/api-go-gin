package controllers

import (
	"net/http"

	"github.com/JoaoVicentim/api-go-gin/database"
	"github.com/JoaoVicentim/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) { //esse parametro é padrão
	//criando lista de alunos
	var alunos []models.Aluno
	//Encontrar todos os alunos da lista no banco de dados
	database.DB.Find(&alunos)
	//retorna o status code e uma mensagem
	c.JSON(200, alunos)
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

	//Validação dos dados do aluno
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	//se nao der erro, salvar as infos no banco de dados
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorID(c *gin.Context) {
	//precisamos de um endereço de memoria de um aluno
	var aluno models.Aluno
	//pegar o ID que colocamos em routes
	id := c.Params.ByName("id")
	//Encontrar o primeiro aluno com o id no DB
	database.DB.First(&aluno, id)

	//tratando caso insira um aluno que nao existe
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	//identificando aluno
	var aluno models.Aluno
	//pegando o id do aluno que vamos deletar
	id := c.Params.ByName("id")
	//excluindo do banco de dados
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	//pesquisando aluno
	database.DB.First(&aluno, id)

	//a requisição PATCH vai ter um corpo com as informações que quero editar do aluno
	//Empacota todo o corpo da requisição com base no endereço de aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	//Validação dos dados do aluno
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	//pegando o modelo de banco de dados do aluno e atualizando as infos do aluno
	//database.DB.Model(&aluno).UpdateColumns(aluno) //pega o modelo de aluno e atualizando todas as colunas com base no aluno que passamos no corpo da requisição

	database.DB.Save(&aluno) //atualiza todas as colunas do aluno
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")

	//pesquisando no banco de dados um aluno com o cpf específico
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	//tratando caso insira um aluno que nao existe
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	///exibir os alunos do banco de dados
	//criando uma lista de alunos
	var alunos []models.Aluno
	//encontra todos os alunos e coloca dentro da variavel alunos
	database.DB.Find(&alunos)
	//segunda informacao, qual o nome do template que queremos renderizar
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

// funcao para rota nao encontrada
func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
