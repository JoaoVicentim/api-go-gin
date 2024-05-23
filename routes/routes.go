package routes

import (
	"github.com/JoaoVicentim/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

// funcao que lida com requisições
func HandleRequests() {
	//variavel que vai configurar a aplicação do gin
	r := gin.Default() //usando caracterisicas default do Gin
	//quando chegar uma requisição GET para /alunos, quem vai atender vai ser o ExibeTodosAlunos
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	//subindo o servidor
	r.Run()
}
