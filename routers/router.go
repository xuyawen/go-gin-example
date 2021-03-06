package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/xuyawen/go-gin-example/docs"

	jwt "github.com/xuyawen/go-gin-example/middleware"
	"github.com/xuyawen/go-gin-example/pkg/setting"
	"github.com/xuyawen/go-gin-example/routers/api"
	v1 "github.com/xuyawen/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取指定标签/全部标签
		apiv1.GET("/tags", v1.GetTags)
		// 添加标签
		apiv1.POST("/tags", v1.AddTag)
		// 编辑标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		// 获取指定文章
		apiv1.POST("/articles/:id", v1.GetArticle)
		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 添加文章
		apiv1.POST("/articles", v1.AddArticle)
		// 编辑文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
