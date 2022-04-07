package main

import (
	"github.com/Dagurasu56/go-gin-api/models"
	"github.com/Dagurasu56/go-gin-api/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{"Théo Souza", "176.123.123-12", "38000000"},
		{"Débora Souza", "123.533.123-12", "54000000"},
	}
	routes.HandleRequests()
}
