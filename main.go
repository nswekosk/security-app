package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/goteleport-interview/fs4/api"
)

const listenPort = 8080

//go:embed web/build
var assets embed.FS

func main() {
	webassets, err := fs.Sub(assets, "web/build")
	if err != nil {
		log.Fatalln("could not embed webassets", err)
	}

	s, err := api.NewServer(webassets)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO: Extend ability to reference TLS cert and key from a secrets solution
	// TODO: Obtain cert and key from a certificate authority
	log.Fatalln(http.ListenAndServeTLS(fmt.Sprintf("localhost:%d", listenPort), "server.crt", "server.key", s))
}
