package main

import (
    "log"
    "net/http"
)

func rootHander(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Header().Set("Content-Type", "text/html; charset=utf8")
    w.Write([]byte("こんにちは"))
}

func main() {
    http.HandleFunc("/", rootHander)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
