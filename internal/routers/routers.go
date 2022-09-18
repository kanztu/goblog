package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kanztu/goblog/internal/controller"
)

func NewRouters() (r *gin.Engine) {
	r = gin.Default()
	r.Use(gin.Recovery())

	blogController := controller.NewBlogController()
	pageController := controller.NewPageController()
	tagController := controller.NewTagController()

	r.LoadHTMLGlob("goblog-frontend/template/*")

	page := r.Group("blog")
	{
		page.GET("/", pageController.GetIndexPage)
		page.GET("/id/:id", pageController.GetBlogPage)
		page.GET("/tag/:id", pageController.GetBlogByTagPage)
		page.GET("/tag", pageController.GetTagPage)
		page.GET("/about", pageController.GetAboutPage)
		page.Static("/static", "./goblog-frontend")
		page.Static("/wasm", "./wasm")
	}

	api := r.Group("api/v1")
	{
		api.GET("blogs", blogController.GetBlog)
		api.GET("blogs/tag", blogController.GetBlogWithTag)
		api.GET("tags", tagController.GetTag)
		api.GET("pagecata", pageController.GetPageCata)

		// admin := api.Group("admin")
		// {
		// 	admin.POST("blogs", blogController.CreateBlog)
		// 	admin.PUT("blogs")
		// 	admin.DELETE("blogs", blogController.DeleteBlog)

		// 	admin.POST("tags")
		// 	admin.DELETE("tags")

		// 	admin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		// }
	}

	// Use nginx to provide basic auth for this route group

	return
}
