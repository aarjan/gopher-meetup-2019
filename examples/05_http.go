/*
	A simple program that runs a server at port :8080.
	It outputs 'hello' at the url address, localhost:8080/home
*/

package main

import "net/http"
import "fmt"

func main() {
	http.HandleFunc("/home", sayHello)
	http.ListenAndServe(":8080", nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
