package posql

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
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

func (a *App) Initialize() {

	var password = os.Getenv("pass")

	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	a.DB.SetMaxIdleConns(5)
	a.DB.SetMaxOpenConns(7)
	a.DB.SetConnMaxIdleTime(4 * time.Second)
	a.DB.SetConnMaxLifetime(20 * time.Second)

	if err := a.DB.Ping(); err != nil {
		log.Fatalf("unable to reach db: %v", err)
	}
	fmt.Println("Reached DB")

	a.Router = mux.NewRouter()
}

func (a *App) Run(address string) {
	err := http.ListenAndServe(":"+address, a.Router)
	if err != nil {
		log.Fatalf("Could not Run the Router, %s", err)
	}
}
