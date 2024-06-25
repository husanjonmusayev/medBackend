package handlers

import (
	"fmt"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the profile page!")
}
