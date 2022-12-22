package nodes

import (
	"errors"

	"github.com/Aayush9029/PiedPiper/database"
	"github.com/Aayush9029/PiedPiper/pkg"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Node struct {
	database.DefaultModel
	Key  string `gorm:"PRIMARY_KEY"`
	Data []byte `json:"data"`
}

func AddUpdate(c *fiber.Ctx) error {
	key := c.Params("key")
	db := database.DB
	var node Node
	err := db.First(&node, "key = ?", key).Error

	// if node not found, create new node
	if errors.Is(err, gorm.ErrRecordNotFound) {
		println("node not found, creatimg a mew mode")
		node = Node{
			Key:  key,
			Data: c.Body(),
		}

		println(node.Key)

		db.Create(&node)
		return c.JSON(node)
	} else {
		// if node found, update node
		updatedNode := Node{
			Key:  key,
			Data: c.Body(),
		}

		if err := db.Model(&node).Updates(updatedNode).Error; err != nil {
			return pkg.Unexpected(err.Error())
		}

		return c.JSON(updatedNode)

	}
}

func Get(c *fiber.Ctx) error {
	key := c.Params("key")
	db := database.DB
	var node Node
	err := db.First(&node, "key = ?", key).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Node not found",
		})
	}
	return c.JSON(node)
}
