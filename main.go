package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	key := flag.String("key", "", "a key for a tls key/cert pair")
	cert := flag.String("cert", "", "a cert for a tls key/cert pair")
	listen := flag.String("listen", ":51004", "address to listen on for http connections")
	listens := flag.String("https-listen", ":51005", "address to listen on for https connections")
	flag.Parse()

	http.HandleFunc("/", Handler)
	go func() {
		if *listen == "" {
			log.Print("no --listen, disabling http")
			return
		}
		log.Print("HTTP Listening on " + *listen)
		panic(http.ListenAndServe(*listen, nil))
	}()
	go func() {
		if *listens == "" {
			log.Print("no --https-listen, disabling https")
			return
		}
		if *key == "" || *cert == "" {
			log.Print("missing --key or --cert, disabling https")
			return
		}
		log.Print("HTTPS Listening on " + *listens)
		panic(http.ListenAndServeTLS(*listens, *cert, *key, nil))
	}()
	select {}
}
