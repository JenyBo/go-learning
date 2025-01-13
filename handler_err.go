package main

import (
	"net/http"
)

func handleErr(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 200, "Got some error")
}
