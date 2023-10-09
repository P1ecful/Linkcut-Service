package main

import (
	"fmt"
	"linkcut/internal/service"
	"linkcut/internal/transport"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var lcs service.LinkCutService = service.LinkcutService{}

	cutHandler := httptransport.NewServer(
		transport.MakeLinkCutEndpoint(lcs),
		transport.DecodeCutRequest,
		transport.EncodeResponse,
	)

	http.Handle("/cut", cutHandler)

	fmt.Println("Server in listening on: localhost:1200/cut")
	log.Fatal(http.ListenAndServe(":1200", nil))
}
