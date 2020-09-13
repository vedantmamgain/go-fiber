package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vedantmamgain/go-fiber/books"
	"github.com/vedantmamgain/go-fiber/database"
)

func setupRoutes(c *fiber.App) {
	c.Get("/api/v1/book", books.GetBooks)
	c.Get("/api/v1/book/:id", books.GetBook)
	c.Post("/api/v1/book", books.NewBook)
	c.Delete("/api/v1/book/:id", books.DeleteBook)
	c.Get("/callback", books.Callback)
}

func initiDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database coneccted sucessfully")
	database.DBConn.AutoMigrate(&books.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initiDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)
	app.Listen(3000)
}
