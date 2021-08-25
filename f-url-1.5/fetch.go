// fet prings the content found at a URL

package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") != true {
			url = "http://" + url
			fmt.Println("If 1 complete") // test code
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
			fmt.Println("If 2 complete") // test code
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
			fmt.Println("If 3 complete")
		}
		fmt.Printf("%s", b)
	}
}
