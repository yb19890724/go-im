package model

import (
	orm "../database"
)

type User struct {
	ID       uint64 `json:"id"`        // 列名为 `id`
	Username string `json:"username"`  // 列名为 `username`
	Password string `json:"password"`  // 列名为 `password`
	Avatar   string `json:"avatar"`    // 列名为  `avatar`
	Token    string `json:"token"`     // 列名为  `token`
	Nickname string `json:"nick_name"` // 列名为  `token`
}

var Users []User

//  添加
func (user User) Insert() (id uint64, err error) {

	//  添加数据
	result := orm.Eloquent.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//  列表
func (user *User) Users() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

//  修改
func (user *User) Update(id uint64) (error) {

	var updateUser User
	
	if err := orm.Eloquent.Select([]string{"id"}).First(&updateUser, id).Error; err != nil {
		return err
	}

	//  参数1:是要修改的数据
	//  参数2:是修改的数据
	if err := orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

//  删除数据
func (user *User) Destroy(id uint64) (Result User, err error) {

	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}

// 登录
func (user *User) Login() (err error) {
	if err = orm.Eloquent.Select([]string{"id"}).
		Where("username = ?", user.Username).
		Where("password = ?", user.Password).
		First(&user).Error; err != nil {
		return
	}
	return
}
