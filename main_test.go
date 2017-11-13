package main_test

import (
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	_, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
}
