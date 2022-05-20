package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/brodiep21/postgresgo/posql"
	_ "github.com/lib/pq"
)

func main() {

	a.Initialize()
	a.Run("8080")

	vehicle := MMquestions()

	fmt.Printf("here is your data. \n" + vehicle)
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
