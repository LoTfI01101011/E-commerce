package main

import (
	"net/http"

	"github.com/LoTfI01101011/E-commerce/Api_gateway/api/rest"
)

func main() {
	r := rest.Router()
	http.ListenAndServe(":8000", r)
}
