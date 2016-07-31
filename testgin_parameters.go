package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"os/exec"
)

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		cmd := "cat /etc/passwd"
		out, _ := exec.Command("sh", "-c", cmd).CombinedOutput()
		//name := c.Param("name")
		//action := c.Param("action")
		//message := name + " is " + action
		c.String(http.StatusOK, string(out))
	})

	router.Run(":8080")
}
