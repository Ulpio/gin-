package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ulpio/gin-api-golang/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTestRoutes() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestStatusCode(t *testing.T) {
	r := SetupTestRoutes()                           // SetupTestRoutes() é uma função que retorna um objeto do tipo *gin.Engine
	r.GET("/alunos", controllers.ExibeTodosAlunos)   // Adiciona a rota /alunos ao objeto r
	req, _ := http.NewRequest("GET", "/alunos", nil) // Cria uma requisição do tipo GET para a rota /alunos
	resp := httptest.NewRecorder()                   // Cria um objeto do tipo *httptest.ResponseRecorder
	r.ServeHTTP(resp, req)                           // Executa a requisição
	if resp.Code != http.StatusOK {                  // Verifica se o código de status da resposta é 200
		t.Errorf("Status code should be 200, but got %d", resp.Code) // Se não for, exibe uma mensagem de erro
	}
}
