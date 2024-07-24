package main

import (
	"cms/v2/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.CmsRouters(r)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
