package main

// app.go

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App Data Structure to define our router and database
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize method to initialize the Database connection utilizing a connection string and the sql driver to connect
// this will also create the necessary new mux router to register the proper routes
func (a *App) Initialize(user, password, dbName string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbName=%s", user, password, dbName)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run method to start the application
func (a *App) Run(addr string) {}
