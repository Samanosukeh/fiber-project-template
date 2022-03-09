package blog

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

type Payload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// CreateItem - Cria um novo post no blog
func CreateItem(title, description string) {
	log.Infof("Creating blog item...")
	textSlugify := slug.Make(title)
	log.Infof(title)
	log.Infof(description)
	log.Infof(textSlugify)
}

func Blog(c *fiber.Ctx) {
	c.JSON(fiber.Map{"message": "blog teste"})
}

func CreateBlogItem(c *fiber.Ctx) error {
	var payload Payload
	jsonBytes, err := ioutil.ReadAll(c.Request) //.Body

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(c.Request.Body)

	if err != nil {
		fmt.Println(err)
	}

	_ = json.Unmarshal(jsonBytes, &payload)

	CreateItem(payload.Title, payload.Description)

	return c.JSON(fiber.Map{"message": "created blog item"})
}
