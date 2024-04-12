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
	app.Post("/books", createBook)
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

func createBook(c *fiber.Ctx) error {
	book := new(Book) // same as *Book, don't need to use &
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, b := range books {
		if b.Id == book.Id || b.Title == book.Title {
			return c.Status(fiber.StatusBadRequest).SendString("The id is already in use.")
		}
	}
	books = append(books, *book)
	return c.JSON(book)
}
