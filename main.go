package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brodiep21/postgresgo/search"
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

	//user input for Make
	fmt.Println("What Make are you looking for?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	make := scanner.Text()
	//user input for Model
	fmt.Println("What Model are you looking for?")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	model := scanner2.Text()

	fullvehicle := make + " " + model

	fmt.Println("Captured:", fullvehicle)

	hp := search.HorsepowerSearch(fullvehicle)
	msrp := search.MsrpSearch(fullvehicle)

	fmt.Println(hp, msrp)

	// fmt.Println("Would you like to add this data into the table?")
	// scanner = bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// response := scanner.Text()
	// switch response {
	// case "y", "yes", "YES", "Y":

	// }
	newCar := Car{
		Make:       make,
		Model:      model,
		Horsepower: hp,
		MSRP:       msrp,
	}

	result, err := db.Exec("INSERT INTO cars (make, model, horsepower, msrp) VALUES ($1, $2, $3, $4)", newCar.Make, newCar.Model, newCar.Horsepower, newCar.MSRP)
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

	if err := row.Scan(&car.Make, &car.Model); err != nil {
		log.Fatalf("Could not scan rows: %v", err)
	}

	fmt.Printf("found car %+v \n", car)

	// rows, err := db.Query("SELECT Make, Model, Horsepower, MSRP FROM cars Limit 10")
	// if err != nil {
	// 	log.Fatalf("couldn't execute query: %v", err)
	// }

}
