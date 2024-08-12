package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//configure the songs directory name and port

	const songDir = "Files"
	const port = 8080

	//add handler for song files
	http.Handle("/", addHeader(http.FileServer(http.Dir(songDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port :%v\n", songDir, port)

	//server the log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func addHeader(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-control -Allow -Origin", "*")
		h.ServeHTTP(w, r)
	}
}
