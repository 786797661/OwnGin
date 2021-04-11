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
	router.Run(":9000")
}
