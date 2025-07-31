package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/melkeydev/femProject/internal/app"
)

func main() {
	app, err := app.NewApplication()
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
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

// Additionally, the http.Request struct is large, so passing it by pointer is more efficient. Handlers and middleware may want to modify the request and those changes should persist for the full life of the request.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available \n")
}
