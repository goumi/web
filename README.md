# Goumi

Goumi is a very simple way to setup middleware, that improves on the `net/http` Handler interface.

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file. We'll call it `server.go`.

~~~ go
package main

import (
  "fmt"
  "github.com/claudiuandrei/goumi"
  "net/http"
)

func main() {

  // Setup the app
  g := goumi.New()

  // Setup a middleware
  g.Use(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("X-Powered-By", "Gibbon")
  }))

  // Load the server
  g.Use(goumi.HandlerFunc(func(ctx goumi.Context) {
    ctx.Response().Header().Add("X-Powered-By", "Goumi")

    // Call next middleware
    ctx.Next()

    // Do stuff after the chain ends
  }))

  // Setup the router
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to Goumi!")
  })
  g.Use(goumi.HTTPHandler(mux));

  // Run the server
  http.ListenAndServe(":3000", g)
}
~~~

Then install the Goumi package (**go 1.3** and greater is required):
~~~
go get github.com/claudiuandrei/goumi
~~~

Then run your server:
~~~
go run server.go
~~~

You will now have a Go net/http webserver running on `localhost:3000`.

## `Use()` - Middleware & Routing

Goumi `Use` function adds middleware to the request execution chain. Everything that acts like a Handler can be a middleware and will run through each of them in the order added as long as there context Next() function is called.

Goumi is BYOR (Bring your own Router). Goumi is fully supporting `net/http`.

## Authors

[Claudiu Andrei](http://claudiuandrei.com/)

## TODO

- Router
- Logger
- Server (with logging)
- Static file serving
- Environment
- Docs
- Tests
- View Layer
- Template / Pongo2 (React)
- Errors
- Session
