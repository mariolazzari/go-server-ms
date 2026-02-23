package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func handleRoot(w http.ResponseWriter, _ *http.Request) {
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

func handleHello(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	userList := params["user"]

	// default user value
	username := "User"
	if len(userList) > 0 {
		username = userList[0]
	}

	var output bytes.Buffer
	output.WriteString("Ciao ")
	output.WriteString(username)
	output.WriteString("!")

	_, err := w.Write(output.Bytes())
	if err != nil {
		slog.Error("error writing response body", "err", err)
		return
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/hello", handleHello)
	mux.HandleFunc("/bye", handleBye)

	log.Fatal(http.ListenAndServe(":8081", mux))
}
