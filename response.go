package ape

import (
	"bufio"
	"bytes"
	"net"
	"net/http"
)

// HTTP Response Writer that signals when the Write function is called
type ResponseWriter struct {

	// Extend http ResponseWriter
	http.ResponseWriter

	// Publi exposed data
	Code    int
	Body    *bytes.Buffer
	Flushed bool

	// Have a flag to check if the response is writtern
	wroteHeader bool
}

// Create a new response writer
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{

		// Pass the previous response writer
		ResponseWriter: w,

		// Public
		Code:    http.StatusInternalServerError,
		Body:    new(bytes.Buffer),
		Flushed: false,

		// Private
		wroteHeader: false,
	}
}

// Set the status code and save it
func (w *ResponseWriter) WriteHeader(code int) {

	// Save the status code when it is writtern
	if !w.wroteHeader {
		w.Code = code
	}

	// Write the header
	w.ResponseWriter.WriteHeader(code)

	// Keep track that we wrote the header
	w.wroteHeader = true
}

// Provided in order to implement the http.ResponseWriter interface.
func (w *ResponseWriter) Write(b []byte) (int, error) {

	// Default to 200 HTTP Status
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}

	// Append data to the body and set the wrote flag
	w.Body.Write(b)

	// Done, return the response write function
	return w.ResponseWriter.Write(b)
}

// Provided in order to implement the http.Flusher interface.
func (w *ResponseWriter) Flush() {

	// Default to 200 HTTP Status
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}

	// Implement the Flush interface
	w.ResponseWriter.(http.Flusher).Flush()

	// Keep track if the data was flushed
	w.Flushed = true
}

// Provided in order to implement the http.CloseNotifier interface.
func (w *ResponseWriter) CloseNotify() <-chan bool {
	return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

// Provide the hijacker interface
func (w *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
