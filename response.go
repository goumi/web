package goumi

import (
	"bufio"
	"bytes"
	"net"
	"net/http"
)

// Setup response interface
type ResponseWriter interface {

	// Extends ResponseWriter
	http.ResponseWriter

	// Extend Flusher, CloseNotify and Hijacker
	http.Flusher
	http.CloseNotifier
	http.Hijacker

	// Setup extra functions
	StatusCode() int
	Body() string
	ContentLength() int
}

// ResponseWriter - HTTP Response Writer that signals when the Write function is called
type response struct {

	// Extend http ResponseWriter
	http.ResponseWriter

	// Private data, it is exposed throught functions
	code        int
	bodyBuffer  *bytes.Buffer
	wroteHeader bool
}

// NewResponseWriter - Create a new response writer
func newResponse(w http.ResponseWriter) ResponseWriter {

	return &response{

		// Pass the previous response writer
		ResponseWriter: w,

		// Private
		code:        http.StatusInternalServerError,
		bodyBuffer:  new(bytes.Buffer),
		wroteHeader: false,
	}
}

// WriteHeader - Set the status code and save it
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

// Provided in order to implement the http.ResponseWriter interface.
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

// Flush - Provided in order to implement the http.Flusher interface.
func (w *response) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

// CloseNotify - Provided in order to implement the http.CloseNotifier interface.
func (w *response) CloseNotify() <-chan bool {
	return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

// Hijack - Provide the hijacker interface
func (w *response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

// StatusCode - Return the status code of the response
func (w *response) StatusCode() int {
	return w.code
}

// Body - Return the wrote body length
func (w *response) Body() string {

	// Return an empty string if no content length
	if w.ContentLength() == 0 {
		return ""
	}

	// Return a string containing the body
	return w.bodyBuffer.String()
}

// ContentLength - Return the wrote body length
func (w *response) ContentLength() int {
	return w.bodyBuffer.Len()
}
