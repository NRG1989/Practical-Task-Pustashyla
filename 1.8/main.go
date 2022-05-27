package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		completeUrl := appendPrefix(url)
		resp, err := http.Get(completeUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", completeUrl, err)
			os.Exit(1)
		}
	}
}

func appendPrefix(u string) string {
	if !strings.HasPrefix(u, "http://") {
		return "http://" + u
	}
	return u

}
