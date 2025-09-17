package auth

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=4,max=50,alphaSpace,alphaMin4"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type UserResponse struct {
	RegisterRequest
	Role string `json:"role" validate:"required,oneof=user admin"`
}
