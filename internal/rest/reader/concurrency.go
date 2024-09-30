package reader

import "net/http"

func (h *handler) WithConcurrency(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Concurrency"))
}
