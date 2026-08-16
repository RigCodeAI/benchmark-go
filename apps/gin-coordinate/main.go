package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/sources/query", func(context *gin.Context) {
		context.Status(http.StatusOK)
		_, _ = fmt.Fprint(context.Writer, context.Query("value"))
	})
	_ = router.Run(":8080")
}
