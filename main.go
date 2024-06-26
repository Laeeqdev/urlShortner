package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	myRouter := InitializeApp(make(map[string]string), make(map[string]string), new(sync.RWMutex))
	r := myRouter.MyRouter()
	defer log.Fatal(http.ListenAndServe(":9090", r))
}
