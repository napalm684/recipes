package http

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/napalm684/recipes/domain"
	"github.com/napalm684/recipes/repositories"
)

// RecipeHandler contains recipe-specific http.HandlerFuncs as
// methods.
type RecipeHandler struct {
	recipeService repositories.RecipeService
	renderJSON    func(w http.ResponseWriter, data interface{}, status int)
}

// All lists all recipes in the system.
func (h *RecipeHandler) All(w http.ResponseWriter, r *http.Request) {
	// Query for all recipes.
	recipes, _ := h.recipeService.All() //TODO Error handling
	if len(recipes) == 0 {
		h.renderJSON(w, recipes, 204)
		return
	}
	h.renderJSON(w, recipes, 200)
	// if err != nil {
	// 	h.renderIndexError(w, r, err)
	// 	return
	// }
	// err = h.renderIndexSuccess(w, r, widgets)
	// if err != nil {
	// 	h.renderIndexError(w, r, err)
	// 	return
	// }
}

// Get retrieves recipe in system using the ID passed on the route.
func (h *RecipeHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	recipe, _ := h.recipeService.Get(int(id)) //TODO error Handling
	if &recipe == nil || (domain.Recipe{}) == recipe {
		h.renderJSON(w, recipe, 204)
	}
	h.renderJSON(w, recipe, 200)
}
