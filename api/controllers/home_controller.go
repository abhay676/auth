package controllers

import (
	"github.com/abhay676/auth/api/responses"
	"net/http"
	"time"
)

func (s *Server) HeathCheck(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, struct {
		Service    string    `json:"service"`
		ServerTime time.Time `json:"server_time"`
	}{
		Service:    "running",
		ServerTime: time.Now(),
	})
}