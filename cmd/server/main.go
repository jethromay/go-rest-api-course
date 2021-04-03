package main

import (
	"fmt"
	"net/http"

	"github.com/jethromay/go-rest-api-course/internal/database"
	transportHTTP "github.com/jethromay/go-rest-api-course/internal/transport/http"
)

// App - Struct which contains pointers and other things like db connections.
type App struct{}

// Run - Sets up application
func (app *App) Run() error {
	fmt.Println("Setting up our APP")

	var err error
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting our REST API.")
		fmt.Println(err)
	}
}
