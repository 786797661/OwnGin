/**
* @Author: Gosin
* @Date: 2021/4/11 22:37
 */

package chapter02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type UserInfo struct {
	Id       int    `form:"id"`
	Age      int    `json:"age" form:"age"`
	Name     string `json:"username" form:"username"`
	Addr     string `json:"addr" form:"addr"`
	Password string `json:"password"  form:"password"`
}

func UserInfoStruct(context *gin.Context) {
	user := UserInfo{Id: 1, Age: 19, Name: "xm", Addr: "xxx"}
	context.HTML(http.StatusOK, "user/user_info.html", user)
}

func ArrayControl(context *gin.Context) {
	var arr = [3]int{1, 2, 3}
	context.HTML(http.StatusOK, "arr/arr_info.html", arr)
}

func MapControl(context *gin.Context) {
	var map_data = make(map[string]string)
	map_data["Name"] = "1"
	map_data["Age"] = "12"
	map_data["Url"] = "map/map_info.html"
	var map_data2 = make(map[string]string)
	map_data2["Name"] = "2"
	map_data2["Age"] = "22"
	var map_map_data = make(map[string]map[string]string)
	map_map_data["map1"] = map_data
	map_map_data["map2"] = map_data2
	context.HTML(http.StatusOK, "map/map_info.html", map_map_data)
}

func MapStructControl(context *gin.Context) {
	user := UserInfo{Id: 1, Age: 19, Name: "xm", Addr: "xxx"}
	user2 := UserInfo{Id: 2, Age: 20, Name: "xh", Addr: "xxx"}
	var map_data = make(map[string]UserInfo)
	map_data["user1"] = user
	map_data["user2"] = user2
	//var map_data2=make(map[string]string)
	//map_data2["Name"]="2"
	//map_data2["Age"]="22"
	//var map_map_data=make(map[string]map[string]string)
	//map_map_data["map1"]=map_data
	//map_map_data["map2"]=map_data2
	context.HTML(http.StatusOK, "map/map_info.html", map_data)
}

func SliceControl(context *gin.Context) {
	var slice = make([]string, 4)
	slice = append(slice, "1", "2", "3", "4", "5", "6", "8")
	context.HTML(http.StatusOK, "arr/arr_info.html", slice)
}

func Pargram(context *gin.Context) {
	context.String(http.StatusOK, context.Param("id")+context.Param("name"))
}

func ArrPargram(context *gin.Context) {
	strs, _ := context.GetQueryArray("id")
	context.String(http.StatusOK, "hello,%s", strs)
}

func MapPargram(context *gin.Context) {
	strs, _ := context.GetQueryMap("id")
	context.String(http.StatusOK, "hello,%s", strs)
}

func GetUserController(context *gin.Context) {
	context.HTML(http.StatusOK, "user/user_do_info.html", nil)
}

func PostUser(context *gin.Context) {
	context.PostForm("username")
	context.PostForm("password")
	context.String(http.StatusOK, context.PostForm("username")+"/"+context.PostForm("password"))
}

func GetUserController2(context *gin.Context) {
	context.HTML(http.StatusOK, "user/user_do_ajax.html", nil)
}

func PostUser2(context *gin.Context) {
	context.PostForm("username")
	context.PostForm("password")
	context.PostForm("age")
	context.JSON(http.StatusOK, gin.H{"code": "200", "msg": "成功"})

}

func GetUserController3(context *gin.Context) {
	context.HTML(http.StatusOK, "user/user_do_struct.html", nil)
}

func PostUser3(context *gin.Context) {
	user := new(UserInfo)
	err := context.ShouldBind(&user)
	fmt.Println(err)
	fmt.Println(user)
	context.JSON(http.StatusOK, gin.H{"code": "200", "msg": "成功"})

}

func GetUploadFileController(context *gin.Context) {
	context.HTML(http.StatusOK, "user/upload_to_file.html", nil)
}

func PostUpload(context *gin.Context) {
	file, _ := context.FormFile("file")
	fmt.Print(file.Filename)
	time_unix_int := time.Now().Unix()
	time_unix_str := strconv.FormatInt(time_unix_int, 10)
	dst := "upload/" + time_unix_str + file.Filename
	err := context.SaveUploadedFile(file, dst)
	if err != nil {
		fmt.Println("上传失败")
		context.String(http.StatusOK, err.Error())
	}

	context.String(http.StatusOK, "上传成功")
}

func GetUploadFileController2(context *gin.Context) {
	context.HTML(http.StatusOK, "user/upload_to_file2.html", nil)
}

func PostUpload2(context *gin.Context) {
	form, _ := context.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		fmt.Print(file.Filename)
		time_unix_int := time.Now().Unix()
		time_unix_str := strconv.FormatInt(time_unix_int, 10)
		dst := "upload/" + time_unix_str + file.Filename
		err := context.SaveUploadedFile(file, dst)
		if err != nil {
			fmt.Println("上传失败")
			context.String(http.StatusOK, err.Error())
		}

		context.String(http.StatusOK, "上传成功")
	}
}

func GetUploadFileController3(context *gin.Context) {
	context.HTML(http.StatusOK, "user/upload_to_file3.html", nil)
}
