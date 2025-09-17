package book

type BookRequest struct {
	Title  string `json:"title" validate:"required" example:"How To Become Backend Engineer"`
	Author string `json:"author" validate:"required,min=4,max=50,alphaSpace" example:"John Smith"`
	Page   int    `json:"page" validate:"required,gt=0" example:"205"`
}
