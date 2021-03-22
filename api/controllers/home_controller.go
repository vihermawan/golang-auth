package controllers

import (
	"net/http"

	"github.com/vickyhermawan/rest-api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "API okey")
}
