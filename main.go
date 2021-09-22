package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	srv := &http.Server{Addr: ":" + os.Getenv("ABX_PORT"), Handler: mux}
	mux.HandleFunc("/", index)
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	k := req.URL.Path[1:]
	v := os.Getenv(k)
	fmt.Fprintf(w, "%s\n", v)
}
