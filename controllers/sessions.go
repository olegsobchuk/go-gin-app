package controllers

import (
	"github.com/gin-gonic/gin"
)

// Root root page
func Root(c *gin.Context) {
	c.HTML(200, "base", gin.H{
		"title": "Root page",
	})
}

// LogIn login user page
func LogIn(c *gin.Context) {
	c.HTML(200, "login", gin.H{
		"title": "LogIn page",
	})
}
