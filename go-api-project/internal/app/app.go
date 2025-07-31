// the responsibility of this
package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/melkeydev/femProject/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// our stores will go here

	// out handlers will go here

	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}
	return app, nil
}

// Additionally, the http.Request struct is large, so passing it by pointer is more efficient. Handlers and middleware may want to modify the request and those changes should persist for the full life of the request.
func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available \n")
}
