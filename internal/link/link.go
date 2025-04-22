package link

import (
	"fmt"
	"net/http"
)

type LinkHandler struct {
	// *configs.Config
}

type LinkHandlerDeps struct {
	// *configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		// Config: deps.Config,
	}
	router.HandleFunc("POST /link/create", handler.Create())
	router.HandleFunc("GET /{alias}", handler.GoTo())
	router.HandleFunc("DELETE /link/delete", handler.Delete())
	router.HandleFunc("PATH /", handler.Update())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Create")
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GoTo")
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Delete")
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update")
	}
}
