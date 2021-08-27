// handler echos the HTTP request
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf("w, %s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header{
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAdd = %q\n", r.RemoteAddr)
	if err:= r.ParseForm(); err != nil{
        log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprint(w, "Form[%q] = %q\n", k, v)
	}
}