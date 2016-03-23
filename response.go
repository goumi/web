package web

import (
	"bytes"
	"net/http"
)

// ResponseWriter is a wrapper around http.ResponseWriter that provides extra information about
// the response.
type ResponseWriter interface {

	// Extends ResponseWriter
	http.ResponseWriter

	// Setup extra functions
	StatusCode() int
	Body() string
	ContentLength() int
}

// response is a simple structure that extends the original http.ResponseWriter
type response struct {
	http.ResponseWriter

	// Store data about the response
	code        int
	bodyBuffer  *bytes.Buffer
	wroteHeader bool
}

// NewResponse creates a new ResponseWriter interface
func NewResponse(w http.ResponseWriter) ResponseWriter {

	return &response{
		ResponseWriter: w,

		// Private data
		code:        http.StatusInternalServerError,
		bodyBuffer:  new(bytes.Buffer),
		wroteHeader: false,
	}
}

// WriteHeader writes the header, no further modifications after this point,
// so it stores the response http code and keeps track that the header was wrote.
func (w *response) WriteHeader(code int) {

	// Don't write the header multiple times
	if w.wroteHeader {
		return
	}

	// Save the code
	w.code = code

	// Write the header
	w.ResponseWriter.WriteHeader(code)

	// Keep track that we wrote the header
	w.wroteHeader = true
}

// Write the response and keep track of it. Before writing will set the header
// if not set to 200.
func (w *response) Write(b []byte) (int, error) {

	// Default to 200 HTTP Status
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}

	// Append data to the body and set the wrote flag
	w.bodyBuffer.Write(b)

	// Done, return the response write function
	return w.ResponseWriter.Write(b)
}

// StatusCode will return the status code of the response
func (w *response) StatusCode() int {
	return w.code
}

// Body will return the response content
func (w *response) Body() string {

	// Return an empty string if no content length
	if w.ContentLength() == 0 {
		return ""
	}

	// Return a string containing the body
	return w.bodyBuffer.String()
}

// ContentLength will return the content length
func (w *response) ContentLength() int {
	return w.bodyBuffer.Len()
}
