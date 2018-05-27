package main

import (
	"database/sql"
	"log"
	stdhttp "net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/napalm684/recipes/http"
	"github.com/napalm684/recipes/repositories"
)

var (
	recipeService repositories.RecipeService
)

func main() {
	//recipeService = &psql.UserService{DB: db}

	db, err := sql.Open("sqlite3", "./recipes_app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := http.NewServer(recipeService)
	log.Fatal(stdhttp.ListenAndServe(":3000", server))
}
