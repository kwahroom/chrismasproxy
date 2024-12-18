package main

import (
	"fmt"
	"net/http"

	"github.com/kwahroom/chrismasproxy/chrismasify"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/301", handle301)
	http.HandleFunc("/302", handle302)
	http.HandleFunc("/400", handle400)
	http.HandleFunc("/401", handle401)
	http.HandleFunc("/402", handle402)
	http.HandleFunc("/403", handle403)
	http.HandleFunc("/404", handle404)
	http.HandleFunc("/418", handle418)
	http.HandleFunc("/500", handle500)
	http.HandleFunc("/501", handle501)
	http.HandleFunc("/502", handle502)
	http.HandleFunc("/503", handle503)
	http.HandleFunc("/504", handle504)

	fmt.Println("Starting server on port 2412")
	if err := http.ListenAndServe(":2412", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
<html><body>
	<h1> All Codes </h1>
	<ul>
		<li><a href="/301">301</a></li>
		<li><a href="/302">302</a></li>
		<li><a href="/400">400</a></li>
		<li><a href="/401">401</a></li>
		<li><a href="/402">402</a></li>
		<li><a href="/403">403</a></li>
		<li><a href="/404">404</a></li>
		<li><a href="/418">418</a></li>
		<li><a href="/500">500</a></li>
		<li><a href="/501">501</a></li>
		<li><a href="/502">502</a></li>
		<li><a href="/503">503</a></li>
		<li><a href="/504">504</a></li>
	</ul>
</body></html>`)
}

func handle301(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 301)
}

func handle302(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 302)
}

func handle400(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 400)
}

func handle401(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 401)
}

func handle402(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 402)
}

func handle403(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 403)
}

func handle404(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 404)
}

func handle418(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 418)
}

func handle500(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 500)
}

func handle501(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 501)
}

func handle502(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 502)
}

func handle503(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 503)
}

func handle504(w http.ResponseWriter, r *http.Request) {
	chrismasify.WriteChrismasResponse(w, nil, 504)
}
