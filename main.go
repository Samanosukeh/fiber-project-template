package main

import (
	"fiber-project/blog"
	"github.com/gofiber/fiber/v2"
)

func HelloWord(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"Hello": "World"})
}

func main() {
	app := fiber.New(
		fiber.Config{
			AppName:           "Test App v1.0.1",
			EnablePrintRoutes: true,
		},
	)

	// GET http://localhost:3000/john
	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is john?")
	}).Name("home")

	// create post route for test
	app.Post("/", HelloWord)

	app.Post("/blog/create", blog.CreateBlogItem)

	app.Listen(":3000")
}
