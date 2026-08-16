package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type sourceDocument struct {
	Value string `json:"value"`
}

func observeSource(writer http.ResponseWriter, value string) {
	encoded := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "\"", "&quot;", "'", "&#39;").Replace(value)
	_, _ = fmt.Fprint(writer, encoded)
}

func registerSourceRoutes(router *gin.Engine) {
	router.GET("/sources/query", func(context *gin.Context) {
		observeSource(context.Writer, context.Query("value"))
	})
	router.GET("/sources/path/:value", func(context *gin.Context) {
		observeSource(context.Writer, context.Param("value"))
	})
	router.POST("/sources/form", func(context *gin.Context) {
		observeSource(context.Writer, context.PostForm("value"))
	})
	router.POST("/sources/json", func(context *gin.Context) {
		var value sourceDocument
		_ = context.ShouldBindJSON(&value)
		observeSource(context.Writer, value.Value)
	})
	router.GET("/sources/header", func(context *gin.Context) {
		observeSource(context.Writer, context.GetHeader("X-Rig-Source"))
	})
	router.GET("/sources/cookie", func(context *gin.Context) {
		value, _ := context.Cookie("rig_source")
		observeSource(context.Writer, value)
	})
	router.POST("/sources/multipart", func(context *gin.Context) {
		_, _ = context.FormFile("value")
		observeSource(context.Writer, context.PostForm("value"))
	})
	router.POST("/sources/body", func(context *gin.Context) {
		value, _ := io.ReadAll(http.MaxBytesReader(context.Writer, context.Request.Body, 64<<10))
		observeSource(context.Writer, string(value))
	})
	router.GET("/sources/middleware", func(context *gin.Context) {
		observeSource(context.Writer, context.GetHeader("X-Rig-Middleware"))
	})
	router.GET("/sources/context", func(context *gin.Context) {
		observeSource(context.Writer, context.GetHeader("X-Rig-Context"))
	})
	router.GET("/sources/principal", func(context *gin.Context) {
		observeSource(context.Writer, strings.TrimPrefix(context.GetHeader("Authorization"), "Bearer "))
	})
}
