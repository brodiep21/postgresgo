package posql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

func (a *App) GetCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		HttpErrorResponse(w, http.StatusBadRequest, "Primary car ID not found")
	}

	c := Car{ID: id}
	if err := c.GetCar(a.DB); err != nil {
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

func HttpErrorResponse(w http.ResponseWriter, Rcode int, message string) {
	JsonResponse(w, Rcode, map[string]string{"error": message})
}

func JsonResponse(w http.ResponseWriter, Rcode int, info interface{}) {
	response, _ := json.Marshal(info)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(Rcode)
	w.Write(response)
}
