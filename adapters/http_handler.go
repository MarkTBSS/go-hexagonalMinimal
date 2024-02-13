package adapters

import (
	"log"

	"github.com/MarkTBSS/go-hexagonalMinimal/core/employee"
	"github.com/gofiber/fiber/v2"
)

type HTTPHandler struct {
	usecase *employee.EmployeeUsecase
}

func NewHTTPHandler(usecase *employee.EmployeeUsecase) *HTTPHandler {
	return &HTTPHandler{usecase}
}
func (h *HTTPHandler) CreateEmployee(c *fiber.Ctx) error {
	log.Println("IN : http_handler.go : adapters.CreateEmployee()")
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
	log.Println("OUT : http_handler.go : adapters.CreateEmployee()")
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
func (h *HTTPHandler) GetAllEmployees(c *fiber.Ctx) error {
	employees, err := h.usecase.GetAllEmployees()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(employees)
}
