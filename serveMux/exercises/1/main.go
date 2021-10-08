package main

import (
    "net/http"
    "html/template"
    "log"
)

func i(w http.ResponseWriter, req *http.Request) {
    tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
    err := tpl.Execute(w, "This is a main page...")
    if err != nil {
        log.Fatalln(err)
    }
}

func m(w http.ResponseWriter, req *http.Request) {
    tpl := template.Must(template.ParseFiles("templates/me.gohtml"))
    me := struct {
        FirstName   string
        LastName    string
    }{
        FirstName:  "Lukas",
        LastName:   "Blaha",
    }
    err := tpl.Execute(w, me)
    if err != nil {
        log.Fatalln(err)
    }
}

func d(w http.ResponseWriter, req *http.Request) {
    tpl := template.Must(template.ParseFiles("templates/dog.gohtml"))
    err := tpl.Execute(w, "This is a dog page...")
    if err != nil {
        log.Fatalln(err)
    }
}

func main () {
    http.Handle("/", http.HandlerFunc(i))
    http.Handle("/dog/", http.HandlerFunc(d))
    http.Handle("/me/", http.HandlerFunc(m))

    http.ListenAndServe(":8080", nil)
}
