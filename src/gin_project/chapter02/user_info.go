/**
* @Author: Gosin
* @Date: 2021/4/11 22:37
 */

package chapter02

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id   int
	Age  int
	Name string
	Addr string
}

func UserInfoStruct(context *gin.Context) {
	user := UserInfo{Id: 1, Age: 19, Name: "xm", Addr: "xxx"}
	context.HTML(http.StatusOK, "user/user_info.html", user)
}
