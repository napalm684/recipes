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
		r := domain.Recipe{}
		if err := rows.Scan(&r.ID, &r.Name, &r.Description, &r.DurationMinutes, &r.Source, &r.Serves); err != nil {
			return []domain.Recipe{}, nil
		}
		recipes = append(recipes, r)
	}

	return recipes, nil
}
