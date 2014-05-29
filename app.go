package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var (
		listen  = flag.String("listen", ":7000", "Http listen")
		rootDir = flag.String("root", "./var", "File root")
	)
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*rootDir)))

	log.Printf("Start HTTP Server listening on %q", *listen)
	if err := http.ListenAndServe(*listen, nil); err != nil {
		log.Fatalf("HTTP Server crashed: %q", err)
	}
}
