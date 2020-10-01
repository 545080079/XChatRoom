package model

import (
	"XChatRoom/data/dataImpl"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
)

var Client *redis.Client
var messageKeyPre string = "M:"
var userIdOffset uint64 = 100

//redis const
var StatusNullUserName int = 0
var StatusNullUserId int = 0


func CreateClient() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		Password:           "",
		DB:                 0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Redis Ping() no respond!")
		panic(err)
	}
	return client
}



/*
	Desc: 	用户注册
	Input: 	Redis实例, Username
	Output:	TrueOrFalse
*/
func Register() {

}



/*
	Desc: 	登陆验证
	Input: 	Redis实例, Username, Password
	Output:	TrueOrFalse
*/
func LoginCheck(client *redis.Client, name, password string) bool{

	fmt.Println("LoginCheck name password: ", name, password)
	user := UserFindByName(client, name)

	if user == nil {
		return false
	}

	fmt.Println("Redis ret: name = ", user.Name, "\t pwd = ", user.Password)
	fmt.Println(user.Password, ":", password, "?", strings.EqualFold(user.Password, password))
	if !strings.EqualFold(user.Password, password) {
		return false
	}
	return true

}


/*
	Desc: 	通过UserId查找昵称
	Input: 	Redis实例, UserId
	Output:	昵称（用户名）
*/
func UserNameFindById(client *redis.Client, id uint64) string{

	//redis数据key格式
	param := "User:" + string(id)
	username, err := client.HGet(param, "Username").Result()
	if err == redis.Nil {
		return "UserId not exist."
	}
	return username

}


/*
	Desc: 	通过昵称查找UserId
	Input: 	Redis实例, Username
	Output:	UserId
*/
func UserIdFindByName(client *redis.Client, name string) uint64{

	//查询UserName:[name]
	param := "UserName:" + name
	fmt.Println("Redis Get UserName:" + name)

	userId, err := client.Get(param).Result()
	if err == redis.Nil {
		return uint64(StatusNullUserId)
	}


	//string -> uint64
	ret, _ := strconv.ParseUint(userId, 10, 64)

	fmt.Println("UserIdFindByName: ret userId = ", ret)
	return ret
}



/*
	Desc: 	通过UserId查找User结构体
	Input: 	Redis实例, UserId
	Output:	User结构体指针对象
*/
func UserFindById(client *redis.Client, id uint64) *dataImpl.User{

	param := "User:" + strconv.FormatUint(id, 10)
	user, err := client.HGetAll(param).Result()
	if err == redis.Nil {
		fmt.Println("UserId not exist.")
		return nil
	}

	tType, _ := strconv.ParseUint(user["Type"], 10, 64)

	//将所有组号提取放入User结构体
	tSplitGroupId := strings.Split(user["ListUserGroupId"], ";")
	var tGroupId []uint64
	for i := 0; i < len(tSplitGroupId); i++ {
		t, _ := strconv.ParseUint(tSplitGroupId[i], 10, 64)
		tGroupId = append(tGroupId, t)
	}

	return &dataImpl.User{
		Id:              id,//uint64
		Type:            tType,//uint64
		Name:            user["Username"],
		Password:        user["Password"],
		Phone:           user["Phone"],
		SecondPassword:  user["SecondPassword"],
		ListUserGroupId: tGroupId,//[]uint64
	}

}


/*
	Desc: 	通过昵称查找User结构体
	Input: 	Redis实例, Username
	Output:	User结构体指针对象
 */
func UserFindByName(client *redis.Client, name string) *dataImpl.User{

	//Username	->	UserId
	userId := UserIdFindByName(client, name)

	//错误返回
	if userId == uint64(StatusNullUserId) {
		return nil
	}

	//Id -> User结构体
	res := UserFindById(client, userId)

	if res == nil {
		return nil
	}

	//return is the same as:

	//res := &dataImpl.User{
	//	Id:              userId,
	//	Type:            0,
	//	Name:            name,
	//	Password:        "root",
	//	Phone:           "13859984262",
	//	SecondPassword:  "nfmh0617",
	//	ListUserGroupId: nil,
	//}

	fmt.Println("UserFindByName: ret password = ", res.Password)
	return res
}

/*
	Desc: 消费消息（从Redis读取）
	Input:
		[username:待读取消息方的用户名]作为key查询Message<K,V>的Value
	Output:
		读取到的消息[]byte:	[username:content]
*/
func ConsumeMessage(client *redis.Client, username string) []byte{

	userId := UserIdFindByName(client, username)
	key := messageKeyPre + strconv.FormatUint(userId, 10)
	key = strings.TrimSpace(key)

	fmt.Println("Consume key =", key)
	messages, err := client.HGetAll(key).Result()
	if err == redis.Nil {
		fmt.Println("ConsumeMessageError: By Redis check, user's message null.")
		return []byte("")
	}

	fmt.Println("Consume content =", messages["content"])
	return []byte(username + ":" + messages["content"])

}

/*
	Desc: 存入消息
 */
func AddMessage(client *redis.Client, mes dataImpl.Message) bool {

    key := messageKeyPre + strconv.FormatUint(mes.Target + userIdOffset, 10)
    key = strings.TrimSpace(key)

	_, err := client.HSet(key, "id", mes.Id).Result()
	fmt.Println("HSet:", key)
	if err != nil {
		fmt.Println("redis AddMessage HSet mes.Target[id] Error:", err.Error())
		return false
	}
	client.HSet(key, "type", mes.Type)
	client.HSet(key, "content", mes.Content)
	client.HSet(key, "source", mes.Source)
	client.HSet(key, "target", mes.Target)
	return true

}