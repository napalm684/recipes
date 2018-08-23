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

// Create will create an ingredient for storage in the system and return the created record ID. An
// error parameter is returned if a problem occurs during the execution.
func (i *IngredientService) Create(ingredient domain.Ingredient) (int, error) {
	result, err := i.DB.Exec(`INSERT INTO Ingredients (ItemName, RecipeID, Quantity, UOMID)
		VALUES ($1, $2, $3, $4);`, ingredient.ItemName, ingredient.RecipeID,
		ingredient.Quantity, ingredient.UOMID)

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), err
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

// Delete will remove an ingredient matching the ID passed from storage.
// An error parameter is returned if a problem occurs during the execution.
func (i *IngredientService) Delete(ingID int) error {
	_, err := i.DB.Exec(`DELETE FROM Ingredients WHERE ID = $1;`, ingID)
	return err
}
