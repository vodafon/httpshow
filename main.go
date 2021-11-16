package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	flagPort = flag.String("port", "8090", "port")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", logRequest)

	http.ListenAndServe(":"+*flagPort, nil)
}

func logRequest(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("%s\n\n\n%s\n-------------------------------------------------\n\n\n",
		time.Now(),
		formatRequest(r),
	)

	fmt.Fprintf(w, msg)
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string // Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)                             // Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host)) // Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		defer r.Body.Close()
	}
	if err == nil && len(body) > 0 {
		request = append(request, "\n")
		request = append(request, fmt.Sprintf("%q", body))
	}
	return strings.Join(request, "\n")
}
