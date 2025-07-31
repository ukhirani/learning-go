package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/melkeydev/femProject/internal/app"
	"github.com/melkeydev/femProject/internal/routes"
)

func main() {
	app, err := app.NewApplication()
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()
	if err != nil {
		panic(err)
	}
	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Println("we are running our app", port)
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
