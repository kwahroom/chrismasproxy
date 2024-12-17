package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/kwahroom/chrismasproxy/chrismasify"
)

// ReverseProxy is a simple reverse proxy that re-maps common HTTP status codes.
func ReverseProxy(target *url.URL) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			r.URL.Path = target.Path + r.URL.Path
			if _, ok := r.Header["User-Agent"]; !ok {
				r.Header.Set("User-Agent", "Chrismas-Proxy")
			}

			// Disable compression otherwise you have to re-compress the body
			r.Header.Del("Accept-Encoding")
		},
		ModifyResponse: chrismasify.ModifyResponse,
	}
}

func main() {
	target, err := url.Parse("http://example.com")
	if err != nil {
		log.Fatal(err)
	}

	proxy := ReverseProxy(target)

	http.Handle("/", proxy)
	fmt.Println("Chrismas proxy server is running on port 2412")
	log.Fatal(http.ListenAndServe(":2412", nil))
}
