package main

import (
	"net/http"
	"io"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, go!\n")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}