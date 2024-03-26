//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/dns-query", queryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method: %v, URL: %v\n", r.Method, r.URL)
	http.NotFound(w, r)
}
