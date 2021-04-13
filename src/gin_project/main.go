/**
* @Author: Gosin
* @Date: 2021/4/8 22:52
 */

package main

import (
	"gin_project/gin_project/chapter02"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//加载页面文件
	router.LoadHTMLGlob("template/**/*")
	//加载静态文件
	router.Static("/static", "static")
	router.GET("/user_info", chapter02.UserInfoStruct)
	//map页面
	router.GET("/map", chapter02.MapStructControl)
	//数组页面
	router.GET("/arr", chapter02.SliceControl)
	//加参页面
	//router.GET("/param/:id", chapter02.Pargram)
	router.GET("/param/*id", chapter02.Pargram)
	//router.GET("/param2", chapter02.QueryPargram)

	router.GET("/user_info_form", chapter02.GetUserController)
	router.POST("/user_add", chapter02.PostUser)

	router.GET("/user_info_form2", chapter02.GetUserController2)
	router.POST("/user_add2", chapter02.PostUser2)

	router.GET("/user_info_form3", chapter02.GetUserController3)
	router.POST("/user_add3", chapter02.PostUser3)

	router.Run(":9000")
}
