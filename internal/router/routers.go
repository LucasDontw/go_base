package router

import (
	"cms/v2/internal/middleware"
	"cms/v2/internal/services"

	"github.com/gin-gonic/gin"
)

const (
	rootPath   = "/api/"
	noAuthPath = "/out/api/"
)

func CmsRouters(r *gin.Engine) {
	cmsApp := services.NewCmsApp()

	session := middleware.NewSessionAuth()

	root := r.Group(rootPath).Use(session.Auth) //api分組
	{
		// /api/cms/hello
		root.GET("/cms/hello", cmsApp.Hello)

		// /api/cms/content/delete
		root.GET("/cms/content/find", cmsApp.ContentFind)

		// /api/cms/content/create
		root.POST("/cms/content/create", cmsApp.ContentCreate)

		// /api/cms/content/update
		root.POST("/cms/content/update", cmsApp.ContentUpdate)

		// /api/cms/content/delete
		root.DELETE("/cms/content/delete", cmsApp.ContentDelete)
	}

	noAuth := r.Group(noAuthPath)
	{
		// /out/api/cms/register
		noAuth.POST("/cms/register", cmsApp.Register)
		// /out/api/cms/login
		noAuth.POST("/cms/login", cmsApp.Login)
	}
}
