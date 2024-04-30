package main

import (
	"api_rest_gin_go_2-validacoes-e-testes/api-go-gin/database"
	"api_rest_gin_go_2-validacoes-e-testes/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
