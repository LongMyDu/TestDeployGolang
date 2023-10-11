package controllers

import (
	"net/http"
	"sample/api/middlewares"
)

// func (server *Server) initializeRoutes() {
// 	// Home
// 	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")

// 	// Login
// 	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

// 	// Users
// server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")
// 	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")
// 	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods("GET")
// 	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.UpdateUser))).Methods("PUT")
// 	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuth(server.DeleteUser)).Methods("DELETE")

// 	// Posts
// 	server.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(server.CreatePost)).Methods("POST")
// 	server.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(server.GetPosts)).Methods("GET")
// 	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(server.GetPost)).Methods("GET")
// 	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuth(middlewares.SetMiddlewareJSON(server.UpdatePost))).Methods("PUT")
// 	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuth(server.DeletePost)).Methods("DELETE")
// }

func (server *Server) initializeRoutes() {
	// Home
	http.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home))
	// Login
	// http.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

	// // Users
	// http.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")
	// http.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")
	// http.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods("GET")
	// http.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.UpdateUser))).Methods("PUT")
	// http.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuth(server.DeleteUser)).Methods("DELETE")

	// // Posts
	// http.HandleFunc("/posts", middlewares.SetMiddlewareJSON(server.CreatePost)).Methods("POST")
	// http.HandleFunc("/posts", middlewares.SetMiddlewareJSON(server.GetPosts)).Methods("GET")
	// http.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(server.GetPost)).Methods("GET")
	// http.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuth(middlewares.SetMiddlewareJSON(server.UpdatePost))).Methods("PUT")
	// http.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuth(server.DeletePost)).Methods("DELETE")

}
