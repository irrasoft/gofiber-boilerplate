package controllers

import (
	"app/models"
	"app/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetAllProducts ...
func GetAllProducts(c *fiber.Ctx) error {
	if data, total, err := models.GetAllProducts(utils.GetParams(c)); err == nil {
		return c.JSON(fiber.Map{
			"data":   data,
			"total":  total,
			"status": "success",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "fail",
		"message": "Can't get data",
	})
}

// GetProduct ...
func GetProduct(c *fiber.Ctx) error {
	id := utils.GetIntParams(c, "id")
	data, err := models.GetProductByID(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"data":    nil,
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":   data,
		"status": "success",
	})
}

// CreateProduct ...
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(&product); err != nil {
		fmt.Println("BodyParser err", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid payload",
		})
	}

	data, err := models.CreateProduct(product)

	if err != nil {
		return c.JSON(fiber.Map{
			"data":    nil,
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":   data,
		"status": "success",
	})
}

// UpdateProduct ...
func UpdateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		fmt.Println("err", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid payload",
		})
	}

	data, err := models.UpdateProduct(product)

	if err != nil {
		fmt.Println("err", err)
		return c.JSON(fiber.Map{
			"status":  "fail",
			"message": "Record not found",
		})
	}

	return c.JSON(fiber.Map{
		"data":   data,
		"status": "success",
	})
}

// DeleteProduct ...
func DeleteProduct(c *fiber.Ctx) error {
	id := utils.GetIntParams(c, "id")
	if err := models.DeleteProduct(id, true); err != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": err})
	}

	return c.JSON(fiber.Map{"status": "success", "data": id})
}

// DeleteProducts ...
func DeleteProducts(c *fiber.Ctx) error {
	var list []int
	// parse data
	if err := c.BodyParser(&list); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid payload",
		})
	}
	// delete
	if err := models.MultiDeleteProducts(list, true); err != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": err})
	}

	return c.JSON(fiber.Map{"status": "success", "data": list})
}
