package controllers

import (
	"fmt"
	"gin_login/utils"
	"time"

	//"fmt"
	"gin_login/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func UpdateGet(c *gin.Context)  {
	//获取session信息
	islogin := GetSession(c)


	c.HTML(http.StatusOK, "update.html", gin.H{"IsLogin": islogin})
}

func UpdatePost(c *gin.Context){
	//get data
	password := c.PostForm("password")
	session := sessions.Default(c)
	loginuser := session.Get("loginuser")
	id := models.QueryUserWithUsername(loginuser.(string))
	fmt.Println("username:", loginuser.(string), ",id:", id)

	islogin := GetSession(c)

	if islogin!=true{
		c.JSON(http.StatusOK, gin.H{"code":0,"message":"请先登录"})
	}
	//var user models.User
	//	//
	//	//db,err := gorm.Open("mysql","root:961013@(127.0.0.1:3306)/gin_login?charset=utf8")
	//	//if err != nil{
	//	//	panic(err)
	//	//}
	//	//defer db.Close()
	//	//
	//	//db.Where("id = ?",id).Find(&user)
	//	//user.PassWord = password
	//	//db.Model(&user).Update("password",password)
	art := models.User{id, loginuser.(string), utils.MD5(password), " ", time.Now().Unix()}
	_, err := models.UpdateUser(art)


	//返回数据给浏览器
	response := gin.H{}
	if err == nil {
		//无误
		response = gin.H{"code": 1, "message": "更新成功"}
	} else {
		response = gin.H{"code": 0, "message": "更新失败"}
	}

	c.JSON(http.StatusOK, response)

}
