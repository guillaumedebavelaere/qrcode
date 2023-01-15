package api

import (
	"net/http"

	"github.com/skip2/go-qrcode"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "text parameter is required", http.StatusBadRequest)
		return
	}

	qr, err := qrcode.New(text, qrcode.Highest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename=qrcode.png")
	if err := qr.Write(256, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/qrcode", Handler)
	http.ListenAndServe(":8080", nil)
}
