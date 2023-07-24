package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
)

var products = []product{}

func main() {

	fmt.Println("Time now: ", time.Now())
	seed()
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.POST("/products", postProducts)

	router.Run("localhost:8080")
}

type product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

func seed() {
	faker := faker.New()

	for i := 0; i < 10; i++ {
		products = append(products, product{ID: strconv.Itoa(i), Name: faker.Person().FirstName(), Description: faker.Lorem().Word(), Price: 2.99, ImageUrl: faker.ProfileImage().Image().Name()})
	}

	fmt.Println(products)

}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func postProducts(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
