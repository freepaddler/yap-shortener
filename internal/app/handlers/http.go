package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/freepaddler/yap-shortener/internal/app"
)

type HTTPHandler struct {
	s app.Storage
}

func NewHTTPHandler(s app.Storage) *HTTPHandler {
	return &HTTPHandler{s: s}
}

func (h *HTTPHandler) Put(w http.ResponseWriter, r *http.Request) {
	ct := r.Header.Get("Content-Type")
	if ct != "text/plain" {
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
	hash := h.s.Put(string(rBody))
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write(hash)
}

func (h *HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	u, ok := h.s.Get(id)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Location", u)
	w.WriteHeader(http.StatusTemporaryRedirect)

}
