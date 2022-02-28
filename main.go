package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/list", listImages())
	http.HandleFunc("/add", addImage())

	fmt.Printf("Starting server at port %d\n", port)
	if err := http.ListenAndServe(":"+fmt.Sprintf("%d", port), nil); err != nil {
		log.Fatal(err)
	}

}
