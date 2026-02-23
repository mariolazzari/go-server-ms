# Go HTTP Servers

[YouTube playlist](https://www.youtube.com/watch?v=lyh9G3-Q7PQ&list=PLlv5lnjOHQo7m2w-KjtTZq10ZNVfNHHdP)

## Introduction

[Github repository](https://github.com/MongooseStudios/GoHTTPServerTutorial)

## Basic HTTP server

```go
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
```
