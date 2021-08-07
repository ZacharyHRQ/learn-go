package main

import (
	"flag"
	"net/http"
	"log"
)

func main() {
	port := flag.String("p", "8000", "port")
	dir := flag.String("d", ".", "dir")
	flag.Parse()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./cat.png")
	})

	log.Printf("Serving %s on Http port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
