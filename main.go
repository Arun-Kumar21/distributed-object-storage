package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Arun-Kumar21/distributed-object-storage/database"
	"github.com/Arun-Kumar21/distributed-object-storage/models"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDB()

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", os.ModePerm)
	}

	app := fiber.New()

	// Upload file to server
	app.Post("/upload", func(c *fiber.Ctx) error{
		file, err := c.FormFile("file")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "Invalid file"})
		}

		filepath := filepath.Join("uploads", file.Filename)

		// save file on disk
		if err = c.SaveFile(file, filepath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : "Failed to save file"})
		}

		// save metadata in DB
		dbFile := models.File{
			Name: file.Filename,
			Size: file.Size,
			Path: filepath,
		}

		database.DB.Create(&dbFile)

		return c.JSON(fiber.Map{"message" : "File uploaded", "file_id" : dbFile.ID})
	})

	// Retrive File
	app.Get("/file/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var file models.File
		if err := database.DB.First(&file, "id =?", id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error" : "File not found"})
		}

		return c.SendFile(file.Path)
	})

	// Get File metadata
	app.Get("/file/metadata/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var file models.File
		if err := database.DB.First(&file, "id=?", id).Error; err != nil{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error" : "File not found"})
		}

		return c.JSON(file)
	})

	log.Fatal(app.Listen(":8000"))
}