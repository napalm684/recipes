package repositories

import (
	"github.com/napalm684/recipes/domain"
)

// RecipeService defines the contract for
// recipe data store interactions.
type RecipeService interface {
	All() ([]domain.Recipe, error)
	Get(id int) (domain.Recipe, error)
	Create(recipe domain.Recipe) error
	Delete(id int) error
}
