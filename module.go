/*
Package web implements a minimal and flexible middleware stack that builds on
go http package in a similar way that KoaJS leverages NodeJS.

It provides adapters for the native http.Handler's but it creates a more modular
way to build applications based on the decorator pattern. The purpose is to provide
a good way to use middlware to decorate the context (response - request wrapper).

A usage example:

	m := web.New()

	// Loaded from elsewhere
	var handler web.Handler
	var httphandler http.Handler

	// Handler
	m.Use(handler)

	// http.Handler
	m.Use(web.HTTPHandler(httphandler))

	// Handler function
	m.Use(web.HandlerFunc(func(ctx Context) {

		// Do your stuff

		// Call next middlware
		ctx.Next()

		// You can do something else after the stack has run
	}))

*/

package web

import "net/http"

// Module is a chain application handler
type Module []Handler

// New module
func New() *Module {
	return &Module{}
}

// ServeHTTP() serves as an entry point into the module, and translates the
// response writer and request into a single structure.
func (m *Module) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Create a context from the response and request
	ctx := newContext(w, r)

	// Serve the app using the new context
	m.Serve(ctx)
}

// Serve() extends the module with a new context, that holds the module
// middleware.
func (m *Module) Serve(ctx Context) {

	// Sandbox the context middleware
	ctx = newAppContext(ctx, m)

	// Run the middleware
	ctx.Next()
}

// Use() adds a Handler to the the module chain
func (m *Module) Use(h Handler) {
	*m = append(*m, h)
}

// appContext extends the context to access the module chain
type appContext struct {
	Context

	// Contain the middleware and the index for the current middleware
	chain []Handler
	index int
}

// newAppContext() creates a new module context from the previous context and
// the module middleware
func newAppContext(ctx Context, m *Module) Context {
	return &appContext{
		Context: ctx,
		chain:   *m,
		index:   -1,
	}
}

// Next() runs the next middleware
func (ctx *appContext) Next() {

	// Increment
	ctx.index++

	// Check if we have middleware in the current chain
	if ctx.index < len(ctx.chain) {

		// Serve the current handler
		ctx.chain[ctx.index].Serve(ctx)

		// Done
		return
	}

	// Exit current chain, advance to the next one
	ctx.Context.Next()
}
