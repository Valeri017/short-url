package auth

type LoginRequests struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequests struct {
	LoginRequests
	Name string `json:"name"     validate:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
