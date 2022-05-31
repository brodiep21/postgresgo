package posql

import (
	"database/sql"
	"fmt"
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

func (c *Car) getCar(db *sql.DB) error {
	return db.QueryRow("SELECT Make, Model, Horsepower, MSRP FROM cars WHERE id=$1", c.ID).Scan(&c.Make, &c.Model, &c.Horsepower, &c.MSRP)
}

func (c *Car) updateCar(db *sql.DB) (string, error) {

	result, err := db.Exec("UPDATE cars SET Make=$1, Model=$2, Horsepower=$3, MSRP=$4 WHERE id=$5", &c.Make, &c.Model, &c.Horsepower, &c.MSRP, &c.ID)
	if err != nil {
		return "", err
	}

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

	return "Changed " + strconv.Itoa(c.ID) + " to " + car.Make + " " + car.Model + " " + car.Horsepower + " " + car.MSRP, nil
}

func (c *Car) deleteCar(db *sql.DB) (string, error) {

	car := Car{}
	res := db.QueryRow("SELECT Make, Model, Horsepower, MSRP FROM cars WHERE id=$1", c.ID)
	if err := res.Scan(&car.Make, &car.Model, &car.Horsepower, &car.MSRP); err != nil {
		return "", nil
	}

	_, err := db.Exec("DELETE from cars WHERE id=$1", c.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Deleted %s %s, %s, %s", car.Make, car.Model, car.Horsepower, car.MSRP), nil
}

func (c *Car) createCar(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO cars (Make, Model, Horsepower, MSRP) VALUES ($1, $2, $3, $4)", &c.Make, &c.Model, &c.Horsepower, &c.MSRP)
	if err != nil {
		return err
	}
	return err
}

func (c *Car) getCars(db *sql.DB, start, counter int) ([]Car, error) {
	rows, err := db.Query("SELECT id, Make, Model, Horsepower, MSRP FROM cars LIMIT $1 OFFSET $2", counter, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cars := []Car{}

	for rows.Next() {
		var c Car
		if err := rows.Scan(&c.ID, &c.Make, &c.Model, &c.Horsepower, &c.MSRP); err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}
