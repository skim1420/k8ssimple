package main

import (
  "io"
  "net/http"
)

func main() {
  http.HandleFunc("/", hello)
  http.ListenAndServe(":80", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "<body style='background-color: dodgerblue'><h1>Hello simple!</h1></body>")
}

