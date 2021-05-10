package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/valdirmendesdev/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like
// pointers to database connections
type App struct{}

// Run - handles the startup of the application
func (app *App) Run() error {
	fmt.Println("Setting up the App")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up the server")
		return err
	}

	return nil
}

// Main entrypoint for the application
func main() {
	app := App{}
	fmt.Println("Go REST API")
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up the REST API")
		fmt.Println(err)
	}
}
