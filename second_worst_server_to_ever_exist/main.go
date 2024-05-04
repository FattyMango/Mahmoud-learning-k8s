package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/google/uuid"
)

var (
	logger = log.New(os.Stdout, "http: ", log.LstdFlags)
	UNIQUE_ID = uuid.New().String()
)

func main() {
	PORT_ENV := os.Getenv("PORT")
	if PORT_ENV == "" {
		panic("PORT environment variable is required")
	}
	PORT, err := strconv.Atoi(PORT_ENV)
	if err != nil {
		panic("PORT environment variable must be a number")
	}

	app := http.NewServeMux()
	app.HandleFunc("/", handler)

	fmt.Println("Server is running on port", PORT)
	err = http.ListenAndServe(fmt.Sprintf(":%d", PORT), app)
	if err != nil {
		fmt.Println("Server stopped with error", err)
	}

	fmt.Println("Server is shutting down")
}

func handler(w http.ResponseWriter, r *http.Request) {
	logger.Println("Request received, by", UNIQUE_ID)
	w.Write([]byte("Hello World from " + UNIQUE_ID))
	logger.Println("Request completed, by", UNIQUE_ID)
}
