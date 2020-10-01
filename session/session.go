package session


import (
	"XChatRoom/data/dataImpl"
	"container/list"
)

//登陆用户Map
var LoginUserSessionsMap map[int]*dataImpl.LoginUserSession

//登陆用户List
var LoginUserSessionsList *list.List
