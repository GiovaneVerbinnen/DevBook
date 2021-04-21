package router

import (
	router "api/src/router/routes"

	"github.com/gorilla/mux"
)

// Gerar Gera rotas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return router.Configurar(r)
}
