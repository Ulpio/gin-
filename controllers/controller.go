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
	if err := models.Validator(&aluno); err != nil {
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

func DeleteAluno(c *gin.Context) {
	var id = c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)
	if aluno.Nome == "" {
		c.JSON(http.StatusNotFound, gin.H{"data": "Aluno não encontrado"})
		return
	}
	database.DB.Delete(&aluno)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

func EditaAluno(c *gin.Context) {
	var id = c.Params.ByName("id") //Recebe o parametro id da URL
	var aluno models.Aluno         //Cria uma variável do tipo Aluno
	database.DB.First(&aluno, id)  //Busca o aluno no banco de dados
	if aluno.Nome == "" {          //Se o aluno não for encontrado
		c.JSON(http.StatusNotFound, gin.H{"data": "Aluno não encontrado"}) //Retorna um erro
		return                                                             //Encerra a função
	}
	if err := models.Validator(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&aluno); err != nil { //Se houver erro ao fazer o bind do JSON
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //Retorna um erro
		return                                                     //Encerra a função
	}
	database.DB.Save(&aluno)     //Salva as alterações no banco de dados
	c.JSON(http.StatusOK, aluno) //Retorna o aluno alterado
}

func GetAlunoPeloCPF(c *gin.Context) {
	var cpf = c.Params.ByName("cpf")                //Recebe o parametro cpf da URL
	var aluno models.Aluno                          //Cria uma variável do tipo Aluno
	database.DB.Where("cpf = ?", cpf).First(&aluno) //Busca o aluno no banco de dados
	if aluno.Nome == "" {                           //Se o aluno não for encontrado
		c.JSON(http.StatusNotFound, gin.H{"data": "Aluno não encontrado"}) //Retorna um erro
		return                                                             //Encerra a função
	}
	c.JSON(http.StatusOK, aluno) //Retorna o aluno encontrado
}
