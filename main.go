package main

import (
	"fav-imgs/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/list", server.ListImages())
	http.HandleFunc("/add", server.AddImage())
	http.HandleFunc("/delete", server.DeleteImage())
	http.HandleFunc("/update", server.UpdateImage())
	http.HandleFunc("/", server.ListImages())

	fmt.Printf("Starting server at port %d\n", port)
	if err := http.ListenAndServe(":"+fmt.Sprintf("%d", port), nil); err != nil {
		log.Fatal(err)
	}

}
