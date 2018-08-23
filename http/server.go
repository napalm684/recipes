package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/napalm684/recipes/repositories"
)

// server is our HTTP server with routes for all our endpoints.
type server struct {
	recipes     *RecipeHandler
	ingredients *IngredientHandler
	router      *mux.Router
}

// NewServer returns an http.Handler for a server providing
// REST APIs using JSON.
func NewServer(rs repositories.RecipeService, is repositories.IngredientService) http.Handler {
	json := JSONServer(rs, is)
	mux := http.NewServeMux()
	mux.Handle("/", StaticServer("../client/build"))
	mux.Handle("/api/", json)
	return mux
}

// JSONServer will construct a server, apply all of the necessary routes,
// then return an http.Handler for the server.
func JSONServer(rs repositories.RecipeService, is repositories.IngredientService) http.Handler {
	rh := &RecipeHandler{
		recipeService: rs,
		renderJSON:    renderJSON,
	}
	ih := &IngredientHandler{
		ingredientService: is,
		renderJSON:        renderJSON,
	}
	s := server{
		recipes:     rh,
		ingredients: ih,
		router:      mux.NewRouter(),
	}
	s.routes()
	return &s
}

// StaticServer will construct a server to serve up static assets.
func StaticServer(path string) http.Handler {
	return http.FileServer(http.Dir(path))
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routes() {
	ar := s.router.PathPrefix("/api/").Subrouter()

	// Recipes routes
	ar.Handle("/recipes", ApplyMwFn(s.recipes.All)).Methods("GET")
	ar.Handle("/recipes/{id}", ApplyMwFn(s.recipes.Get)).Methods("GET")
	ar.Handle("/recipes", ApplyMwFn(s.recipes.Create)).Methods("POST")
	ar.Handle("/recipes/{id}", ApplyMwFn(s.recipes.Delete)).Methods("DELETE")
	ar.Handle("/recipes", ApplyMwFn(s.recipes.Update)).Methods("PUT")

	// Ingredient routes
	ar.Handle("/ingredients/recipes/{recipeID}", ApplyMwFn(s.ingredients.Get)).Methods("GET")
	ar.Handle("/ingredients/{ingID}", ApplyMwFn(s.ingredients.Delete)).Methods("DELETE")
	ar.Handle("/ingredients", ApplyMwFn(s.ingredients.Create)).Methods("POST")
}
