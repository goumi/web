package web

import "net/http"

// Handler is a addapter of net/http's http.Handler to use a single extensible
// context, instead of response writer and request.
type Handler interface {
	Serve(Context)
}

// HandlerFunc converts a function into a Handler
type HandlerFunc func(ctx Context)

// Serve calls the parent function
func (fn HandlerFunc) Serve(ctx Context) {
	fn(ctx)
}

// HTTPHandler extends http.Handler to act as a Handler
type httpHandler struct {
	http.Handler
}

// HTTPHandler creates a new Handler from a http.Handler
func HTTPHandler(h http.Handler) Handler {
	return &httpHandler{
		Handler: h,
	}
}

// Serve runs the http.Handler ServeHTTP then goes to the next middleware.
func (h *httpHandler) Serve(ctx Context) {

	// Serve the handler
	h.ServeHTTP(ctx.Response(), ctx.Request())

	// Load the next context
	ctx.Next()
}
