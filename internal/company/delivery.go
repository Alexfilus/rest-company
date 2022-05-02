package company

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetList(ctx *fiber.Ctx) error
}
