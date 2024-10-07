package utils

import (
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx, name string) (string, error) {

	form, err := c.MultipartForm()

	if err != nil {
		return "", err
	}
	files := form.File[name]

	fileName := ""
	for _, file := range files {
		fileName = randomName(file, name)

		if err := c.SaveFile(file, "./uploads/"+name+"s"+"/"+fileName); err != nil {

			return "", err
		}

	}
	return fileName, nil

}

func randomName(f *multipart.FileHeader, name string) string {
	arr := strings.Split(f.Filename, ".")
	ex := arr[len(arr)-1]
	tmp := strconv.Itoa(int(time.Now().Unix()))

	return name + "_" + tmp + "." + ex
}
