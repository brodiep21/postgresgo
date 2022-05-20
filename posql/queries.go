package posql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Car struct {
	Make       string `json:"Make"`
	Model      string `json:"Model"`
	Horsepower string `json:"Horsepower"`
	MSRP       string `json:"MSRP"`
}

func (c *Car) TableInsert(db *sql.DB) string {

	result, err := db.Exec("INSERT INTO cars (make, model, horsepower, msrp) VALUES ($1, $2, $3, $4)", &c.Make, &c.Model, &c.Horsepower, &c.MSRP)
	if err != nil {
		log.Fatalf("could not insert into cars, %v", err)
	}

	changedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("could not get rows affected by INSERT %v", err)
	}
	fmt.Println("inserted", changedRows, "rows")

	car := Car{}
	row := db.QueryRow("SELECT Make, Model, Horsepower, MSRP FROM cars ")

	if err := row.Scan(&car.Make, &car.Model, &car.Horsepower, &car.MSRP); err != nil {
		log.Fatalf("Could not scan rows: %v", err)
	}

	return "Added in " + car.Make + " " + car.Model + " " + car.Horsepower + " " + car.MSRP
}
