// the responsibility of this
package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &Application{
		Logger: logger,
	}
	return app, nil
}

// Additionally, the http.Request struct is large, so passing it by pointer is more efficient. Handlers and middleware may want to modify the request and those changes should persist for the full life of the request.
func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available \n")
}
