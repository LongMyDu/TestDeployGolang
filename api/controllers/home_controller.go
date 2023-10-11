package controllers

import (
	"net/http"

	"sample/api/responses"
)

// Home is..
func (server *Server) Home(res http.ResponseWriter, req *http.Request) {
	responses.JSON(res, http.StatusOK, "Welcome to our API!")
}
