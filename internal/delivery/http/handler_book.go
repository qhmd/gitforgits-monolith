package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/config"
	bukuStruct "github.com/qhmd/gitforgits/internal/domain/book"
	"github.com/qhmd/gitforgits/internal/dto/book"
	"github.com/qhmd/gitforgits/internal/middleware"
	"github.com/qhmd/gitforgits/internal/usecase"
	"github.com/qhmd/gitforgits/utils"
)

type BookHandler struct {
	Usecase *usecase.BookUseCase
}

func NewBookHandler(app *fiber.App, uc *usecase.BookUseCase) {
	h := &BookHandler{Usecase: uc}
	app.Get("/books", h.ListBook)
	app.Get("/books/:id<^[0-9]+$>", h.GetBookByID)
	app.Post("/books", middleware.JWT(), middleware.ValidateBook(), h.Create)
	app.Delete("/books/:id<^[0-9]+$>", h.Delete)
	app.Put("/books/:id<^[0-9]+$>", middleware.ValidateBook(), h.Update)
}

// ListBook godoc
// @Summary Get all books
// @Description Retrieve all books from the database
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} book.SuccessGetListBook
// @Failure 500 {object} book.ErrorResponse
// @Router /books [get]
func (h *BookHandler) ListBook(c *fiber.Ctx) error {
	books, err := h.Usecase.List(c.Context())
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "something went wrong", nil)
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "successfully get list book", books)
}

// GetBookByID godoc
// @Summary Get book by ID
// @Description Retrieve a single book by its ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} book.SuccessGetBook
// @Failure 500 {object} book.ErrorResponse
// @Router /books/{id} [get]
func (h *BookHandler) GetBookByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	book, err := h.Usecase.GetByID(c.Context(), id)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "books not found", nil)
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "successfully get the book", book)
}

// Create godoc
// @Summary Create a new book
// @Description Add a new book to the database
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param book body book.BookRequest true "Book data"
// @Success 201 {object} book.SuccessfullCreate
// @Failure 409 {object} book.TitleAlreadytaken
// @Failure 400 {object} book.MissingAuthorization
// @Failure 500 {object} book.ErrorResponse
// @Security ApiKeyAuth
// @Router /books [post]
func (h *BookHandler) Create(c *fiber.Ctx) error {
	req := c.Locals("validateBook").(book.BookRequest)
	book := &bukuStruct.Book{
		Title:  req.Title,
		Author: req.Author,
		Page:   req.Page,
	}
	if err := h.Usecase.Create(c.Context(), book); err != nil {
		if err == config.ErrBookTitleExists {
			return utils.ErrorResponse(c, fiber.StatusConflict, "title already exits, choose antoher title", config.ErrBookTitleExists)
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "something went wrong", nil)
	}
	return utils.SuccessResponse(c, fiber.StatusCreated, "successfully add the book", book)
}

// Update godoc
// @Summary Update book by ID
// @Description Update book information by its ID, u have to login first to access your access token
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Security ApiKeyAuth
// @Param id path int true "Book ID"
// @Param book body book.BookRequest true "Updated book data"
// @Success 200 {object} book.SuccessfullUpdate
// @Failure 400 {object} book.InvalidId
// @Failure 404 {object} book.BookNotFoundResponse
// @Failure 409 {object} book.ErrorResponse
// @Failure 500 {object} book.ErrorResponse
// @Router /books/{id} [put]
func (h *BookHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid id", err)
	}
	req := c.Locals("validateBook").(book.BookRequest)
	existing, err := h.Usecase.GetByID(c.Context(), id)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "book not found", err)
	}
	existing.Title = req.Title
	existing.Author = req.Author
	existing.Page = req.Page

	if err := h.Usecase.Update(c.Context(), existing); err != nil {
		if err == config.ErrBookTitleExists {
			return utils.ErrorResponse(c, fiber.StatusConflict, "choose another title", err.Error())
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "status went wrong", err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "update successfully", existing)
}

// Delete godoc
// @Summary Delete book by ID
// @Description Remove a book from the database using its ID
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} book.DeleteSuccesfully
// @Failure 400 {object} book.InvalidId
// @Failure 404 {object} book.BookNotFoundResponse
// @Failure 500 {object} book.ErrorResponse
// @Router /books/{id} [delete]
func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid id", err)
	}
	_, err = h.Usecase.GetByID(c.Context(), id)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "book with id {id} does not exist", "book not found")
	}
	err = h.Usecase.Delete(c.Context(), id)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "something went wrong", nil)
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "delete successfully", nil)
}
