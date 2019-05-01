package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
)

var _counter = 0
var _hostname = getHostName()
var _executable = getExecutable()
var _goos = runtime.GOOS
var _goarch = runtime.GOARCH
var _cpus = runtime.NumCPU()

func default_handler(w http.ResponseWriter, r *http.Request) {
	env_desc := fmt.Sprintf("%s - %s (%d-core): %s - %s", _goos, _goarch, _cpus, _hostname, _executable)

	_counter++
	fmt.Println("Handler called: " + strconv.Itoa(_counter))
	fmt.Fprintf(w, "Hello World from ["+env_desc+"] - "+strconv.Itoa(_counter))
}

func getHostName() string {
	hostname, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	return hostname
}

func getExecutable() string {
	executable, err := os.Executable()

	if err != nil {
		panic(err)
	}

	return executable
}

func main() {
	fmt.Println("Go Docker Tutorial!")

	http.HandleFunc("/", default_handler)

	log.Fatal(http.ListenAndServe(":8088", nil))
}
