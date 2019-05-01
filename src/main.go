package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var counter = 0

func main() {
	fmt.Println("Go Docker Tutorial!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter++
		fmt.Fprintf(w, "Hello World! " + strconv.Itoa(counter))
	})

	log.Fatal(http.ListenAndServe(":8088", nil))
}
