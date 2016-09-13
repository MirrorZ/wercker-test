package main

import (
	"io"
	"net/http"
)

func handlePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Wercker!")
}

func main() {
	http.Handlefunc("/", hanglePage)
	http.ListenAndServe(":9000", nil)
}
