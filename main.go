package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Hello World!")
}

func main() {
    http.HandleFunc("/", hello)
    if err := http.ListenAndServe(":8111", nil); err != nil {
        panic(err)
    }
}