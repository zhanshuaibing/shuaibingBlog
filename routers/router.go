package routers

import (
	"github.com/gin-gonic/gin"
	"shuaibingBlog/controllers"
)

func RegisterRouters(e *gin.Engine) {
	// Ping test
	e.GET("/index", controllers.Index)

	// Ping test
	e.GET("/ping", controllers.Ping)

	//// Get user value
	e.GET("/user/:name", controllers.GetUserByName)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := e.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		c.MustGet(gin.AuthUserKey)
		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			c.JSON(200, gin.H{"status": json.Value})
		}

	})

}