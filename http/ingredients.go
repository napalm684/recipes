package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/napalm684/recipes/domain"
	"github.com/napalm684/recipes/repositories"
)

// IngredientHandler contains ingredient-specific http.HandlerFuncs as
// methods.
type IngredientHandler struct {
	ingredientService repositories.IngredientService
	renderJSON        func(w http.ResponseWriter, data interface{}, status int)
}

// Create adds new recipes in system based on the JSON payload posted
// with the request.
func (h *IngredientHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ingredient domain.Ingredient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ingredient); err != nil {
		renderJSON(w, "Invalid request payload", http.StatusBadGateway)
		return
	}
	defer r.Body.Close()

	id, err := h.ingredientService.Create(ingredient)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.renderJSON(w, id, 200)
	return
}

// Get retrieves ingredients in system using the RecipeID passed on the route.
func (h *IngredientHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recipeID, err := strconv.ParseInt(vars["recipeID"], 10, 32)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	ingredients, _ := h.ingredientService.Get(int(recipeID))
	if &ingredients == nil || len(ingredients) == 0 {
		h.renderJSON(w, ingredients, 204)
	}
	h.renderJSON(w, ingredients, 200)
}

// Delete removes ingredient in system using the IngredientID passed on the route.
func (h *IngredientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ingID, err := strconv.ParseInt(vars["ingID"], 10, 32)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = h.ingredientService.Delete(int(ingID))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(204)
	return
}
