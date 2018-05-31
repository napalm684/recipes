package http

import (
	"net/http"

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
