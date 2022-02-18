package routes

import (
	"myproject/controllers"
	"myproject/redisCache"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	//-------------------------------------------------------------------//
	// 		FOR REGISTERING USER
	//-------------------------------------------------------------------//
	app.Post("/api/registerUser", controllers.RegisterUser)

	//-------------------------------------------------------------------//
	// 		FOR LOGGING IN
	//-------------------------------------------------------------------//
	app.Post("/api/login", controllers.Login)

	//-------------------------------------------------------------------//
	// 		FOR INVENTORY PRODUCT TABLE
	//-------------------------------------------------------------------//
	app.Post("/api/addProduct", controllers.AddProduct)
	app.Get("/api/getAllProducts", controllers.GetAllProducts)
	app.Get("/api/getProduct/:id", controllers.GetProduct)
	app.Delete("/api/deleteProduct/:id", controllers.DeleteProduct)
	app.Put("/api/updateProduct/:id", controllers.UpdateProduct)

	//-------------------------------------------------------------------//
	// 		FOR REDIS CACHE
	//-------------------------------------------------------------------//
	app.Get("/api/redis/getAllProductCache", redisCache.GetAllProductCache)
	app.Get("/api/redis/addProductCache/:id", redisCache.AddProductCache)
	app.Delete("/api/redis/deleteProductCache/:id", redisCache.DeleteProductCache)
	app.Get("/api/redis/getProductCache/:id", redisCache.GetProductCache)

	//-------------------------------------------------------------------//
	// 		FOR SUPPLIER DETAILS TABLE
	//-------------------------------------------------------------------//
	app.Post("/api/addSupplier", controllers.AddSupplier)
	app.Get("/api/getAllSuppliers", controllers.GetAllSuppliers)
	app.Get("/api/getSupplier/:id", controllers.GetSupplier)
	app.Delete("/api/deleteSupplier/:id", controllers.DeleteSupplier)
	app.Put("/api/updateSupplier/:id", controllers.UpdateSupplier)
	app.Get("/api/updateSupplierProductCountINC/:id", controllers.UpdateSupplierProductCountINC)
	app.Get("/api/updateSupplierProductCountDEC/:id", controllers.UpdateSupplierProductCountDEC)

}
