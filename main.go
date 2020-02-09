package main

import (
    "os"
    "fmt"
    "net/http"
)
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        hostname, err := os.Hostname()
        if err != nil {
            fmt.Fprint(w, err)
        }
        fmt.Fprint(w, hostname)
    })
    http.ListenAndServe(":9191", nil)
}
