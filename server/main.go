package main

import(
	"fmt"
	"log"
	"net/http"
	"./router"
)
	


func main() {
	r := router.Router()

	fmt.Println("Starting server in port localhost....")

	log.Fatal(http.ListenAndServe(":9090", r))
}