package ape

import "net/http"

// Simple handler with middleware attached
type Application struct {
	mw []Handler
}

// Create a new app
func New() *Application {
	return &Application{}
}

// Entry point into the applicaton
func (app *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Create the context
	ctx := NewContext(w, r)

	// Setup a new app
	app.Serve(ctx)
}

// Entry point into the applicaton
func (app *Application) Serve(ctx *Context) {

	// Add the data to the chain and serve the next middleware
	ctx.Push(app.mw)

	// Run it now
	ctx.Next()
}

// Add handlers to the middleware stack
func (app *Application) Use(h Handler) {
	app.mw = append(app.mw, h)
}

// Add handlers to the middleware stack
func (app *Application) UseFunc(f HandlerFunc) {
	app.Use(HandlerFunc(f))
}

// Run the server
func (app *Application) Run(addr string) {
	http.ListenAndServe(addr, app)
}
