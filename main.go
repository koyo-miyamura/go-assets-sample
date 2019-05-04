package main

import "net/http"

//go:generate go-assets-builder -s="/data" -o bindata.go data

func main() {
	http.Handle("/", http.FileServer(Assets))
	http.ListenAndServe(":8080", nil)
}
