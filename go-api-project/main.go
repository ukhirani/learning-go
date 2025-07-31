package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/melkeydev/femProject/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	app.Logger.Println("we are running our app")
	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

// Additionally, the http.Request struct is large, so passing it by pointer is more efficient. Handlers and middleware may want to modify the request and those changes should persist for the full life of the request.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available \n")
}
