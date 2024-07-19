package controllers

import (
	"net/http"

	"github.com/Ulpio/gin-api-golang/database"
	"github.com/Ulpio/gin-api-golang/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func ExibeAlunoPorID(c *gin.Context) {
	var id = c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)
	if aluno.Nome == "" {
		c.JSON(http.StatusNotFound, gin.H{"data": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
