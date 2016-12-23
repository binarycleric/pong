package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func pong(res http.ResponseWriter, req *http.Request) {
	fmt.Println(time.Now().Format(time.RFC3339) + " :: pong")
	fmt.Fprintln(res, "pong")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.SetOutput(os.Stdout)
	http.HandleFunc("/", pong)
	fmt.Println("listening...")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}

}
