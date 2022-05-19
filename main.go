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

	fmt.Printf("here is your data. \n"+fullvehicle+" %s %s \n", hp, msrp)
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

}
