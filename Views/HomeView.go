package Views

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HomeView(c *gin.Context) {
	obj := make(map[string]string)
	obj["title"] = "Hello gin-gonic"
	c.HTML(http.StatusOK, "home.html",obj)
}