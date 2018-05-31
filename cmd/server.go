package main

import (
	"database/sql"
	"log"
	stdhttp "net/http"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
	"github.com/napalm684/recipes/http"
	"github.com/napalm684/recipes/repositories"
	"github.com/napalm684/recipes/sqlite"
)

var (
	recipeService repositories.RecipeService
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := path.Dir(ex)
	log.Print(dir)

	db, err := sql.Open("sqlite3", "./recipes_app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	recipeService = &sqlite.RecipeService{DB: db}

	server := http.NewServer(recipeService)
	log.Fatal(stdhttp.ListenAndServe(":3000", server))
}
