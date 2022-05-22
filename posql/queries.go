package posql

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

type Car struct {
	ID         int    `json:"id"`
	Make       string `json:"Make"`
	Model      string `json:"Model"`
	Horsepower string `json:"Horsepower"`
	MSRP       string `json:"MSRP"`
}

// func (c *Car) TableInsert(db *sql.DB) (string, error) {

// 	result, err := db.Exec("INSERT INTO cars (make, model, horsepower, msrp) VALUES ($1, $2, $3, $4)", &c.Make, &c.Model, &c.Horsepower, &c.MSRP)
// 	if err != nil {
// 		return "", err
// 	}

// 	changedRows, err := result.RowsAffected()
// 	if err != nil {
// 		return "", err
// 	}
// 	fmt.Println("inserted", changedRows, "rows")

// 	car := Car{}
// 	row := db.QueryRow("SELECT Make, Model, Horsepower, MSRP FROM cars ")

// 	if err := row.Scan(&car.Make, &car.Model, &car.Horsepower, &car.MSRP); err != nil {
// 		return "", err
// 	}

// 	return "Added in " + car.Make + " " + car.Model + " " + car.Horsepower + " " + car.MSRP, nil
// }

func (c *Car) GetCar(db *sql.DB) error {
	return db.QueryRow("SELECT make, model, MSRP, horsepower FROM cars WHERE id =$1", c.ID).Scan(&c.Make, &c.Model, &c.MSRP, &c.Horsepower)
}

func (c *Car) UpdateCar(db *sql.DB) (string, error) {

	result, err := db.Exec("UPDATE cars SET Make=$1, Model=$2, MSRP=$3, Horsepower=$4 WHERE id=$5", &c.Make, &c.Model, &c.MSRP, &c.Horsepower, &c.ID)

	_, err = result.RowsAffected()
	if err != nil {
		return "", err
	}

	car := Car{}
	idconv := strconv.Itoa(c.ID)
	if err != nil {
		return "", err
	}
	row := db.QueryRow("SELECT Make, Model, Horsepower, MSRP FROM cars WHERE id=$1", idconv)

	if err := row.Scan(&car.Make, &car.Model, &car.Horsepower, &car.MSRP); err != nil {
		return "", err
	}

	return "Added in " + car.Make + " " + car.Model + " " + car.Horsepower + " " + car.MSRP, nil
}

func (c *Car) DeleteCar(db *sql.DB) error {
	result, err := db.Exec("DELETE from cars WHERE id=$1", c.ID)
	return err
}

func (c *Car) CreateProduct(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO cars (make, model, horsepower, msrp) VALUES ($1, $2, $3, $4)", &c.Make, &c.Model, &c.Horsepower, &c.MSRP)
	if err != nil {
		return err
	}
	return err
}
