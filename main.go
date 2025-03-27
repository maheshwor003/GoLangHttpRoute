package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Error in Parsing Form", err)
	}
	fmt.Fprintf(w, "Post request Successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name= %s\n", name)
	fmt.Fprintf(w, "Address= %s\n", address)

}
func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	fmt.Println("Startiing server at port 80")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error in listening", err)
	}

}
