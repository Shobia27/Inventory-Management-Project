package controllers

import (
	"context"
	"log"
	"myproject/database"
	"myproject/models"
	"net/http"
	"strconv"

	"myproject/kafka"

	"github.com/gofiber/fiber/v2"
)

// func toJson(val []byte) models.Product {
// 	prod := models.Product{}
// 	err := json.Unmarshal(val, &prod)
// 	if err != nil {

// 		panic(err)
// 	}
// 	return prod
// }

type ProductSample struct {
	// ID            uint    `json:"ID"`
	// CreatedAt     string  `json:"CreatedAt"`
	// UpdatedAt     string  `json:"UpdatedAt"`
	// DeletedAt     string  `json:"DeletedAt"`
	Id            uint    `json:"Id"`
	Product_Name  string  `json:"productname"`
	Quantity      float64 `json:"quantity"`
	SupplierId    string  `json:"supplierid"`
	Selling_Price float64 `json:"sp"`
}

func AddProduct(c *fiber.Ctx) error {

	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return err
	}
	database.DB.Create(&product)

	//to add record to redis
	insertedproductid := product.Id
	redisAddProductApi := "http://localhost:3000/api/redis/addProductCache/" + strconv.FormatInt(int64(insertedproductid), 10)
	_, err := http.Get(redisAddProductApi)

	if err != nil {
		log.Fatal(err)
	}

	//call producer of kafka
	ctx := context.Background()
	sid := product.SupplierId
	kafka.Produce(ctx, sid, "ADD")
	kafka.Consume(ctx)

	//return added product
	return c.JSON(&product)
}

func GetAllProducts(c *fiber.Ctx) error {

	var products []models.Product
	database.DB.Find(&products)

	return c.JSON(&products)

}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	database.DB.Find(&product, id)
	return c.JSON(&product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	database.DB.First(&product, id)
	if product.Product_Name == "" {
		return c.JSON(fiber.Map{
			"message": "Product Not Found :(",
		})
	}

	database.DB.Delete(&product)

	//call producer of kafka
	ctx := context.Background()
	sid := product.SupplierId
	kafka.Produce(ctx, sid, "DEL")
	kafka.Consume(ctx)

	return c.JSON(fiber.Map{
		"message": "Product is Deleted!!",
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(models.Product)

	database.DB.First(&product, id)
	if product.Product_Name == "" {
		return c.JSON(fiber.Map{
			"message": "Product Not Found :(",
		})
	}

	if err := c.BodyParser(product); err != nil {
		return err
	}

	database.DB.Save(&product)

	//to update record to redis
	redisAddProductApi := "http://localhost:3000/api/redis/addProductCache/" + id
	_, err := http.Get(redisAddProductApi)

	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(&product)
}
