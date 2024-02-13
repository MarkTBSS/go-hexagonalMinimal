package adapters

import (
	"github.com/MarkTBSS/go-hexagonalMinimal/core"
	"github.com/gofiber/fiber/v2"
)

type HTTPHandler struct {
	usecase *core.EmployeeUsecase
}

func NewHTTPHandler(usecase *core.EmployeeUsecase) *HTTPHandler {
	return &HTTPHandler{usecase}
}

func (h *HTTPHandler) CreateEmployee(c *fiber.Ctx) error {
	var req struct {
		Name   string `json:"name"`
		Salary string `json:"salary"`
		Age    string `json:"age"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := h.usecase.CreateEmployee(req.Name, req.Salary, req.Age); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (h *HTTPHandler) GetEmployeeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := h.usecase.GetEmployeeByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	return c.JSON(employee)
}
