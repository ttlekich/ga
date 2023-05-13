package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ttlekich/ga/api"
)

func main() {
	port := flag.String("PORT", ":1337", "the server listening port")

	server := api.NewServer(*port)

	fmt.Println("listening on port: ", *port)

	log.Fatal(server.Start())
}
