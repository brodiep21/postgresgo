package posql

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

func TableInsert(make, model, hp, msrp string) string {

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

	if err := row.Scan(&car.Make, &car.Model, &car.Horsepower, &car.MSRP); err != nil {
		log.Fatalf("Could not scan rows: %v", err)
	}

	return "Added in " + car.Make + " " + car.Model + " " + car.Horsepower + " " + car.MSRP
}
