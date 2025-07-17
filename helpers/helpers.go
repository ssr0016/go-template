package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/ssr0016/booking/internal/config"
)

var app *config.AppConfig

// NewHelpers sets the config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

// ClientError sends a status code
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

// ServerError sends a status code
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
