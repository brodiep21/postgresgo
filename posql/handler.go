package posql

import (
	"net/http"
	"database/sql"
	"fmt"

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

	a.DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	
	a.Router = mux.NewRouter()
}

func (a *App) Run(address string) {
	err := http.ListenAndServe(":" + address, handler)
}
