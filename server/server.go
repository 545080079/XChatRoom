package server

import (
	"XChatRoom/controller"
	"XChatRoom/data/dataImpl"
	"XChatRoom/model"
	"XChatRoom/session"
	"container/list"
	"github.com/gin-gonic/gin"
)


var StatusUserNameError int = 2
var StatusPwdError int = 1
var StatusNoneError int = 0
var StatusSendError int = 3


func init() {
	//初始化登陆Session
	session.LoginUserSessionsMap = make(map[int]*dataImpl.LoginUserSession)
	session.LoginUserSessionsList = list.New()

	//初始化Redis
	model.Client = model.CreateClient()
}


func Begin() {
    server := gin.Default()

    v1 := server.Group("/v1")
    {

    	/*
    		测试API
    	 */
    	test := v1.Group("/test")
    	{
    		test.GET("/message/add", controller.ProcessAdd)
    		test.GET("/message/consume", controller.ProcessConsumeMessage)
			test.GET("/hello", controller.ProcessHello)
		}

		/*
			POST API
		 */
		post := v1.Group("/post")
		{
			post.POST("/send", controller.ProcessSendMessage)
			post.POST("/receive", controller.ProcessReceiveMessage)
			post.POST("/login", controller.ProcessLogin)
		}

	}


    server.Run()
}

