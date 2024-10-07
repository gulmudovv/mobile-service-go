package controllers

import (
	"gobackend/database"
	"gobackend/models"
	"gobackend/utils"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AllWords(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(utils.Paginate(database.DB, &models.Word{}, page))

}
func CreateWord(c *fiber.Ctx) error {

	var word models.Word

	if err := c.BodyParser(&word); err != nil {
		log.Println(err)
		return err
	}

	database.DB.Create(&word)

	return c.JSON(word)
}

func UpdateWord(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	word := models.Word{
		Id: uint(id),
	}

	if err := c.BodyParser(&word); err != nil {

		return err
	}

	database.DB.Model(&word).Updates(word)

	return c.JSON(word)
}

func GetWord(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	word := models.Word{
		Id: uint(id),
	}

	database.DB.Find(&word)

	return c.JSON(word)
}

func WordsByTopicId(c *fiber.Ctx) error {

	id := c.Params("id")
	var words []models.Word
	database.DB.Where("topic_id = ?", id).Find(&words)
	return c.JSON(words)
}

func RandomWordsByTopicId(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))
	var words []models.Word
	if id == -3 {
		database.DB.Order(gorm.Expr("random()")).Find(&words)
	} else if id == -2 {
		database.DB.Order(gorm.Expr("random()")).Limit(20).Find(&words)
	} else {
		database.DB.Order(gorm.Expr("random()")).Limit(5).Find(&words)
	}

	return c.JSON(words)
}
