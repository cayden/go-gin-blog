// Package routers -----------------------------
// @file      : router.go
// @author    : cayden
// @contact   : cuiran2001@163.com
// @time      : 2022/6/16 13:56
// -------------------------------------------

package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog/pkg/setting"
	v1 "go-gin-blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)

		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r

}
