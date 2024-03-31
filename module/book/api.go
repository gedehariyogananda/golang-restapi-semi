package book

import (
	"test/golang/domain"
	"test/golang/helper"

	"github.com/gofiber/fiber/v2"
)

type api struct {
	bookService domain.BookService
}

func NewApi(app *fiber.App, bookService domain.BookService) {
	api := api{bookService}

	apiRoute := app.Group("/api")

	bookRoute := apiRoute.Group("/book")

	// api books
	bookRoute.Get("/:author", api.findByAuthor)
	bookRoute.Get("/", api.findAll)
	bookRoute.Post("/", api.createData)
}

func (a *api) findAll(c *fiber.Ctx) error {
	response := a.bookService.FindAll()

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (a *api) findByAuthor(c *fiber.Ctx) error {
	author := c.Params("author")

	if author == "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponse{
			Message: "Author is required",
			Status:  false,
		})
	}

	response := a.bookService.FindByAuthor(author)

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (a *api) createData(c *fiber.Ctx) error {
	execute := domain.Book{}

	if execute.Author == "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponse{
			Message: "Author is required",
			Status:  false,
		})
	}

	if execute.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponse{
			Message: "Title is required",
			Status:  false,
		})
	}

	if execute.Description == "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponse{
			Message: "Description is required",
			Status:  false,
		})

	}

	if err := c.BodyParser(&execute); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(helper.ApiResponse{
			Message: "Failed to parse request body",
			Status:  false,
		})
	}

	response := a.bookService.Create(execute)

	return c.Status(fiber.StatusCreated).JSON(response)
}
