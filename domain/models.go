package domain

// Recipe represents the metadata/header data
// associated to a recipe in the system.
type Recipe struct {
	ID              int
	Name            string
	Description     string
	DurationMinutes float32
	Source          string
}

// Ingredient represents the data describing
// a recipe ingredient in the system.
type Ingredient struct {
	ID       int
	ItemName string
	RecipeID int
	Quantity float32
	UOMID    int
}

// UOM represents the unit of measure for
// an ingredient in the system.
type UOM struct {
	ID        int
	UnitName  string
	ShortName string
}
