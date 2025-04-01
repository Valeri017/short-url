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
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Registred")

		body, err := req.HundleBody[RegisterRequests](&w, r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(body)

	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Login")
		//прочитать боди
		//       var payload LoginRequests

		body, err := req.HundleBody[LoginRequests](&w, r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(body)

		// 	err := json.NewDecoder(r.Body).Decode(&payload)
		// 	if err != nil {
		// 		res.Json(w, err.Error(), 402)
		// 	}
		// 	//Валидация
		// 	validate := validator.New()
		// 	validate.Struct(payload)

		// 	fmt.Println(payload)
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}
}
