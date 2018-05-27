package http

import (
	"net/http"

	"github.com/napalm684/recipes/repositories"
)

// RecipeHandler contains recipe-specific http.HandlerFuncs as
// methods.
type RecipeHandler struct {
	recipeService repositories.RecipeService
}

// All lists all recipes in the system.
func (h *RecipeHandler) All(w http.ResponseWriter, r *http.Request) {
	// Query for all recipes.
	// recipes, err := h.recipeService.All()
	// if err != nil {
	// 	h.renderIndexError(w, r, err)
	// 	return
	// }
	// err = h.renderIndexSuccess(w, r, widgets)
	// if err != nil {
	// 	h.renderIndexError(w, r, err)
	// 	return
	// }
	w.Write(([]byte)("All called!"))
}
