package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kanztu/goblog/internal/controller"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouters() (r *gin.Engine) {
	r = gin.Default()
	r.Use(gin.Recovery())

	blogController := controller.NewBlogController()
	pageController := controller.NewPageController()
	tagController := controller.NewTagController()

	page := r.Group("blog")
	{
		page.GET("/", pageController.GetIndexPage)
		page.Static("css", "src/css")
		page.Static("js", "src/js")
		page.Static("images", "src/images")
		page.Static("wasm", "src/wasm")
	}

	api := r.Group("api/v1")
	{
		api.GET("blogs", blogController.GetBlog)
		api.POST("blogs", blogController.CreateBlog)
		api.PUT("blogs")
		api.DELETE("blogs", blogController.DeleteBlog)

		api.GET("blogs/tag", blogController.GetBlogWithTag)

		api.GET("tags", tagController.GetTag)
		api.POST("tags")
		api.DELETE("tags")

		api.GET("pagecata", pageController.GetPageCata)
	}

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return
}
