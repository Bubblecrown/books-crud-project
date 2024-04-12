package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// database
var books []Book

func main() {
	app := fiber.New()
	books = append(books, Book{Id: 1, Title: "Bubble", Author: "Bubble Crown"})
	app.Get("/books", getAllBooks)
	app.Get("/books/:id", getBook)
	app.Listen(":8080")
}

func getAllBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	booksId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, book := range books {
		if book.Id == booksId {
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("The book does not exist")
}
