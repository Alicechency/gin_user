package routers

import (
	"gin_login/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() * gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("views/*")

	//定义加密（将session信息存储在redis服务器）
	//store1, _ := redis.NewStore(10, "tcp", "172.200.1.254:6379", "", []byte("secret"))
	// 设置session midddleware
	store := cookie.NewStore([]byte("loginuser"))
	router.Use(sessions.Sessions("mysession", store))
	{
		// 注册：
		router.GET("/register", controllers.RegisterGet)
		router.POST("/register", controllers.RegisterPost)

		// 登录
		router.GET("/login", controllers.LoginGet)
		router.POST("/login", controllers.LoginPost)

		//修改用户信息
		router.GET("/user/update", controllers.UpdateGet)
		router.POST("/user/update", controllers.UpdatePost)


	}

	//绑定session中间件
	//router.Use(sessions.Sessions("mySession", store1))

	//定义GET方法
	//router.GET("/session", func(context *gin.Context) {
	//	//初始化session对象
	//	session := sessions.Default(context)
	//
	//	//如果浏览器第一次访问返回状态码401，第二次访问则返回状态码200
	//	if session.Get("user") != "yinzhengjie" {
	//		session.Set("user", "yinzhengjie")
	//		session.Save()
	//		context.JSON(http.StatusUnauthorized, gin.H{"user": session.Get("user")})
	//	} else {
	//		context.String(http.StatusOK, "Successful second visit")
	//	}
	//
	//})

	return router
}
