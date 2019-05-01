package main

import (
	"fmt"
	"log"
	"net/http"
	"os"	
	"runtime"
	"strconv"
)

var counter = 0

func main() {
	fmt.Println("Go Docker Tutorial!")

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}	

	env_desc := fmt.Sprintf("%s - %s (%d-core): %s - %s", runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), hostname, executable)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter++
		fmt.Fprintf(w, "Hello World from ["+env_desc+"] - "+strconv.Itoa(counter))
	})

	log.Fatal(http.ListenAndServe(":8088", nil))
}
