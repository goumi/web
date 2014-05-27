package ape

import "net/http"

// Context is a wrapper f
type Context struct {
	Res *ResponseWriter
	Req *http.Request

	// Next
	chain []Handler
	index int
}

// Setup a new context
func NewContext(w http.ResponseWriter, r *http.Request) *Context {

	return &Context{
		Res: NewResponseWriter(w),
		Req: r,

		index: -1,
	}
}

// Append middleware to the context
func (ctx *Context) Push(chain []Handler) {
	ctx.chain = append(ctx.chain, chain[0:]...)
}

// Next function
func (ctx *Context) Next() {

	// Advance the index
	ctx.index += 1

	// Load the element from the chain
	if ctx.index >= len(ctx.chain) {
		return
	}

	// Run the Handler
	ctx.chain[ctx.index].Serve(ctx)
}
