package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/api/list", func(c *fiber.Ctx) error {
		list := []map[string]interface{}{
			{"Name": "Bengo"},
			{"Name": "Benguela"},
			{"Name": "Bié"},
			{"Name": "Cabinda"},
			{"Name": "Cuando Cubango"},
			{"Name": "Cuanza Norte"},
			{"Name": "Cuanza Sul"},
			{"Name": "Cunene"},
			{"Name": "Huambo"},
			{"Name": "Huíla"},
			{"Name": "Luanda"},
			{"Name": "Lunda Norte"},
			{"Name": "Lunda Sul"},
			{"Name": "Malanje"},
			{"Name": "Moxico"},
			{"Name": "Namibe"},
			{"Name": "Uíge"},
			{"Name": "Zaire"},
		}
		return c.JSON(list)
	})

	app.Listen(":3000")
}
