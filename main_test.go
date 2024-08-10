package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JoaoVicentim/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
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

//funcao de teste que sabemos que vai falhar
//ponteiro apontando para o teste que vamos fazer
// func TestFalhador(t *testing.T) {
// 	t.Fatalf("Teste falhou de propósito, não se preocupe!")
// }
