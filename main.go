package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	webhost := os.Getenv("Hostname")

	helloHandler := func(write http.ResponseWriter, response *http.Request) {
		io.WriteString(write, " Welcome to AWK Web")
	}
	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at " + webhost)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
