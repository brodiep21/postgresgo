package main

import (
	"log"
	"os"

	"github.com/brodiep21/postgresgo/posql"
	_ "github.com/lib/pq"
)

var password = os.Getenv("pass")

func main() {
	a := posql.App{}
	err := a.Initialize(password)
	if err != nil {
		log.Fatal(err)
	}
	err = a.Run("8080")
	if err != nil {
		log.Fatal(err)
	}

	// vehicle, err = MMquestions()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("here is your data. \n" + vehicle)
	// fmt.Println("Would you like to add this data into the table? Yes or No?")

	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// response := scanner.Text()
	// response = strings.ToLower(response)
	// switch response {
	// case "y", "yes":
	// 	posql.TableInsert(make, model, hp, msrp)
	// case "n", "no":
	// 	break
	// }

}
