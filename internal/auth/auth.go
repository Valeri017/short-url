package auth

import (
	"fmt"
	"log"
	"net/http"
	"short-url/configs"
	"short-url/pkg/req"
	"short-url/pkg/res"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())

}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Registred")
		body, err := req.HundleBody[RegisterRequests](&w, r)
		if err != nil {
			w.Write([]byte("Введите\002 адрес\002 электронной\002 почты\002 или\002 пароль\002 не \002может\002 быть\002 пустым"))
			log.Println(err)
		}
		fmt.Println(body)

	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Login")
		//прочитать боди
		body, err := req.HundleBody[LoginRequests](&w, r)
		if err != nil {
			w.Write([]byte("Введите \002 адрес\002 электронной \002почты\002 или \002пароль\002 не\002 может\002 быть\002 пустым"))
			log.Println(err)
		}
		fmt.Println(body)
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}
}
