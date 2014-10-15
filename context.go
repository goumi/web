package web

import "net/http"

// Context is a per request object which will be passed arround to all middleware.
// It is meant to be extended in order to provide more functionality.
type Context interface {

	// Readers for request and response writer
	Request() *http.Request
	Response() ResponseWriter

	// Iterator
	Next()
}

// context is the most simple implementation of the Context interface
type context struct {
	res ResponseWriter
	req *http.Request
}

// newContext() sets up new context from the response writer and request
func newContext(w http.ResponseWriter, r *http.Request) Context {
	return &context{
		res: newResponse(w),
		req: r,
	}
}

// Response() gets the response from the context
func (ctx *context) Response() ResponseWriter {
	return ctx.res
}

// Request() gets the request from the context
func (ctx *context) Request() *http.Request {
	return ctx.req
}

// Next() is not implemented by default
func (ctx *context) Next() {}
