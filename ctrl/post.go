package ctrl

import (
	"Calculator/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Getstring(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	//c.String(http.StatusOK,"%s",service.Calculate(name))
	c.JSON(200, service.Calculate(service.Convert(name)))
}
