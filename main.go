package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Car struct {
	Make       string
	Model      string
	Horsepower string
}

func main() {
	db, err := sql.Open("cars", "postgresql://localhost:5432/bp21")
	if err != nil {
		log.Fatalf("could not connect to the db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach db: %v", err)
	}
	fmt.Println("Reached DB")

	row := db.QueryRow("SELECT Make, description FROM cars LIMIT 1")

	car := Car{}

	if err := row.Scan(&car.Make, &car.Model, car.Horsepower); err != nil {
		log.Fatalf("Could not scan rows: %v", err)

	}

	rows, err := db.Query("SELECT Make, description FROM cars Limit 10")
	if err != nil {
		log.Fatalf("couldn't execute query: %v", err)
	}

	cars := []Car{}

	for rows.Next() {
		car := Car{}
		if err := rows.Scan(&car.Make, &car.Model, &car.Horsepower); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}

		cars = append(cars, car)
	}
	fmt.Printf("found %d cars: %+v", len(cars), cars)

	carName := "Ferrari"

	row := db.QueryRow("SELECT make, description FROM cars, WHERE make = $1 LIMIT $2", carName, 1)
}
