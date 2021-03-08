package controllers

import "github.com/haryanapnx/go-blog-crud/api/middlewares"

func (s *Server) initializeRoutes() {

	// HEALTH CHECK
	s.Router.HandleFunc("/health-check", middlewares.SetMiddlewareJSON(s.HealthCheck)).Methods("GET")
}
