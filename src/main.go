package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"os"
)

var counter = 0

func main() {
	fmt.Println("Go Docker Tutorial!")
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter++
		fmt.Fprintf(w, "Hello World from " + hostname + " - " + strconv.Itoa(counter))
	})

	log.Fatal(http.ListenAndServe(":8088", nil))
}
