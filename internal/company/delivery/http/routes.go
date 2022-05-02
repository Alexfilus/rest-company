package http

import (
	"github.com/gofiber/fiber/v2"
	"rest-company/internal/company"
)

func MapCompaniesRoutes(group fiber.Router, handlers company.Handlers) {
	group.Get(":id", handlers.GetByID)
	group.Get("", handlers.GetList)
	group.Post("", handlers.Create)
	group.Put(":id", handlers.Update)
	group.Delete(":id", handlers.Delete)
}
