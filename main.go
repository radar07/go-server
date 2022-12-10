package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Hello!!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error %v", err)
	}

	name := r.FormValue("name")
	age := r.FormValue("age")

	fmt.Fprintf(w, "Hello, %s! I bet you're atleast %s", name, age)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting the server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error!!", err)
	}
}
