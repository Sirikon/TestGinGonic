package Controllers

import (
	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	obj := make(map[string]string)
	obj["title"] = "Hello gin-gonic"
	c.Set("data", obj)
}