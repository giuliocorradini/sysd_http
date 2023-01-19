package main

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func AddServiceSuffixIfNeeded(service string) string {
	matched, _ := regexp.MatchString("^.*\\.service$", service)

	if matched {
		return service
	} else {
		return "" + service + ".service"
	}
}

func GetServiceByName(c *gin.Context) {
	name := c.Param("name")
	name = AddServiceSuffixIfNeeded(name)

	service := QueryService(name)

	c.IndentedJSON(http.StatusOK, service)
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "sysd_http running")
	})

	router.GET("/service/:name", GetServiceByName)

	router.Run("0.0.0.0:8000")
}
