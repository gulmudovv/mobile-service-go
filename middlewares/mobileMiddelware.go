package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

const token = "A0gPCdf8Vn!4otgc!1=/yJ1I7ycVOxX4SWWHT!WePNVvFSu2!?AkhTMB7m12PXJs"

type Jwt struct {
	Token string `json:"token"`
}

func IsMobileJWT(c *fiber.Ctx) error {

	jwt := new(Jwt)

	if err := c.BodyParser(jwt); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "mobile unauthenticated",
		})
	}

	if jwt.Token != token {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "mobile unauthorized",
		})
	}

	return c.Next()
}
