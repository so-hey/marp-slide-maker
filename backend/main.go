package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodOptions {
			c.Set("Access-Control-Allow-Origin", "*")
			c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			return c.Status(http.StatusOK).JSON(fiber.Map{
				"message": "preflight request is accepted",
			})
		}
		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	setupApiRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)

}

func setupApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/convertmd", convertMd)
	api.Post("/convertpdf", convertPdf)
	api.Post("convertpptx", convertPptx)
}

func convertMd(c *fiber.Ctx) error {
	requestBody := string(c.Body())
	var jsonBody map[string]interface{}

	err := json.Unmarshal([]byte(requestBody), &jsonBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid JSON")
	}

	mdText := jsonBody["mdText"].(string)

	tmp := os.TempDir()
	mdFilePath := filepath.Join(tmp, "new.md")

	err = os.WriteFile(mdFilePath, []byte(mdText), 0644)
	if err != nil {
		return c.Status(500).SendString("Failed to write md file")
	}

	return c.SendFile(mdFilePath)
}

func convertPdf(c *fiber.Ctx) error {
	requestBody := string(c.Body())
	var jsonBody map[string]interface{}

	err := json.Unmarshal([]byte(requestBody), &jsonBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid JSON")
	}

	mdText := jsonBody["mdText"].(string)

	tmp := os.TempDir()
	mdFilePath := filepath.Join(tmp, "new.md")
	pdfFilePath := filepath.Join(tmp, "new.pdf")
	themeFilePath := filepath.Join("public", "css", "dakken_dark_theme.css")

	err = os.WriteFile(mdFilePath, []byte(mdText), 0644)
	if err != nil {
		return c.Status(500).SendString("Failed to write md file")
	}

	cmd := exec.Command("npx", "@marp-team/marp-cli", "--theme", themeFilePath, "--html", "--pdf", mdFilePath, "-o", pdfFilePath)
	err = cmd.Run()
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Failed to convert to PDF: \n%v", err))
	}

	c.Set("Content-Type", "application/pdf")
	return c.SendFile(pdfFilePath)
}

func convertPptx(c *fiber.Ctx) error {
	requestBody := string(c.Body())
	var jsonBody map[string]interface{}

	err := json.Unmarshal([]byte(requestBody), &jsonBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid JSON")
	}

	mdText := jsonBody["mdText"].(string)

	tmp := os.TempDir()
	mdFilePath := filepath.Join(tmp, "new.md")
	pptxFilePath := filepath.Join(tmp, "new.pptx")
	themeFilePath := filepath.Join("public", "css", "dakken_dark_theme.css")

	err = os.WriteFile(mdFilePath, []byte(mdText), 0644)
	if err != nil {
		return c.Status(500).SendString("Failed to write md file")
	}

	cmd := exec.Command("npx", "@marp-team/marp-cli", "--theme", themeFilePath, "--html", "--pptx", mdFilePath, "-o", pptxFilePath)
	err = cmd.Run()
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Failed to convert to PDF: \n%v", err))
	}

	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.presentationml.presentation")
	return c.SendFile(pptxFilePath)
}
