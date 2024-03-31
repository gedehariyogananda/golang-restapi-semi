package main

import (
	"fmt"
	"test/golang/config"
	"test/golang/module/book"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("tes project")

	// Connect to the database
	config.Connect()

	dbConnection := config.DB

	// book init modul
	bookRepository := book.NewRepository(dbConnection)
	bookService := book.NewService(bookRepository)

	// create fiber app
	app := fiber.New()

	// create book api
	book.NewApi(app, bookService)

	// port
	app.Listen(":3000")

}
