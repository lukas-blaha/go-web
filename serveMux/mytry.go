package main

import (
    "io"
    "net/http"
)

type about string

func (a about) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "This is an about page...")
}

type contact string

func (c contact) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "This is a contact page...")
}

func main () {
    var a about
    var c contact

    mux := http.NewServeMux()
    mux.Handle("/about", a)
    mux.Handle("/contact/", c)

    http.ListenAndServe(":8080", mux)
}
