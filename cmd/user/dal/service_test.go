package dal

import (
	"fmt"
	"simple-douyin/cmd/user/model"
	"testing"
)

func TestService(t *testing.T) {
	// 结构体查询用户服务测试------------------------------------------------------
	// userRoot:=model.User{
	// 	UserName: "zm",
	// }
	// user:=SearchUserService(&userRoot)
	// fmt.Println(user)

	// 用户ID查询服务测试---------------------------------------------------------
	// user:=SearchUserByIDService(2)
	// fmt.Println(user)

	// 注册服务测试-----------------------------------------------------------------
	// user:=model.User{
	// 	UserName: "zm2",
	// 	EncryptedPassword: "zm2",
	// }
	// RegisterService(&user)

	// 登录服务测试----------------------------------------------------------------
	// user:=model.User{
	// 	UserName: "zm",
	// 	EncryptedPassword: "zm",
	// }
	// var tokenStr string
	// LoginService(&user,&tokenStr)
	// println(tokenStr)

	// 关注服务测试------------------------------------------------------------------------
	// userRoot:=model.User{
	// 	UserID: 1,
	// 	UserName: "zm",
	// 	EncryptedPassword: "zm",
	// }
	// user0:=model.User{
	// 	UserID: 2,
	// 	UserName: "zm0",
	// 	EncryptedPassword: "zm0",
	// }
	// FollowService(&user0,&userRoot)
	// fmt.Println(userRoot)

	userRoot := model.User{
		UserID:            1,
		UserName:          "zm",
		EncryptedPassword: "zm",
	}
	user0 := model.User{
		UserID:            3,
		UserName:          "zm1",
		EncryptedPassword: "zm1",
	}
	flag := SearchFollowService(&userRoot, &user0)
	fmt.Println(flag)
}
