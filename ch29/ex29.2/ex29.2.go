package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func userHandler(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "jini"
	}
	age, _ := strconv.Atoi(values.Get("age"))
	fmt.Fprintf(writer, "name: %s, age: %d", name, age)
}
func main() {
	http.HandleFunc("/test", userHandler)
	http.ListenAndServe(":3000", nil)
}
