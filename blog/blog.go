package blog

import (
	"fiber-project/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title       string `json:"title" required:"true"`
	Description string `json:"description" required:"true"`
	Slug        string `json:"slug"`
}

// CreateItem - Cria um novo post no blog
//func CreatePost(title, description string) Post {
//	db := database.DBConn
//	textSlugify := slug.Make(title)
//	post := Post{
//		Title:       title,
//		Description: description,
//		Slug:        textSlugify,
//	}
//	db.Create(&post)
//}

func GetPosts(c *fiber.Ctx) error {
	db := database.DBConn
	var posts []Post
	db.Find(&posts)
	return c.JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var post Post
	db.First(&post, id)
	return c.JSON(post)
}

func NewPost(c *fiber.Ctx) error {
	db := database.DBConn
	post := new(Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	post.Slug = slug.Make(post.Title)

	db.Create(&post)
	return c.JSON(fiber.Map{"message": "created blog item"})
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var post Post
	db.First(&post, id)
	db.Delete(&post)
	return c.JSON(fiber.Map{"message": "deleted blog item"})
}
