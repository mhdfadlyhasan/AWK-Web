package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	webhost := os.Getenv("Hostname")

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, " Welcome to AWK Web")
	}
	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at %s \n", webhost)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
