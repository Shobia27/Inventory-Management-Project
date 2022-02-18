package redisCache

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/gofiber/fiber/v2"
)

type ProductSample struct {
	ID            uint    `json:"ID"`
	CreatedAt     string  `json:"CreatedAt"`
	UpdatedAt     string  `json:"UpdatedAt"`
	DeletedAt     string  `json:"DeletedAt"`
	Id            uint    `json:"Id"`
	Product_Name  string  `json:"productname"`
	Quantity      float64 `json:"quantity"`
	SupplierId    string  `json:"supplierid"`
	Selling_Price float64 `json:"sp"`
}

var cache = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})
var ctx = context.Background()

//to convert the byte arr to productsample
func toJson1(val []byte) ProductSample {
	prod := ProductSample{}
	err := json.Unmarshal(val, &prod)
	if err != nil {
		panic(err)
	}
	return prod
}

func (s ProductSample) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s ProductSample) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}

//------------------------------------------------------------------------------------------
// to get all data(values for key-value pairs) from redis and convert it to suitable format
//-------------------------------------------------------------------------------------------

func GetAllProductCache(c *fiber.Ctx) error {

	var keys []string
	var err error
	var cursor uint64

	//getting all the keys form the redis db
	keys, _, err = cache.Scan(ctx, cursor, "*", 10).Result()
	if err != nil {
		panic(err)
	}

	//sorting the keys in ascending order...
	var intKeys = []int{}

	for _, i := range keys {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intKeys = append(intKeys, j)
	}
	sort.Ints(intKeys)
	keys = nil
	for _, i := range intKeys {
		k := strconv.Itoa(i)
		if err != nil {
			panic(err)
		}
		keys = append(keys, k)
	}
	//sorting of keys end..

	//getting all the values form the redis db
	var allProduct []ProductSample
	for _, keyval := range keys {
		val, err := cache.Get(ctx, keyval).Bytes()
		if err != nil {
			return err
		}

		iterdata := toJson1(val)
		allProduct = append(allProduct, iterdata)
	}

	return c.JSON(&allProduct)

}

//------------------------------------------------------------------------------------------
// to get new product from db and update in redis
//-------------------------------------------------------------------------------------------
func AddProductCache(c *fiber.Ctx) error {

	//waiting for 2 sec o get updated in postgresql product db
	time.Sleep(2 * time.Second)

	id := c.Params("id")
	res, err := http.Get("http://localhost:3000/api/getProduct/" + id)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var product ProductSample
	parseErr := json.Unmarshal(body, &product)
	if parseErr != nil {
		return parseErr
	}

	cacheErr := cache.Set(ctx, id, body, 0).Err()
	if cacheErr != nil {
		return cacheErr
	}

	//to retrieve data from redis with it "id"
	val, err := cache.Get(ctx, id).Bytes()
	if err != nil {
		return err
	}

	data := toJson1(val)
	return c.JSON(fiber.Map{"cached": data})

}

//------------------------------------------------------------------------------------------
// to delete product from redis
//-------------------------------------------------------------------------------------------
func DeleteProductCache(c *fiber.Ctx) error {
	id := c.Params("id")

	iter := cache.Scan(ctx, 0, id, 0).Iterator()
	for iter.Next(ctx) {
		cache.Del(ctx, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	var keys []string
	var err error
	var cursor uint64

	keys, _, err = cache.Scan(ctx, cursor, "*", 10).Result()
	if err != nil {
		panic(err)
	}
	return c.JSON(fiber.Map{"keys": keys})
}

//------------------------------------------------------------------------------------------
// to get detail of product from redis with given id
//-------------------------------------------------------------------------------------------
func GetProductCache(c *fiber.Ctx) error {
	id := c.Params("id")

	val, err := cache.Get(ctx, id).Bytes()
	if err != nil {
		return err
	}

	data := toJson1(val)
	return c.JSON(fiber.Map{"cached": data})
}

//------------------------------------------------------------------------------
//gets product from db api and store it in redis
// func AddProductCache(c *fiber.Ctx) error {

// 	//product := new(ProductSample)
// 	// product := new(models.Product)
// 	// if err := c.BodyParser(product); err != nil {
// 	// 	return err
// 	// }
// 	// return c.JSON(&product)

// 	res, err := http.Get("http://localhost:3000/api/getProduct/4")
// 	if err != nil {
// 		return err
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return err
// 	}

// 	var product ProductSample
// 	parseErr := json.Unmarshal(body, &product)
// 	if parseErr != nil {
// 		return parseErr
// 	}

// 	//to store data in redis with id "id"
// 	id := strconv.FormatUint(uint64(product.Id), 10)
// 	cacheErr := cache.Set(ctx, id, body, 0).Err()
// 	if cacheErr != nil {
// 		return cacheErr
// 	}

// 	//to retrieve data from redis with it "id"
// 	val, err := cache.Get(ctx, id).Bytes()
// 	if err != nil {
// 		return err
// 	}

// 	data := toJson(val)
// 	return c.JSON(fiber.Map{"cached": data})

// }

//--------------------------------------------------------------------------------
//to convert the byte arr to productsample arr
// func toJson(val []byte) []ProductSample {
// 	prod := []ProductSample{}
// 	err := json.Unmarshal([]byte(val), &prod)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return prod
// }

//--------------------------------------------------------------------------------
//sample fn to get all the data from redis and convert it to suitable json format
// func ProductCache(c *fiber.Ctx) error {
// 	var keys []string
// 	var err error
// 	var cursor uint64

// 	//getting all the keys form the redis db
// 	keys, _, err = cache.Scan(ctx, cursor, "*", 10).Result()
// 	if err != nil {
// 		panic(err)
// 	}

// 	//getting all the values form the redis db
// 	var allProduct []ProductSample
// 	for _, keyval := range keys {
// 		if keyval != "productData" {
// 			val, err := cache.Get(ctx, keyval).Bytes()
// 			if err != nil {
// 				return err
// 			}

// 			iterdata := toJson1(val)
// 			allProduct = append(allProduct, iterdata)
// 		}
// 	}

// 	return c.JSON(&allProduct)
// }

//---------------------------------------------------------------------------------
//gets all the data from the db and store as a single key value in redis
// func GetAllProductCache(c *fiber.Ctx) error {

// 	id := "productData"

// 	//to retrieve data from redis with it "id"
// 	val, err := cache.Get(ctx, id).Bytes()
// 	if err != nil {
// 		return err
// 	}

// 	data := toJson(val)
// 	return c.JSON(&data)

// }

//-------------------------------------------------------------------------
//used to update the entire data we stored as a single key-value pair from above function
// func UpdateProductCache(c *fiber.Ctx) error {
// 	res, err := http.Get("http://localhost:3000/api/getAllProducts")
// 	if err != nil {
// 		return err
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return err
// 	}

// 	var product []ProductSample
// 	parseErr := json.Unmarshal(body, &product)
// 	if parseErr != nil {
// 		return parseErr
// 	}

// 	//to store data in redis with id "id"
// 	// id := strconv.FormatUint(uint64(product.Id), 10)
// 	id := "productData"
// 	cacheErr := cache.Set(ctx, id, body, 0).Err()
// 	if cacheErr != nil {
// 		return cacheErr
// 	}

// 	//to retrieve data from redis with it "id"
// 	val, err := cache.Get(ctx, id).Bytes()
// 	if err != nil {
// 		return err
// 	}

// 	data := toJson(val)
// 	return c.JSON(&data)
// }
