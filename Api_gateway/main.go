package main

import (
	"log"
	"net/http"

	"github.com/LoTfI01101011/E-commerce/Api_gateway/api/gRPC"
	"github.com/LoTfI01101011/E-commerce/Api_gateway/api/rest"
)

func main() {
	log.Println("Starting API Gateway...")
	user := gRPC.User{}
	user.Start("localhost:9001") // Ensure correct gRPC server address
	defer user.Stop()
	r := rest.Router(&user)
	log.Println("startign the router")
	http.ListenAndServe(":8000", r)
	log.Println("startign the router")
}
