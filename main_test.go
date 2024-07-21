package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Ulpio/gin-api-golang/controllers"
	"github.com/Ulpio/gin-api-golang/database"
	"github.com/Ulpio/gin-api-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	ID int
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome de Teste", CPF: "12312312300", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

// func TestCriaAluno(t *testing.T) {
// 	database.ConnectDB()
// 	rotas := SetupTestRoutes()
// 	rotas.POST("/alunos", controllers.CriaNovoAluno)

// 	req, _ := http.NewRequest("POST", "/alunos", strings.NewReader(`
// 		{
// 		"nome":"Ulpiooo",
// 		"cpf":"13366671416",
// 		"rg":"41510771"
// 		}
// 		`))
// 	resp := httptest.NewRecorder()
// 	rotas.ServeHTTP(resp, req)
// 	if resp.Code != http.StatusCreated {
// 		t.Fatalf("Aluno nao Criado")
// 	}
// }

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestSaudacaoStatusCode(t *testing.T) {
	rotas := SetupTestRoutes()
	rotas.GET("/saudacao/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/saudacao/ulpio", nil)
	resp := httptest.NewRecorder()
	rotas.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Response code is %v", resp.Code)
	}
	responseMock := `{"data":"Olá ulpio"}`
	respondeBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, responseMock, string(respondeBody))
}

func TestListandoTodosAlunosHandler(t *testing.T) {
	database.ConnectDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupTestRoutes()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuscaPorCPF(t *testing.T) {
	database.ConnectDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupTestRoutes()
	r.GET("/alunos/cpf/:cpf", controllers.GetAlunoPeloCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/13366671416", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuscaAlunoPorID(t *testing.T) {
	database.ConnectDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupTestRoutes()
	r.GET("/alunos/:id", controllers.ExibeAlunoPorID)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var alunoMock models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome de Teste", alunoMock.Nome)
	assert.Equal(t, "12312312300", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resp.Code)
}
