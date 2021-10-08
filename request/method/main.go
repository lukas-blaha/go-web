package main

import (
    "html/template"
    "log"
    "net/http"
    "net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    err := req.ParseForm()
    if err != nil {
        log.Fatalln(err)
    }
    
    data := struct {
        Method          string
        URL             *url.URL
        Submissions     map[string][]string // url.Values
        Header          http.Header
        Host            string
        ContentLength   int64
    }{
        Method:         req.Method,
        URL:            req.URL,
        Submissions:    req.Form,
        Header:         req.Header,
        Host:           req.Host,
        ContentLength:  req.ContentLength,
    }

    tpl.ExecuteTemplate(w, "postForm.gohtml", data)
}

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("postForm.gohtml"))
}

func main() {
    var d hotdog
    http.ListenAndServe(":8080", d)
}
