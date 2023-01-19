package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handle_root_get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "sysd_http running");
}

func get_services(c *gin.Context) {

}

func get_service_by_name(c *gin.Context) {
	name := c.Param("name")

	c.IndentedJSON(http.StatusOK, name)
}

func main() {
	router := gin.Default()
	router.GET("/", handle_root_get)

	router.GET("/service", get_services)
	router.GET("/service/<name>", get_service_by_name)

	router.Run("localhost:8000")
}