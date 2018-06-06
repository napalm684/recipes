package sqlite

import (
	"database/sql"

	"github.com/napalm684/recipes/domain"
)

// IngredientService is a SQLite implementation of a
// ingredient datastore.
type IngredientService struct {
	DB *sql.DB
}

// Get will return a slice of ingredients for the recipeID parameter passed and
// an additional error parameter should a problem occur.
func (i *IngredientService) Get(recipeID int) ([]domain.Ingredient, error) {
	ingredients := make([]domain.Ingredient, 0)
	rows, err := i.DB.Query(`SELECT Id, ItemName, RecipeID, Quantity, UOMID 
		FROM Ingredients WHERE RecipeID=?;`, recipeID)
	if err != nil {
		return []domain.Ingredient{}, nil
	}
	defer rows.Close()

	for rows.Next() {
		ingredient := domain.Ingredient{}
		if err := rows.Scan(&ingredient.ID, &ingredient.ItemName,
			&ingredient.RecipeID, &ingredient.Quantity, &ingredient.UOMID); err != nil {
			return []domain.Ingredient{}, nil
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}
