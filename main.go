package main

import "github.com/gofiber/fiber/v2"

func main() {
	ConnectToDB()

	app := fiber.New()

	app.Get("/verify/:username", Verify)
	app.Post("/signup/:username", SignUp)
	app.Post("/reset/:username", Reset)

	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
}
