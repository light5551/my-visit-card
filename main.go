package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world2!")
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}