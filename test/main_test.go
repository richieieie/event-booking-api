package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/richieieie/event-booking/internal/router"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	r := router.NewGinRouter()
	req, _ := http.NewRequest("GET", "/health-check", nil)
	w := httptest.NewRecorder()

	// Call the route handler
	r.ServeHTTP(w, req)

	// Assert the response status and body
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"ok"}`, w.Body.String())
}
