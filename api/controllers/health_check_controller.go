package controllers

import (
	"net/http"

	"github.com/haryanapnx/go-blog-crud/api/responses"
)

func (server *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Everything is ok: 200")
}
