package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hellow from api! 23 ")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println()
	log.Fatal(http.ListenAndServe(":5000", nil))
}