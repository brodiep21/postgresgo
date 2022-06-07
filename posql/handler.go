package posql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "bp21"
	dbname = "bp21"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(password string) error {

	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	a.DB.SetMaxIdleConns(5)
	a.DB.SetMaxOpenConns(7)
	a.DB.SetConnMaxIdleTime(4 * time.Second)
	a.DB.SetConnMaxLifetime(20 * time.Second)

	if err := a.DB.Ping(); err != nil {
		return err
	}
	fmt.Println("Reached DB")

	a.Router = mux.NewRouter()
	return nil
}

func (a *App) Run(address string) error {
	err := http.ListenAndServe(":"+address, a.Router)
	if err != nil {
		return err
	}
	return nil
}

func HttpErrorResponse(w http.ResponseWriter, Rcode int, message string) {
	JsonResponse(w, Rcode, map[string]string{"error": message})
}

func JsonResponse(w http.ResponseWriter, Rcode int, info interface{}) {
	response, _ := json.Marshal(info)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(Rcode)
	w.Write(response)
}

func (a *App) GetCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HttpErrorResponse(w, http.StatusBadRequest, "Primary car ID not found")
	}

	c := Car{ID: id}
	if err := c.getCar(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			HttpErrorResponse(w, http.StatusNotFound, "Car not found")
		default:
			HttpErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	JsonResponse(w, http.StatusOK, c)
}

func (a *App) GetCars(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	c := Car{}
	cars, err := c.getCars(a.DB, start, count)
	if err != nil {
		HttpErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	JsonResponse(w, http.StatusOK, cars)
}

func (a *App) CreateCar(w http.ResponseWriter, r *http.Request) {

	var c Car
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		HttpErrorResponse(w, http.StatusBadRequest, "Invalid request information")
		return
	}
	defer r.Body.Close()

	if err := c.createCar(a.DB); err != nil {
		HttpErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	JsonResponse(w, http.StatusCreated, c)
}

func (a *App) UpdateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HttpErrorResponse(w, http.StatusBadRequest, "Invalid Car ID")
		return
	}

	var c Car
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		HttpErrorResponse(w, http.StatusBadRequest, "Invalid request information")
		return
	}

	defer r.Body.Close()

	str, err := c.updateCar(a.DB)
	if err != nil {
		HttpErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	c = Car{
		ID:         id,
		Make:       str.Make,
		Model:      str.Model,
		Horsepower: str.Horsepower,
		MSRP:       str.MSRP,
	}

	JsonResponse(w, http.StatusOK, c)
}

func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HttpErrorResponse(w, http.StatusBadRequest, "Invalid Car ID")
		return
	}

	c := Car{ID: id}

	car, err := c.deleteCar(a.DB)
	if err != nil {
		HttpErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	c = Car{
		ID:         car.ID,
		Make:       car.Make,
		Model:      car.Model,
		Horsepower: car.Horsepower,
		MSRP:       car.MSRP,
	}

	JsonResponse(w, http.StatusOK, c)
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/cars", a.GetCars).Methods("GET")
	a.Router.HandleFunc("/car", a.CreateCar).Methods("POST")
	a.Router.HandleFunc("/car/{id:[0-9]+}", a.GetCar).Methods("GET")
	a.Router.HandleFunc("/car/{id:[0-9]+}", a.UpdateCar).Methods("PUT")
	a.Router.HandleFunc("/car/{id:[0-9]+}", a.DeleteProduct).Methods("DELETE")
}
