package controllers

import "github.com/haryanapnx/go-blog-crud/api/middlewares"

func (s *Server) initializeRoutes() {

	// HEALTH CHECK
	s.Router.HandleFunc("/health-check", middlewares.SetMiddlewareJSON(s.HealthCheck)).Methods("GET")

	//USER
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON((s.GetAllUser))).Methods("GET")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatedUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.RemoveUser)).Methods("DELETE")
}
