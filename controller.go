package main

import (
	"github.com/gofiber/fiber/v2"
)

func ParseRequest(ctx *fiber.Ctx) (string, string) {
	return ctx.Params("username"), ctx.Query("password")
}

func Verify(ctx *fiber.Ctx) error {
	username, password := ParseRequest(ctx)
	hash := GetHash(&username)

	if hash == nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	} else if VerifyPassword(&hash, &password) {
		return ctx.SendStatus(fiber.StatusOK)
	} else {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
}

func SignUp(ctx *fiber.Ctx) error {
	username, password := ParseRequest(ctx)

	passwordHash := HashPassword(&password)

	return Insert(&username, &passwordHash)
}

func Reset(ctx *fiber.Ctx) error {
	username, password := ParseRequest(ctx)
	passwordHash := HashPassword(&password)

	return Update(&username, &passwordHash)
}
