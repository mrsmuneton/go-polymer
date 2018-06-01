package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	appengine.Main()
	http.ListenAndServe("localhost:8080", nil)
}
