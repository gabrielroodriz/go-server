package api

import (
	"net/http"

	"github.com/gabrielroodriz/go-server/internal/store/pgstore"

	"github.com/go-chi/chi/v5"
)

type apiHandler struct {
	query *pgstore.Queries
	r     *chi.Mux
}

func (h apiHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.r.ServeHTTP(writer, request)
}
func NewHandler(query *pgstore.Queries) http.Handler {
	api := apiHandler{query: query}
	r := chi.NewRouter()
	api.r = r
	return api
}
