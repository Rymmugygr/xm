package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "Ok\n")
	})
	fmt.Println("Starting")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
