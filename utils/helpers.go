package utils

import (
	"gobackend/models"
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, entity models.Entity, page int) fiber.Map {
	limit := 14
	offset := (page - 1) * limit
	lastPage := 0

	data := entity.Take(db, limit, offset)
	total := entity.Count(db)
	if int(total) <= limit {
		lastPage = 1
	} else {
		lastPage = int(math.Ceil(float64(int(total)/limit)) + 1)
	}
	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage,
		},
	}
}
