package main

import (
    "log"
    "os"
    "text/template"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("templates/tpl2.gohtml"))
}

func main() {
    err := tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", `Release self-focus; embrace other-focus.`)
    if err != nil {
        log.Fatalln(err)
    }

}
