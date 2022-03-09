package main

import (
	"fiber-project/blog"
	"fiber-project/database"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/post", blog.GetPosts).Name("getPosts")
	app.Get("/post/:id", blog.GetPost).Name("getPost")
	app.Post("/post", blog.NewPost).Name("newPost")
	//app.Put("/blog/:id", blog.Update)
	app.Delete("/post/:id", blog.DeletePost)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "blog.db")
	if err != nil {
		log.Infof("failed to connect database")
	}
	log.Infof("database connected")

	database.DBConn.AutoMigrate(&blog.Post{})
	log.Infof("database migrated")
}

func main() {
	app := fiber.New(
		fiber.Config{
			AppName:           "Test App v1.0.1",
			EnablePrintRoutes: true,
		},
	)

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(":3000")
}
