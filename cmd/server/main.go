package main

import (
	"fmt"
	"github.com/valdirmendesdev/go-rest-api/internal/database"
	"github.com/valdirmendesdev/go-rest-api/internal/services"
	"net/http"

	transportHTTP "github.com/valdirmendesdev/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like
// pointers to database connections
type App struct{}

// Run - handles the startup of the application
func (app *App) Run() error {
	fmt.Println("Setting up the App")

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := services.NewCommentService(db)

	handler := transportHTTP.NewHandler(commentService)
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
