package main

import (
    "log"
    "os"
    "text/template"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl2.gohtml"))
}

func main() {
    sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

    err := tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", sages)
    if err != nil {
        log.Fatalln(err)
    }

}
