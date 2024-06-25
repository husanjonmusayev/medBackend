package handlers

import (
    "fmt"
    "net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    fmt.Fprintln(w, "Auth handler ishga tushdi!")
}
