package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
)

func main() {
    li, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Println(err)
    }
    defer li.Close()

    for {
        conn, err := li.Accept()
        if err != nil {
            log.Println(err)
        }

        go handle(conn)
    }
}

func handle(conn net.Conn) {
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ln := scanner.Text()
        bs := []byte(ln)
        r := rot13(bs)

        fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
    }
    defer conn.Close()
}

func rot13(bs []byte) string {
    // A-Z - 65-90
    // a-z - 97-122
    var tb []byte
    var nbs byte
    for i := 0; i < len(bs); i++ {
        nbs = bs[i] + 13
        if nbs > 90 && nbs < 97 {
            nbs -= 25
        } else if nbs > 122 {
            nbs -= 25
        } else if bs[i] == 32 {
            nbs = 32
        }
        tb = append(tb, nbs)
    }

    return string(tb)
}
