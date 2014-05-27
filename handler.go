package ape

// Serve a context
type Handler interface {
	Serve(ctx *Context)
}

// Just build a handler
type HandlerFunc func(ctx *Context)

// Serve calls f(ctx)
func (f HandlerFunc) Serve(ctx *Context) {
	f(ctx)
}
