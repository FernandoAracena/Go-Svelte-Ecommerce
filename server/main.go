package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type Coffee struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	PictureURL  string `json:"picture_url"`
}

var DB *gorm.DB

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=Myrer123# dbname=CoffeeDB port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = DB.AutoMigrate(&Coffee{})
	if err != nil {
		log.Fatal("Failed to migrate to DB:", err)
	}

	// coffees := []Coffee{
	// 	{Name: "Expresso", Description: "Strong and black coffee", Price: 100, PictureURL: "https://unsplash.com/photos/dAYJfrtVjh0/download?force=true"},
	// 	{Name: "Latte", Description: "Cofee with milk", Price: 150, PictureURL: "https://unsplash.com/photos/3n3mPoGko8g/download?force=true"},
	// 	{Name: "Capuccino", Description: "Coffee with steamed milk foam", Price: 200, PictureURL: "https://unsplash.com/photos/gNfXpDPl6Vk/download?force=true"},
	// }
	// for _, coffee := range coffees {
	// 	DB.Create(&coffee)
	// }
}

func getCoffees(c *fiber.Ctx) error {
	var coffees []Coffee
	DB.Find(&coffees)
	return c.JSON(coffees)
}

func getCoffee(c *fiber.Ctx) error {
	id := c.Params("id")
	var coffee Coffee
	if err := DB.First(&coffee, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Coffee not found"})
	}
	return c.JSON(coffee)
}

func createCoffee(c *fiber.Ctx) error {
	coffee := new(Coffee)
	if err := c.BodyParser(coffee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	DB.Create(&coffee)
	return c.JSON(coffee)
}

func editCoffee(c *fiber.Ctx) error {
	id := c.Params("id")
	var coffee Coffee
	if err := DB.First(&coffee, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Coffee not found"})
	}
	var updatedCoffee map[string]interface{}
	if err := c.BodyParser(&updatedCoffee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot Parse JSON"})
	}

	if err := DB.Model(&coffee).Updates(updatedCoffee).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update Coffee"})
	}
	return c.Status(200).JSON(coffee)
}

func deleteCoffee(c *fiber.Ctx) error {
	id := c.Params("id")
	var coffee Coffee
	if err := DB.First(&coffee, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Coffee not found"})
	}
	if err := DB.Delete(&coffee).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete Coffee"})
	}
	return c.Status(200).JSON(coffee)
}

func main() {
	app := fiber.New()

	initDatabase()

	app.Get("/api/coffees", getCoffees)
	app.Get("/api/coffees/:id", getCoffee)
	app.Post("/api/coffees", createCoffee)
	app.Patch("/api/coffees/:id", editCoffee)
	app.Delete("/api/coffees/:id", deleteCoffee)

	log.Fatal(app.Listen(":3000"))
}
