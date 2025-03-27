package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Println("Startiing server at port 80")
	if err := http.ListenAndServe(":8180", nil); err != nil {
		log.Fatal("Error in listening", err)

	}

}
