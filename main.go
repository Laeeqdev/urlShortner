package main

import (
	"log"
	"net/http"
)

func main() {
	myRouter := InitializeApp()
	r := myRouter.MyRouter()
	defer log.Fatal(http.ListenAndServe(":9090", r))
}
