package Controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SayHelloForm struct {
    Name     string `form:"name" binding:"required"`
    Surname     string `form:"surname" binding:"required"`
}

func SayHelloController(c *gin.Context) {
	var form SayHelloForm
	c.BindWith(&form, binding.MultipartForm)

	obj := make(map[string]string)
	obj["name"] = form.Name

	c.HTML(http.StatusOK, "sayhello.html",obj)
}