package main

import (
	"fmt"
	"noahcarniglia/bookings/internal/config"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing, test passed

	default:
		t.Error(fmt.Sprintf("type is not *chi.mux, type is %T", v))
	}
}
