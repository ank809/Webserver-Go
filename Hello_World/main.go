package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "URL NOT FOUND", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Hello!")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/data" {
		http.Error(w, "URL NOT FOUND", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "METHOD NOT ALLOWED", http.StatusBadRequest)
		return
	}
	queryparams := r.URL.Query()
	param1 := queryparams.Get("param1")
	param2 := queryparams.Get("param2")

	fmt.Fprintf(w, "Param 1 %s Param2 %s", param1, param2)

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	fmt.Printf("server starting on port 8081\n")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/data", dataHandler)
	http.ListenAndServe(":8081", nil)
}
