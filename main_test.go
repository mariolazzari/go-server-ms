package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHello(t *testing.T) {
	w := httptest.NewRecorder()
	handleHello(w, nil)
	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code: expected %v but get %v\nbody: %s\n", desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Ciao Mario!")
	if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
		t.Errorf("bad response: expected %q but get %q", expectedMessage, w.Body.String())
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
