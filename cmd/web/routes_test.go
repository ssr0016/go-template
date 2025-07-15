package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/ssr0016/booking/internal/config"
)

func TestRoutes(t *testing.T) {
	var appConf config.AppConfig

	mux := routes(&appConf)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, but is %T", v))
	}
}
