package controllers

import (
	"gobackend/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	filename, err := utils.UploadFile(c, "image")
	if err != nil {
		log.Println("Server", err)
		return err
	}
	return c.JSON(fiber.Map{
		"url": filename,
	})
}
