package book

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"something went wrong"`
}

type InvalidId struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Your id is invalid"`
	Errors  string `json:"error" example:"Invalid id"`
}

type TitleAlreadytaken struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Title already exist, choose another titile"`
	Errors  string `json:"error" example:"book with this title already exists"`
}

type MissingAuthorization struct {
	Success bool `json:"success" example:"false"`

	Error string `json:"error" example:"missing authorization header"`
}
type BookNotFoundResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"book not found"`
	Errors  string `json:"errors" example:"book with id {id} does not exist"`
}
