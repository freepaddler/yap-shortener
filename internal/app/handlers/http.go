package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi"

	"github.com/freepaddler/yap-shortener/internal/app"
)

type HTTPHandler struct {
	s app.Storage
	b string
}

func NewHTTPHandler(s app.Storage, b string) *HTTPHandler {
	return &HTTPHandler{s: s, b: b}
}

func (h *HTTPHandler) Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Put Handler:", r.URL)
	ct := r.Header.Get("Content-Type")
	if !strings.Contains(ct, "text/plain") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var rBody []byte
	defer r.Body.Close()
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(rBody) < 1 {
		fmt.Println("Empty body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hash := h.s.Put(string(rBody))
	hash = h.b + "/" + hash
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(hash))
}

func (h *HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Handler:", r.URL)
	id := chi.URLParam(r, "id")
	u, ok := h.s.Get(id)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("Get Handler return:", u)
	w.Header().Add("Location", u)
	w.WriteHeader(http.StatusTemporaryRedirect)

}
