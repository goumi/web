package ape

// Simple handler with middleware attached
type Application struct {
	Handler
	middleware []Handler
}

// Entry point into the applicaton
func (app *Application) Serve(ctx *Context) {

	// Add the data to the chain and serve the next middleware
	ctx.Push(app.middleware)

	// Run it now
	ctx.Next()
}

// Create a new app
func New() *Application {
	return &Application{}
}

// Add handlers to the middleware stack
func (app *Application) Use(f HandlerFunc) {
	app.Mount(HandlerFunc(f))
}

// Add handlers to the middleware stack
func (app *Application) Mount(h Handler) {
	app.middleware = append(app.middleware, h)
}
