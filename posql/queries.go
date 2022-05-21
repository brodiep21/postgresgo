package posql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Car struct {
	ID         string `json:"id"`
	Make       string `json:"Make"`
	Model      string `json:"Model"`
	Horsepower string `json:"Horsepower"`
	MSRP       string `json:"MSRP"`
}

func (c *Car) TableInsert(db *sql.DB) (string, error) {

	result, err := db.Exec("INSERT INTO cars (make, model, horsepower, msrp) VALUES ($1, $2, $3, $4)", &c.Make, &c.Model, &c.Horsepower, &c.MSRP)
	if err != nil {
		return "", err
	}

	changedRows, err := result.RowsAffected()
	if err != nil {
		return "", err
	}
	fmt.Println("inserted", changedRows, "rows")

	car := Car{}
	row := db.QueryRow("SELECT Make, Model, Horsepower, MSRP FROM cars ")

	if err := row.Scan(&car.Make, &car.Model, &car.Horsepower, &car.MSRP); err != nil {
		return "", err
	}

	return "Added in " + car.Make + " " + car.Model + " " + car.Horsepower + " " + car.MSRP, nil
}
