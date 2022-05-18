package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/brodiep21/postgresgo/posql"
	"github.com/brodiep21/postgresgo/search"
	_ "github.com/lib/pq"
)

func main() {

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

	fmt.Println("Would you like to add this data into the table? Yes or No?")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()
	response = strings.ToLower(response)
	switch response {
	case "y", "yes":
		posql.TableInsert(make, model, hp, msrp)
	case "n", "no":
		break
	}

	// rows, err := db.Query("SELECT Make, Model, Horsepower, MSRP FROM cars Limit 10")
	// if err != nil {
	// 	log.Fatalf("couldn't execute query: %v", err)
	// }

}
