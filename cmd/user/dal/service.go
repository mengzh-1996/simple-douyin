package dal

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"simple-douyin/cmd/user/model"
	"simple-douyin/cmd/user/utils"
)

// 通过用户结构体查询用户
func SearchUserService(searchUser *model.User) model.User {
	db := GetDB()
	var user model.User
	result := db.Where(&searchUser).First(&user)
	if result.Error != nil {
		panic("该用户不存在")
	}
	return user
}

// 通过用户ID查询用户
func SearchUserByIDService(userID uint) model.User {
	db := GetDB()
	var user model.User
	result := db.Where(&model.User{UserID: userID}).First(&user)
	if result.Error != nil {
		panic("该用户不存在")
	}
	return user
}

// 注册
func RegisterService(regisUser *model.User) {
	db := GetDB()

	if err := db.Create(&regisUser).Error; err != nil {
		panic("用户注册失败！")
	}
}

// 登录
func LoginService(loginUser *model.User, tokenstr *string) {
	db := GetDB()
	var user model.User
	if result := db.Where("user_name=?", loginUser.UserName).First(&user); result.Error != nil {
		panic("用户不存在")
	}
	if loginUser.EncryptedPassword != user.EncryptedPassword {
		panic("密码错误！")
	}
	*tokenstr = utils.GetToken(user.UserID)
}

// A关注B
func FollowService(userA, userB *model.User) {
	db := GetDB()

	db.Model(&userB).Association("Fans").Append(userA)

}

// A取关B
func UnFollowService(userA, userB *model.User) {

	db := GetDB()

	db.Model(&userB).Association("Fans").Delete(userA)

}

// 查询粉丝列表
func SearchFansService(user *model.User) (fans []int) {
	db := GetDB()
	rows, err := db.Table("follow").Where("user_user_id = ?", user.UserID).Select("fan_user_id").Rows()
	if err != nil {
		fmt.Println("Counting Following error:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		rows.Scan(&id)
		fans = append(fans, id)
	}
	return
}

// 查询用户粉丝数
func SearchFansNum(user *model.User) int {
	num := SearchFansService(user)
	return len(num)
}

// 查询关注列表
func SearchFavsService(user *model.User) (fans []int) {
	db := GetDB()
	rows, err := db.Table("follow").Where("fan_user_id = ?", user.UserID).Select("user_user_id").Rows()
	if err != nil {
		fmt.Println("Counting Following error:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		rows.Scan(&id)
		fans = append(fans, id)
	}
	return
}

// 查询用户关注数
func SearchFavssNum(user *model.User) int {
	num := SearchFavsService(user)
	return len(num)
}

// 查询用户A是否关注用户B
func SearchFollowService(userA, userB *model.User) (flag bool) {
	db := GetDB()
	rows, err := db.Table("follow").Where("user_user_id = ? AND fan_user_id= ?", userA.UserID, userB.UserID).Rows()
	flag = !errors.Is(err, gorm.ErrRecordNotFound)
	defer rows.Close()
	return
}
