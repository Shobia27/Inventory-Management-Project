package controllers

import (
	"myproject/database"
	"myproject/models"

	"github.com/gofiber/fiber/v2"
)

func AddSupplier(c *fiber.Ctx) error {

	supplier := new(models.Supplier)
	if err := c.BodyParser(supplier); err != nil {
		return err
	}
	database.DB.Create(&supplier)
	return c.JSON(&supplier)

}

func GetAllSuppliers(c *fiber.Ctx) error {

	var suppliers []models.Supplier
	database.DB.Find(&suppliers)

	return c.JSON(&suppliers)

}

func GetSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplier models.Supplier
	database.DB.Where("supplier_id = ?", id).Find(&supplier)
	return c.JSON(&supplier)
}

func DeleteSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplier models.Supplier
	database.DB.Where("supplier_id = ?", id).Find(&supplier)
	if supplier.Supplier_Name == "" {
		return c.JSON(fiber.Map{
			"message": "Supplier Not Found :(",
		})
	}

	database.DB.Delete(&supplier)
	return c.JSON(fiber.Map{
		"message": "Supplier is Deleted!!",
	})
}

func UpdateSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	supplier := new(models.Supplier)

	database.DB.Where("supplier_id = ?", id).First(&supplier)
	if supplier.Supplier_Name == "" {
		return c.JSON(fiber.Map{
			"message": "Supplier Not Found :(",
		})
	}

	if err := c.BodyParser(supplier); err != nil {
		return err
	}

	database.DB.Save(&supplier)

	return c.JSON(&supplier)
}

func UpdateSupplierProductCountINC(c *fiber.Ctx) error {

	id := c.Params("id")
	supplier := new(models.Supplier)

	database.DB.Where("supplier_id = ?", id).First(&supplier)
	supplier.Num_Inventory_Product += 1
	database.DB.Save(&supplier)

	return c.JSON(&supplier)
}

func UpdateSupplierProductCountDEC(c *fiber.Ctx) error {

	id := c.Params("id")
	supplier := new(models.Supplier)

	database.DB.Where("supplier_id = ?", id).First(&supplier)
	supplier.Num_Inventory_Product -= 1
	database.DB.Save(&supplier)

	return c.JSON(&supplier)
}
