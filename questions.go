package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/brodiep21/postgresgo/search"
)

// requests make and model from cmdline
func MMquestions() (string, error) {

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

	hp, err := search.HorsepowerSearch(fullvehicle)
	if err != nil {
		return "", error
	}
	msrp, err = search.MsrpSearch(fullvehicle)
	if err != nil {
		return "", error
	}

}
