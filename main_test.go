package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/JoaoVicentim/api-go-gin/controllers"
	"github.com/JoaoVicentim/api-go-gin/database"
	"github.com/JoaoVicentim/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	//criando um aluno mock
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	//criando no banco de dados
	database.DB.Create(&aluno)
	//armazenando o ID do aluno
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	//criando instancia de aluno
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParamentro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	//dizendo como a requisicao vai ser feita
	req, _ := http.NewRequest("GET", "/Joao", nil)
	//armazenar a resposta
	resposta := httptest.NewRecorder()
	//realizando a requisicao
	r.ServeHTTP(resposta, req)
	//passa a instacia do teste, o status code que esperamos e o status code que recebemos
	assert.Equal(t, 200, resposta.Code, "Deveriam ser iguais")

	//quero verificar o corpo da requisição
	//mock significa o que esperamos que aconteça
	mockDaResposta := `{"API diz:":"E ai Joao, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody), "Deveriam ser iguais")
}

func TestListaTodosOsAlunosHandler(t *testing.T) {
	//para listar todos os alunos, precisamos acessar o banco de dados (dev)
	database.ConectaComBancoDeDados()

	//criando um aluno teste, para garantir que sempre tenha pelo menos um aluno
	CriaAlunoMock()
	defer DeletaAlunoMock()
	//rota do gin
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

	//verificando corpo da resposta
	fmt.Println(resposta.Body)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	//criando a rota de busca
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	//vai pegar todos os bytes e transformar em   json e armazenar na variavel AlunoMock
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	fmt.Println(alunoMock.Nome)
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaUmAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	//armazenando a configuracao do aluno
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "47123456789", RG: "123456700"}
	//transformando em json
	valorJson, _ := json.Marshal(aluno)
	//passar o path que vamos editar
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	//o valor em json que vamos passar deve ser passado em forma de bytes
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	//armazenar a resposta
	resposta := httptest.NewRecorder()
	//realizar a requisicao
	r.ServeHTTP(resposta, req)
	fmt.Println("Response Body:", resposta.Body.String())
	//criano aluno para fazer verificação
	var alunoMockAtualizado models.Aluno
	//transformar o corpo da resposta em json e armazenar na variavel alunoMockAtualizado
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	//verificar se o nome foi alterado
	//primeiro passar a instancia do teste, segunda parametro é o que esperamos, terceiro é o que recebemos
	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome)
	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456700", alunoMockAtualizado.RG)

	fmt.Println(alunoMockAtualizado.Nome)
}

//funcao de teste que sabemos que vai falhar
//ponteiro apontando para o teste que vamos fazer
// func TestFalhador(t *testing.T) {
// 	t.Fatalf("Teste falhou de propósito, não se preocupe!")
// }
