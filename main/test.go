package main

//
//func main() {
//	m := &dataImpl.Message{
//		Id:      1,
//		Type:    "test1",
//		Content: "This is a globelTest message.",
//		Source:  101,
//		Target:  102,
//	}
//	fmt.Println(m.String())
//
//	u := &dataImpl.User{
//		Id:             101,
//		Type:           0,
//		Name:           "Xi Ming",
//		Password:       "admin",
//		Phone:          "13899999999",
//		SecondPassword: "root",
//		ListUserGroupId: []int64{0,1},
//	}
//	fmt.Println(u.String())
//
//	//ug := &dataImpl.UserGroup{
//	//	Id:         0,
//	//	Type:       0,
//	//	Size:       0,
//	//	Name:       "Group_Xiamen",
//	//	ListUserId: []int64{1,2},
//	//}
//
//
//	data, err := proto.Marshal(m)
//	if err != nil {
//		log.Fatal("marshaling error: ", err)
//	}
//	m2 := &dataImpl.Message{}
//	proto.Unmarshal(data, m2)
//	fmt.Println("Unmarshal ::")
//	fmt.Println(m2)
//
//}