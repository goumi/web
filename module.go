package goumi

import "net/http"

// Application handler with middleware attached
type Module []Handler

// New app
func New() *Module {
	return &Module{}
}

// Entry point into the applicaton
func (m *Module) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Create the context
	ctx := newContext(w, r)

	// Setup a new app
	m.Serve(ctx)
}

// Serve - Entry point into the applicaton
func (m *Module) Serve(ctx Context) {

	// Add the data to the chain and serve the next middleware
	ctx.Push(*m)

	// Run it now
	ctx.Next()
}

// Use - Add handlers to the middleware stack
func (m *Module) Use(h Handler) {
	*m = append(*m, h)
}
