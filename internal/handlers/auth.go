package handlers

import (
	"fmt"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // Change to POST
		http.Error(w, "Metod noto'g'ri POST methodidan foydalaning", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "Auth handler ishga tushdi!")
}
