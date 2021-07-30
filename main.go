package main

import (
	"net/http"

	"br.com.github/raphasalomao/go-marvel-heroes-api/api"
	"br.com.github/raphasalomao/go-marvel-heroes-api/api/database"
)

func main() {
	database.InitDatabase()
	api.InitController()
	http.ListenAndServe(":8080", api.Router)
}
