package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "net"
    "strings"
)

func handle(conn net.Conn) {
    defer conn.Close()

    m, u := request(conn)
    respond(conn, m, u)
}

func request(conn net.Conn) (string, string){
    var counter int

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        txt := scanner.Text()
        if txt != "" {
            if counter == 0 {
                sl := strings.Fields(txt)
                // m := fmt.Sprintf("Method: %s\n", sl[0])
                // u := fmt.Sprintf("URI: %s\n", sl[1])
                return sl[0], sl[1]
            }
        //} else {
        //    break
        }
        counter++
    }
    return "none", "none"
}

func createBody(txt string) string {
    body := fmt.Sprintf(`
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Code Gangsta</title>
    </head>
    <body>
        <h1>%s</h1>
        <a href="/">index</a></br>
        <a href="/dog">dog</a>
        <form action="/dog" method="post">
            <input type="hidden" value="">
            <input type="submit" value="submit">
        </form>
    </body>
    </html>
    `, txt)

    return body
}

func respond(conn net.Conn, m string, u string) {
    var body string

    if m == "GET" && u == "/" {
        body = createBody("This is a main page...")
    } else if m == "GET" && (u == "/dog" || u == "/dog/") {
        body = createBody("This is a dog page...")
    } else if m == "POST" && u == "/dog" {
        body = createBody("This is a dog page, method POST...")
    } else {
        body = createBody("Something else...")
    }


    io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
    fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    io.WriteString(conn, "\r\n")
    io.WriteString(conn, body)
}

func main() {
    li, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalln(err)
    }
    defer li.Close()

    for {
        conn, err := li.Accept()
        if err != nil {
            log.Fatalln(err)
        }

        go handle(conn)
    }
}
