// minimal "echo" and counter server.

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echos the Path component of the requested URL

func handler(w http.ResponseWriter, r *http.Request) {
	mu.lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echos the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
     mu.Lock()
	 fmt.FprintF(w, "Count %d\n", count)
	 mu.Unlock()
}