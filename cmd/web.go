package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func crash(w http.ResponseWriter, _ *http.Request) {
	panic("foo")

}

func health(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("ok"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", health)
	r.HandleFunc("/crash", crash)

	// todo: #2 register handlers

	// todo: #6 middlewares
	// 	* logger: %timestamp%: %method% %uri%
	//  * recovery: panic recovery

	log.Fatal(http.ListenAndServe(":8080", r))
}
