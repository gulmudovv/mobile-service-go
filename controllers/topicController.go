package controllers

import (
	"gobackend/database"
	"gobackend/models"
	"gobackend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateTopic(c *fiber.Ctx) error {
	var topic models.Topic

	if err := c.BodyParser(&topic); err != nil {
		return err
	}

	database.DB.Create(&topic)

	return c.JSON(topic)
}

func AllTopics(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "0"))
	if page == 0 {
		var topics []models.Topic
		database.DB.Order("id").Find(&topics)
		return c.JSON(topics)
	}

	return c.JSON(utils.Paginate(database.DB, &models.Topic{}, page))

}

func Topics(c *fiber.Ctx) error {

	var topics []models.Topic

	database.DB.Find(&topics)

	return c.JSON(topics)

}

func GetTopic(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	topic := models.Topic{
		Id: uint(id),
	}

	database.DB.Find(&topic)

	return c.JSON(topic)
}

func UpdateTopic(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	topic := models.Topic{
		Id: uint(id),
	}

	if err := c.BodyParser(&topic); err != nil {
		return err
	}

	database.DB.Model(&topic).Updates(topic)

	return c.JSON(topic)
}

func DeleteTopic(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	topic := models.Topic{
		Id: uint(id),
	}

	database.DB.Delete(&topic)

	return nil
}
