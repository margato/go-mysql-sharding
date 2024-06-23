package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(customerRepository *CustomerRepository) *gin.Engine {
	router := gin.Default()

	router.GET("/customers/:id", func(c *gin.Context) {
		id := c.Param("id")

		customer, err := customerRepository.GetById(id)

		if err != nil {
			if errors.Is(err, errors.New("customer not found")) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, customer)
	})

	router.POST("/customers", func(c *gin.Context) {
		var data map[string]interface{}
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		customer := NewCustomer(data["name"].(string))
		err := customerRepository.Save(customer)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, customer)
	})

	return router
}
