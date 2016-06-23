package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"

)

func TestGetUsersOK(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("a")

	// Assertions
	if assert.NoError(t, GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "a", rec.Body.String())
	}
}

func TestGetUsersNotFound(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("nothing")

	// Assertions
	if assert.NoError(t, GetUsers(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "user not found", rec.Body.String())
	}

}