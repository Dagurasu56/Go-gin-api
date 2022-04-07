package main

import (
	"github.com/Dagurasu56/go-gin-api/database"
	"github.com/Dagurasu56/go-gin-api/models"
	"github.com/Dagurasu56/go-gin-api/routes"
)

func main() {

	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Théo Souza", CPF: "176.123.123-12", RG: "38000000"},
		{Nome: "Débora Souza", CPF: "123.533.123-12", RG: "54000000"},
	}
	routes.HandleRequests()
}
