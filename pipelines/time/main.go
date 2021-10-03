package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    fmt.Println(t)
    // tf := t.Format(time.Kitchen)
    tf := t.Format("02-01-2006")
    fmt.Println(tf)
}
