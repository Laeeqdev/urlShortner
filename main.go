package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("app started")
	myRouter := InitializeApp()
	r := myRouter.MyRouter()
	defer log.Fatal(http.ListenAndServe(":9090", r))
}
