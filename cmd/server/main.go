package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/jethromay/go-rest-api-course/internal/transport/http"
)

// App - Struct which contains pointers and other things like db connections.
type App struct{}

// Run - Sets up application
func (app *App) Run() error {
	fmt.Println("Setting up our APP")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":9000", handler.Router); err != nil {
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
