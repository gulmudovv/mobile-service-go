package controllers

import (
	"gobackend/utils"

	"github.com/gofiber/fiber/v2"
)

func UploadAudio(c *fiber.Ctx) error {
	filename, err := utils.UploadFile(c, "audio")
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"url": filename,
	})
}
