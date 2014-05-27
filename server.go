package ape

import "net/http"

type Server struct {
	http.Handler

	// Load the app
	*Application
}

func NewServer(app *Application) *Server {
	return &Server{
		Application: app,
	}
}

// Entry point into the applicaton
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Create the context
	ctx := NewContext(w, r)

	// Setup a new app
	s.Application.Serve(ctx)
}

// Run the server
func (s *Server) Run(addr string) {
	http.ListenAndServe(addr, s)
}
