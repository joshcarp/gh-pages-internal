package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joshcarp/gh-pages-internal/pkg/pages"
)

func main() {
	port := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	}
	server := pages.New()
	log.Println("Serving on "+port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(lis, &server))
}

