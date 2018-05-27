/*
sqlite3 recipes_app.db
*/

CREATE TABLE Recipes(
    Id INTEGER PRIMARY KEY,
    Name TEXT,
    Description TEXT,
    DurationMinutes DECIMAL(5,2),
    Source TEXT
);

CREATE TABLE Ingredients (
    	Id       INTEGER PRIMARY KEY,
	ItemName TEXT,
	RecipeID INTEGER,
	Quantity DECIMAL(5,2),
	UOMID    INTEGER
);

CREATE TABLE UnitsOfMeasure (
	Id        INTEGER PRIMARY KEY,
	UnitName  TEXT,
	ShortName TEXT
);
