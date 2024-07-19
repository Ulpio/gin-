package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupTestRoutes() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestFalha(t *testing.T) {
	t.Fatalf("Teste falhou, pois foi forçado a falhar")
}
