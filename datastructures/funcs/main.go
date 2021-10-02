package main

import (
    "log"
    "os"
    "strings"
    "text/template"
)

var tpl *template.Template

type sage struct {
    Name    string
    Motto   string
}

type car struct {
    Manufacturer    string
    Model           string
    Doors           int
}

type items struct {
    Wisdom      []sage
    Transport   []car
}

var fm = template.FuncMap{
    "uc": strings.ToUpper,
    "ft": firstThree,
}

func firstThree(s string) string {
    s = strings.TrimSpace(s)
    s = s[:3]
    return s
}

func init() {
    tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
    buddha := sage {
        Name:   "Buddha",
        Motto:  "The belief of no beliefs",
    }

    mlk := sage {
        Name:   "Martin Luther King",
        Motto:  "Hatred never ceases with hatred but with love alone is healed",
    }

    jesus := sage {
        Name:   "Jesus",
        Motto:  "Love all",
    }

    ford := car {
        Manufacturer:   "Ford",
        Model:          "F150",
        Doors:          2,
    }

    toyota := car {
        Manufacturer:   "Toyota",
        Model:          "Corrola",
        Doors:          4,
    }

    sages := []sage{buddha, mlk, jesus}
    cars := []car{ford, toyota}

    data := items{
        Wisdom:     sages,
        Transport:  cars,
    }

    err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
    if err != nil {
        log.Fatalln(err)
    }

}
