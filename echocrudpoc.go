package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JamesHovious/echocrudpoc/database"
	"github.com/JamesHovious/echocrudpoc/routes"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

const db = "poc.gob"

func main() {

	// Setup
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// Load the database if it exists
	if _, err := os.Stat(db); !os.IsNotExist(err) {
		if err != nil {
			e.Logger.Info("[-] Could not load a local gobdb")
		}
		err = database.LoadDB(db, &database.GobDB)
		if err != nil {
			e.Logger.Fatal(err)
		}
		e.Logger.Info("[+] Loaded gobdb")
	}

	// Set up the routes
	e.POST("/users", routes.CreateUser)
	e.GET("/users/:username", routes.GetUser)
	e.PUT("/users/:username", routes.UpdateUser)
	e.DELETE("/users/:username", routes.DeleteUser)
	e.GET("/database", routes.ShowDatabase)

	// Wait for interrupt signal to gracefully shutdown the server with
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("[*] Saving the database")
		err := database.SaveDB(db, &database.GobDB)
		if err != nil {
			e.Logger.Error("[-] Could not save database")
			e.Logger.Fatal(err)
		}
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
