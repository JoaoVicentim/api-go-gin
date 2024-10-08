package routes

import (
	"github.com/JoaoVicentim/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

// funcao que lida com requisições
func HandleRequests() {
	//variavel que vai configurar a aplicação do gin
	r := gin.Default() //usando caracterisicas default do Gin
	//configurando o template (temos paginas para serem renderizadas)
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	//quando chegar uma requisição GET para /alunos, quem vai atender vai ser o ExibeTodosAlunos
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	//PATCH para editar
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	//subindo o servidor
	r.Run()
}
