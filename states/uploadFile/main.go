package main

import (
    "fmt"
    "os"
    "io"
    "io/ioutil"
    "net/http"
    "path/filepath"
)

func foo(w http.ResponseWriter, req *http.Request) {
    var s string
    fmt.Println(req.Method)
    if req.Method == http.MethodPost {
        f, h, err := req.FormFile("q")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer f.Close()

        fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)
        
        bs, err := ioutil.ReadAll(f)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        s = string(bs)

        dst, err := os.Create(filepath.Join("./files/", h.Filename))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer dst.Close()

        _, err = dst.Write(bs)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    io.WriteString(w, `
    <form method="POST" enctype="multipart/form-data">
        <input type="file" name="q">
        <input type="submit">
    </form>
    <br>`+s)
}

func main() {
    http.HandleFunc("/", foo)
    http.Handle("/favicon.ico", http.NotFoundHandler())
    http.ListenAndServe(":8080", nil)
}
