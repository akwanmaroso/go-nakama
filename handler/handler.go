package handler

import (
	"net/http"

	"github.com/akwanmaroso/nakama-app/internal/service"
	"github.com/matryer/way"
)

type handler struct {
	*service.Service
}

// New create an http.Handler with predefined routing
func New(s *service.Service) http.Handler {
	h := &handler{s}
	api := way.NewRouter()
	api.HandleFunc("POST", "/login", h.login)
	api.HandleFunc("GET", "/auth_user", h.authUser)
	api.HandleFunc("POST", "/users", h.createUser)

	r := way.NewRouter()
	r.Handle("*", "/api...", http.StripPrefix("/api", h.withAuth(api)))

	return r
}
