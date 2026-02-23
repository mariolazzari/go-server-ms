package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRoot(t *testing.T) {
	w := httptest.NewRecorder()
	handleRoot(w, nil)
	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %v but get %v\nbody: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Ciao Mario!")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response: expected %q but get %q", expectedMessage, w.Body.String())
	}
}

func TestHandleHello(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?user=Mario", nil)
	w := httptest.NewRecorder()
	handleHello(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %v, get %v\nbody: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := "Ciao Mario!"
	if !bytes.Equal([]byte(expectedMessage), w.Body.Bytes()) {
		t.Errorf("bad response: expected %q, get %q\n", desiredCode, w.Body.String())
	}
}

func TestHandleHelloNoParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()
	handleHello(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %v, get %v\nbody: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := "Ciao User!"
	if !bytes.Equal([]byte(expectedMessage), w.Body.Bytes()) {
		t.Errorf("bad response: expected %q, get %q\n", desiredCode, w.Body.String())
	}
}

func TestHandleHelloWrongParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?foo=bar", nil)
	w := httptest.NewRecorder()
	handleHello(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %v, get %v\nbody: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := "Ciao User!"
	if !bytes.Equal([]byte(expectedMessage), w.Body.Bytes()) {
		t.Errorf("bad response: expected %q, get %q\n", desiredCode, w.Body.String())
	}
}

func TestHandleGoodbye(t *testing.T) {
	w := httptest.NewRecorder()
	handleBye(w, nil)
	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %v but get %v\nbody: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Bye bye!")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response: expected %q but get %q", expectedMessage, w.Body.String())
	}
}
