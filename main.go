package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func handleHello(w http.ResponseWriter, _ *http.Request) {
	wc, err := w.Write([]byte("Ciao Mario!"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
	fmt.Printf("%d bytes written", wc)
}

func handleBye(w http.ResponseWriter, _ *http.Request) {
	wc, err := w.Write([]byte("Bye bye!"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
	fmt.Printf("%d bytes written", wc)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHello)
	mux.HandleFunc("/bye", handleBye)

	log.Fatal(http.ListenAndServe(":8081", mux))
}
