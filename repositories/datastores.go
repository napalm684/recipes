package repositories

import (
	"github.com/napalm684/recipes/domain"
)

// RecipeService defines the contract for
// recipe data store interactions.
type RecipeService interface {
	All() ([]domain.Recipe, error)
	Get(id int) (domain.Recipe, error)
	Create(recipe domain.Recipe) (int, error)
	Delete(id int) error
	Update(domain.Recipe) error
}

// IngredientService defines the contract
// for ingredient store interactions.
type IngredientService interface {
	Get(recipeID int) ([]domain.Ingredient, error)
	Delete(ingID int) error
}
