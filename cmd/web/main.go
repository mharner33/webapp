package main

import (
	"fmt"
	"net/http"

	"github.com/mharner33/webapp/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting webserver on port: %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
