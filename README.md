# Goumi

Goumi is a very simple way to sandboxed modular applications. It improves on the `net/http` and tries to be as minimal as possible. It is inspired by [KoaJS](http://koajs.com/).

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file. We'll call it `server.go`.

~~~ go
package main

import (
  "fmt"
  "github.com/goumi/web"
  "github.com/goumi/logger"
  "github.com/goumi/mount"
  "net/http"
)

func main() {

  // Setup the app
  m := web.New()

  // Logger
  m.Use(logger.New())

  // Router
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to Goumi!")
  })
  m.Use(web.HTTPHandler(mux))

  // Mount another application
  mw := web.New()

  // Run another middleware
  mw.Use(web.HandlerFunc(func(ctx Context) {

    // Do your stuff
    ctx.Response().Header().Add("X-Powered-By", "Goumi")

    // Call next middlware
    ctx.Next()

    // You can do something else after the stack has run
    ctx.Response().Write([]byte("Hello!"))
  }))

  // Mount it on hello
  m.Use(mount.New("/hello", mw))

  // Run the server
  http.ListenAndServe(":3000", m)
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

- Docs
- Tests
- Errors
- Recovery

## Midleware

- Logger
- Server
- Static
- Router
- Session
- Templates (Pongo2)
