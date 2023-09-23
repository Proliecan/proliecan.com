package main

import (
	"log"
	"net/http"
	"strconv"
)

func initServer() {
	// handle all routes and branch out into handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(colorize(r.RemoteAddr, Yellow),
			colorize(">", Red), r.Method, r.URL.Path)

		switch r.URL.Path {
		case "/":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("🤌"))
			log.Println(colorize(r.RemoteAddr, Yellow),
				colorize("<", Green), "🤌")
		case "/favicon.ico":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(""))
			log.Println(colorize(r.RemoteAddr, Yellow),
				colorize("<", Green), "")
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404"))
			log.Println(colorize(r.RemoteAddr, Yellow),
				colorize("<", Green), "404")
		}
	})
}

func startServer(domain string, port int) {
	// start server
	log.Fatal(http.ListenAndServe(domain+":"+strconv.Itoa(port), nil))
}
