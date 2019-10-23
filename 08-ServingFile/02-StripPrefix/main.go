package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", test)
	http.Handle("/resources/", http.StripPrefix("/resources",http.FileServer(http.Dir("./test"))))
	http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/test.png">`)
}