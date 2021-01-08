package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	uriHealthCheck = "/healthz"

	goImportDomain = "nosoxon.net"
	goSourceRoot   = "https://github.com/nosoxon"
)

// Keys are prefixed with `nosoxon.net/'
var goSourceRedirects = map[string]string{
	"targetd-provisioner": "https://github.com/nosoxon/targetd-provisioner",
}

func main() {
	router := gin.New()
	router.LoadHTMLGlob("web/templates/*")
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter, uriHealthCheck))

	router.GET(uriHealthCheck, healthCheck)
	router.GET("/", func(c *gin.Context) {
		if c.Query("go-get") == "1" {
			c.HTML(http.StatusOK, "go-source-meta.tmpl", gin.H{
				"importPath": goImportDomain,
				"repoURI":    goSourceRoot,
			})
		} else {
			c.HTML(http.StatusOK, "main.tmpl", gin.H{"clientIP": c.ClientIP()})
		}
	})
	for path, uri := range goSourceRedirects {
		router.GET(fmt.Sprintf("/%v", path), goSourceRedirect(path, uri))
	}

	router.Run(":8080")
}

func healthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func goSourceRedirect(importPath, repoURI string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query("go-get") == "1" {
			c.HTML(http.StatusOK, "go-source-meta.tmpl", gin.H{
				"importPath": fmt.Sprintf("%v/%v", goImportDomain, importPath),
				"repoURI":    repoURI,
			})
		} else {
			c.Redirect(http.StatusMovedPermanently, repoURI)
		}
	}
}
