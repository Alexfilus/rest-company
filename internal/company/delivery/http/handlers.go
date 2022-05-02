package http

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"rest-company/config"
	"rest-company/internal/company"
	"rest-company/internal/models"
	"rest-company/pkg/logger"
)

type companyHandlers struct {
	cfg    *config.Config
	uc     company.UseCase
	logger logger.Logger
}

const companyID = "id"

func NewCompanyHandlers(
	cfg *config.Config,
	uc company.UseCase,
	logger logger.Logger,
) company.Handlers {
	return &companyHandlers{
		cfg:    cfg,
		uc:     uc,
		logger: logger,
	}
}

func (h *companyHandlers) Create(c *fiber.Ctx) error {
	companyObj := new(models.Company)
	err := json.Unmarshal(c.Body(), companyObj)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		h.logger.Error("companyHandlers-Create-json.Unmarshal", err)
		return err
	}
	res, err := h.uc.Create(c.UserContext(), companyObj)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		h.logger.Error("companyHandlers-Create-h.uc.Create", err)
		return err
	}
	err = c.JSON(res)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		h.logger.Error("companyHandlers-Create-c.JSON", err)
		return err
	}
	c.Status(fiber.StatusOK)
	return nil
}

func (h *companyHandlers) Update(c *fiber.Ctx) error {
	companyObj := new(models.Company)
	err := json.Unmarshal(c.Body(), companyObj)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		h.logger.Error("companyHandlers-Update-json.Unmarshal", err)
		return err
	}
	id := c.Params(companyID)
	err = h.uc.Update(c.UserContext(), id, companyObj)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		h.logger.Error("companyHandlers-Update-h.uc.Create", err)
		return err
	}
	c.Status(fiber.StatusOK)
	return nil
}

func (h *companyHandlers) Delete(c *fiber.Ctx) error {
	id := c.Params(companyID)
	err := h.uc.Delete(c.UserContext(), id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		h.logger.Error("companyHandlers-Delete-h.uc.Create", err)
		return err
	}
	c.Status(fiber.StatusOK)
	return nil
}

func (h *companyHandlers) GetByID(c *fiber.Ctx) error {
	id := c.Params(companyID)
	res, err := h.uc.GetByID(c.UserContext(), id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		h.logger.Error("companyHandlers-GetByID-h.uc.Create", err)
		return err
	}
	err = c.JSON(res)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		h.logger.Error("companyHandlers-GetByID-c.JSON", err)
		return err
	}
	c.Status(fiber.StatusOK)
	return nil
}

func (h *companyHandlers) GetList(c *fiber.Ctx) error {
	search := models.CompanySearch{}
	err := c.QueryParser(&search)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		h.logger.Error("companyHandlers-GetList-QueryParser", err)
		return err
	}
	res, err := h.uc.GetList(c.UserContext(), search)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		h.logger.Error("companyHandlers-GetList-h.uc.GetList", err)
		return err
	}
	err = c.JSON(res)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		h.logger.Error("companyHandlers-GetList-c.JSON", err)
		return err
	}
	c.Status(fiber.StatusOK)
	return nil
}
