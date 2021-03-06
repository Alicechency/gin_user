package main

import (
	"fmt"
	db "gin_login/database"
	router "gin_login/routers"
	"gin_login/utils"
)

func main()  {

	err := db.InitMysql()
	if err != nil {
		 fmt.Println("initdb is failed,err:%v\n",err)
		return
	}else {
		fmt.Printf("initdb is succ")
	}

	// 路由
	router := router.InitRouter()

	// 注册一个全局中间件
	router.Use(utils.StatCost())

	// 静态资源
	router.Static("/static", "./static")


	router.Run(":8081")
}
