/*
sqlite3 recipes_app.db
*/

CREATE TABLE Recipes(
    Id INTEGER PRIMARY KEY,
    Name TEXT,
    Description TEXT,
    DurationMinutes DECIMAL(5,2),
    Source TEXT,
	Serves INTEGER
);

CREATE TABLE Ingredients (
	Id       INTEGER PRIMARY KEY,
	ItemName TEXT,
	RecipeID INTEGER,
	Quantity DECIMAL(5,2),
	UOMID    INTEGER
);

CREATE TABLE Steps (
	Id INTEGER PRIMARY KEY,
	StepText TEXT,
	SortOrder INTEGER,
	RecipeID INTEGER
);

CREATE TABLE UnitsOfMeasure (
	Id        INTEGER PRIMARY KEY,
	UnitName  TEXT,
	ShortName TEXT
);
