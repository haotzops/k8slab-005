package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        hostname, _:= os.Hostname()

        message := os.Getenv("MESSAGE")
        if message == "" {
            message = "no message configured"
        }

        fmt.Fprintf(w, "hostname: %s\n", hostname)
        fmt.Fprintf(w, "message: %s\n", message)
    })

    log.Println("listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

