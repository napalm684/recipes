package sqlite

import (
	"database/sql"

	"github.com/napalm684/recipes/domain"
)

// RecipeService is a SQLite implementation of a
// recipe datastore.
type RecipeService struct {
	DB *sql.DB
}

// All will return a slice of every recipe in the system,
// returning an additional error parameter should a problem
// occur.
func (r *RecipeService) All() ([]domain.Recipe, error) {
	recipes := make([]domain.Recipe, 0)
	rows, err := r.DB.Query("SELECT Id, Name, Description, DurationMinutes, Source, Serves FROM Recipes")
	if err != nil {
		return []domain.Recipe{}, nil
	}
	defer rows.Close()

	for rows.Next() {
		recipe := domain.Recipe{}
		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.DurationMinutes, &recipe.Source, &recipe.Serves); err != nil {
			return []domain.Recipe{}, nil
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

// Get will return a recipe for the id paramter passed, and
// an additional error parameter should a problem occur.
func (r *RecipeService) Get(id int) (domain.Recipe, error) {
	recipe := domain.Recipe{}
	err := r.DB.QueryRow("SELECT Id, Name, Description, DurationMinutes, Source, Serves FROM Recipes WHERE Id=?", id).Scan(
		&recipe.ID, &recipe.Name, &recipe.Description, &recipe.DurationMinutes, &recipe.Source, &recipe.Serves)
	return recipe, err
}

// Create will create a recipe for storage in the system and return the created record ID. An
// error parameter is returned if a problem occurs during the
// execution.
func (r *RecipeService) Create(recipe domain.Recipe) (int, error) {
	result, err := r.DB.Exec(`INSERT INTO Recipes (Name, Description, DurationMinutes, Source, Serves) 
		VALUES ($1, $2, $3, $4, $5)`, recipe.Name, recipe.Description,
		recipe.DurationMinutes, recipe.Source, recipe.Serves)

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), err
}

// Delete will remove a recipe matching the ID passed from storage.
// An error parameter is returned if a problem occurs during the execution.
func (r *RecipeService) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM Recipes WHERE ID = $1`, id)
	return err
}
