package main

import "net/http"

func main() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":3000", nil)
}
