package main

import (
    "github.com/satori/go.uuid"
    "html/template"
    "net/http"
)

type user struct {
    UserName    string
    First       string
    Last        string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
    tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
    http.HandleFunc("/", foo)
    http.Handle("/favicon.ico", http.NotFoundHandler())
    http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
    cookie, err := req.Cookie("session")
    if err != nil {
        sID, _ := uuid.NewV4()
        cookie = &http.Cookie{
            Name:       "session",
            Value:      sID.String(),
            HttpOnly:   true,
        }
        http.SetCookie(w, cookie)
    }
    var u user
    if un, ok := dbSessions[cookie.Value]; ok {
        u = dbUsers[un]
    }

    if req.Method == http.MethodPost {
        un := req.FormValue("username")
        f := req.FormValue("firstname")
        l := req.FormValue("lastname")
        u = user{un, f, l}
        dbSessions[cookie.Value] = un
        dbUsers[un] = u
    }
    tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
    cookie, err := req.Cookie("session")
    if err != nil {
        http.Redirect(w, req, "/", http.StatusSeeOther)
        return
    }
    un, ok := dbSessions[cookie.Value]
    if !ok {
        http.Redirect(w, req, "/", http.StatusSeeOther)
        return
    }
    u := dbUsers[un]
    tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
