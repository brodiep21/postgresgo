package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "bp21"
	dbname = "bp21"
)

var password = os.Getenv("pass")

type Car struct {
	Make       string
	Model      string
	Horsepower string
	MSRP       string
}

func init() {

}
func main() {

	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("could not connect to the db: %v", err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(7)
	db.SetConnMaxIdleTime(4 * time.Second)
	db.SetConnMaxLifetime(20 * time.Second)

	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach db: %v", err)
	}
	fmt.Println("Reached DB")

	defer db.Close()

	// cars := []Car{}

	// for rows.Next() {
	// 	car := Car{}
	// 	if err := rows.Scan(&car.Make, &car.Model, &car.Horsepower); err != nil {
	// 		log.Fatalf("could not scan row: %v", err)
	// 	}

	// 	cars = append(cars, car)
	// }
	// fmt.Printf("found %d cars: %+v", len(cars), cars)

	//user input for Make of vehicle
	var CarMake string
	fmt.Println("What Make are you looking for?")
	fmt.Scanln(&CarMake)
	// newCar :=

	// row = db.QueryRow("SELECT make, description FROM cars, WHERE make = $1 LIMIT $2", carName, 1)

	row := db.QueryRow("SELECT Make, Model, Horsepower, MSRP FROM cars LIMIT 1")

	car := Car{}

	if err := row.Scan(&car.Make, &car.Model, car.Horsepower); err != nil {
		log.Fatalf("Could not scan rows: %v", err)

	}

	// rows, err := db.Query("SELECT Make, Model, Horsepower, MSRP FROM cars Limit 10")
	// if err != nil {
	// 	log.Fatalf("couldn't execute query: %v", err)
	// }

}
