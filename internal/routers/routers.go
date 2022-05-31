package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kanztu/goblog/internal/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouters() (r *gin.Engine) {
	r = gin.Default()
	r.Use(gin.Recovery())

	blogController := controller.NewBlogController()
	pageController := controller.NewPageController()
	tagController := controller.NewTagController()

	r.LoadHTMLGlob("template/*")

	page := r.Group("blog")
	{
		page.GET("/", pageController.GetIndexPage)
		page.GET("/id/:id", pageController.GetBlogPage)
		page.GET("/tag/:id", pageController.GetBlogByTagPage)
		page.GET("/tag", pageController.GetTagPage)
		page.GET("/about", pageController.GetAboutPage)
		page.Static("static", "static")
	}

	api := r.Group("api/v1")
	{
		api.GET("blogs", blogController.GetBlog)
		api.GET("blogs/tag", blogController.GetBlogWithTag)
		api.GET("tags", tagController.GetTag)
		api.GET("pagecata", pageController.GetPageCata)
	}

	admin := r.Group("admin")
	{
		api := admin.Group("api/v1")
		{
			api.POST("blogs", blogController.CreateBlog)
			api.PUT("blogs")
			api.DELETE("blogs", blogController.DeleteBlog)

			api.POST("tags")
			api.DELETE("tags")

		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return
}
