package routes

import (
	"github.com/Ulpio/gin-api-golang/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/saudacao/:nome", controllers.Saudacao)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/cpf/:cpf", controllers.GetAlunoPeloCPF)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorID)
	r.PUT("/alunos/:id", controllers.EditaAluno)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.GET("/", controllers.ShowIndex)
	r.NoRoute(controllers.Rota404)
	r.Run()
}
