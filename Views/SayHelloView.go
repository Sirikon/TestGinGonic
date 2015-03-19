package Views

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SayHelloView(c *gin.Context) {
	obj := c.MustGet("data").(map[string]string)
	c.HTML(http.StatusOK, "sayhello.html",obj)
}