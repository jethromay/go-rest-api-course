package main

import "fmt"

// App - Struct which contains pointers and other things like db connections.
type App struct {
}

// Run - Sets up application
func (app *App) Run() error {
	fmt.Println("Setting up our APP")
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
