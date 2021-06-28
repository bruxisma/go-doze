package test

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router mux.Router

func NewRouter() *Router {
	return (*Router)(mux.NewRouter())
}

func (router *Router) inner() *mux.Router {
	return (*mux.Router)(router)
}

func routerHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		handler(writer, request)
	}
}

func (router *Router) Method(method, path string, handler http.HandlerFunc) *Router {
	router().inner().Methods(method).Path(path).HandlerFunc(routerHandler(handler))
	return router
}

func (router *Router) Get(path string, handler http.HandlerFunc) *Router {
	return router.Method("GET", path, handler)
}

func (router *Router) Patch(path string, handler http.HandlerFunc) *Router {
	return router.Method("PATCH", path, handler)
}

func (router *Router) Post(path string, handler http.HandlerFunc) *Router {
	return router.Method("POST", path, handler)
}

func (router *Router) Put(path string, handler http.HandlerFunc) *Router {
	return router.Method("PUT", path, handler)
}

func (router *Router) Delete(path string, handler http.HandlerFunc) *Router {
	return router.Method("DELETE", path, handler)
}
