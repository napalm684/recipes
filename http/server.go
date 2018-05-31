package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/napalm684/recipes/repositories"
)

// server is our HTTP server with routes for all our endpoints.
type server struct {
	recipes *RecipeHandler
	router  *mux.Router
}

// NewServer returns an http.Handler for a server providing
// REST APIs using JSON.
func NewServer(rs repositories.RecipeService) http.Handler {
	json := JSONServer(rs)
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", json))
	return mux
}

// JSONServer will construct a server, apply all of the necessary routes,
// then return an http.Handler for the server.
func JSONServer(rs repositories.RecipeService) http.Handler {
	rh := &RecipeHandler{
		recipeService: rs,
		renderJSON:    renderJSON,
	}
	s := server{
		recipes: rh,
		router:  mux.NewRouter(),
	}
	s.routes()
	return &s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routes() {
	// Recipes routes
	s.router.Handle("/recipes", ApplyMwFn(s.recipes.All)).Methods("GET")
}
