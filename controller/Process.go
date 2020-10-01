package controller

import (
	"XChatRoom/data/dataImpl"
	"XChatRoom/messageQueue"
	"XChatRoom/model"
	"XChatRoom/session"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


var StatusUserNameError int = 2
var StatusPwdError int = 1
var StatusNoneError int = 0
var StatusSendError int = 3


func ProcessAdd(context *gin.Context) {
	mes := dataImpl.Message{
		Id:      0,
		Type:    "normal",
		Content: "This is globelTest message.",
		Source:  0,
		Target:  1,
	}
	model.AddMessage(model.Client, mes)
	//return json
	context.JSON(http.StatusOK, gin.H{
		"status" : StatusNoneError,//send status
		"target" : 1,
		"content" : "[Add OK]This is globelTest message.",
	})

}



func ProcessConsumeMessage(context *gin.Context) {
	res := model.ConsumeMessage(model.Client, "lyt")
	//return json
	context.JSON(http.StatusOK, gin.H{
		"status" : StatusNoneError,//send status
		"target" : "lyt",
		"content" : "[Add OK]" + string(res[:]),
	})

}

func ProcessSendMessage(context *gin.Context) {
	username := context.PostForm("username")

	//mock有消息
	//from := "default"//系统默认发件人
	//content := "this a mock message."

	from := ""
	content := ""
	fmt.Printf("username = %s\n", context.PostForm("username"))
	//读取RabbitMQ的[Exchange:content]中的指向[收件人:username]的消息

	msg := model.ConsumeMessage(model.Client, username)

	//读取消息list
	//for it := msgs.Front(); it != nil; it = it.Next() {
	//	msg := (it.Value).([]byte)
	//
	//	//解析格式——发件人:消息内容
	//	msgStr := string(msg[:])
	//	splitStr := strings.Split(msgStr, ":")
	//	//splitStr[0]:发件人
	//	//splitStr[1]:消息内容
	//	fmt.Printf("splitStr[0] = %s, splitStr[1] = %s", splitStr[0], splitStr[1])
	//	from = splitStr[0]
	//
	//	if strings.EqualFold(content, "") {
	//		content = splitStr[0] + ":" + splitStr[1]
	//	} else {
	//		content = content + ";" + splitStr[0] + ":" + splitStr[1]	//追加到content，分号代表一条消息的结束
	//	}
	//}

	//读取单条消息（简化）
	msgStr := string(msg[:])
	splitStr := strings.Split(msgStr, ":")
	//splitStr[0]:发件人
	//splitStr[1]:消息内容
	fmt.Printf("splitStr[0] = %s, splitStr[1] = %s", splitStr[0], splitStr[1])
	from = splitStr[0]

	if strings.EqualFold(content, "") {
		content = splitStr[0] + ":" + splitStr[1]
	} else {
		content = content + ";" + splitStr[0] + ":" + splitStr[1]	//追加到content，分号代表一条消息的结束
	}

	//return json
	context.JSON(http.StatusOK, gin.H{
		"status" : StatusNoneError,//send status
		"target" : username,//收件人
		"content" : content,
		"from" : from,//发件人
	})
}

func ProcessLogin(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	fmt.Println("POST login username password:", username, password)


	//登陆校验
	res := model.LoginCheck(model.Client, username, password)
	if res == false {
		fmt.Println("Password Error!")
		context.JSON(http.StatusOK, gin.H{
			"status" : StatusPwdError,
		})
		return
	}


	userMessage := model.UserFindByName(model.Client, username)
	loginUserSession := &dataImpl.LoginUserSession{
		Id:      userMessage.Id,
		Type:    userMessage.Type,
		Name:    username,
		IsLogin: true,
	}
	//保存登陆信息
	session.LoginUserSessionsMap[int(userMessage.Id)] = loginUserSession
	session.LoginUserSessionsList.PushBack(int(userMessage.Id))


	fmt.Println("Login Success, Id ->", session.LoginUserSessionsMap[int(userMessage.Id)])
	messageQueue.BgRecMessage(userMessage.Id, userMessage.Name) //后台开启消息监听
	context.JSON(http.StatusOK, gin.H{
		"status" : StatusNoneError,
	})
}

func ProcessReceiveMessage(context *gin.Context) {
	target := context.PostForm("target")
	content := context.PostForm("content")
	from := context.PostForm("username")

	//log
	fmt.Println("提交请求: from->", from, "To->" + target + ", 内容->" + content)

	//发送消息
	fmt.Println("Call PublishMessage: from", from, "发送消息To->", target, ", 内容->", content)
	messageQueue.PublishMessage(from, target, content)

	//return json
	context.JSON(http.StatusOK, gin.H{
		"status" : StatusNoneError,//send status
		"target" : target,
		"content" : content,
	})
}

func ProcessHello(context *gin.Context) {
	context.String(http.StatusOK, "hello!")
}
