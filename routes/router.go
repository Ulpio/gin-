package routes

import (
	"github.com/Ulpio/gin-api-golang/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/cpf/:cpf", controllers.GetAlunoPeloCPF)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorID)
	r.PUT("/alunos/:id", controllers.EditaAluno)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	
	r.Run()
}
