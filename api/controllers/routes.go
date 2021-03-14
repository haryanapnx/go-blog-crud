package controllers

import "github.com/haryanapnx/go-blog-crud/api/middlewares"

func (s *Server) initializeRoutes() {

	// HEALTH CHECK
	s.Router.HandleFunc("/health-check", middlewares.SetMiddlewareJSON(s.HealthCheck)).Methods("GET")

	// LOGIN
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//USER
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON((s.GetAllUser))).Methods("GET")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatedUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.RemoveUser)).Methods("DELETE")

	// ARTICLE
	s.Router.HandleFunc("/articles", middlewares.SetMiddlewareJSON(s.GetAllArticle)).Methods("GET")
	s.Router.HandleFunc("/articles", middlewares.SetMiddlewareJSON(s.CreateArticle)).Methods("POST")
	s.Router.HandleFunc("/articles/{id}", middlewares.SetMiddlewareJSON(s.GetArticle)).Methods("GET")
	s.Router.HandleFunc("/articles/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateArticle))).Methods("PUT")
	s.Router.HandleFunc("/articles/{id}", middlewares.SetMiddlewareJSON(s.RemoveArticle)).Methods("DELETE")
}
