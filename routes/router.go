package routes

import (
	"github.com/Ulpio/gin-api-golang/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.PUT("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.GetAlunoPeloCPF)
	r.Run()
}
