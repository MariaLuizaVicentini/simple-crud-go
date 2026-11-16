package main

import (
	"log"
	"net/http"
	"github.com/marialuizavicentini/simple-crud-go/config"
)

// É a função chamada quando a aplicação é iniciada
func main(){
	dbConnection := config.SetupDB()

	// chama a função close pra fechar a conexao com o banco de dados
	defer dbConnection.Close()

	log.Fatal(http.ListenAndServe(":8000", nil))


}