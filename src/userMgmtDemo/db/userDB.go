package db

import (
	"github.com/jinzhu/gorm"
	"userMgmtDemo/model"
)

//query * from users?
func (baseGorm *BaseGorm) QueryAllUserInfo() (userInfo []model.Users, err error) {
	if err := baseGorm.GetDB().Model(&model.Users{}).Find(&userInfo).Error; err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

//query data from users by username and pwd
func (baseGorm *BaseGorm) QueryUserInfoByUnamePwd(userName,password string) (userInfo model.Users, count int,err error) {
	if err = baseGorm.GetDB().Table("users").Where("name = ? AND password = ?", userName,password).Find(&userInfo).Count(&count).Error; err != nil {
		return  userInfo, count,err
	}
	return userInfo,count, err
}

//query data from users by username
func (baseGorm *BaseGorm) QueryUserInfoByUname(userName string) (count int,err error) {
	if err = baseGorm.GetDB().Table("users").Where("userName = ?", userName).Count(&count).Error; err != nil {
		return  count,err
	}
	return count, err
}



//query data from Users where Id=?
func (baseGorm *BaseGorm) QueryUserInfoById(userId int) (userInfo model.Users, err error) {
	if err = baseGorm.GetDB().Where("Id = ?", userId).Find(&userInfo).Error; err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

//delete data from Users where Id=?
func DeleteUserInfo(db *gorm.DB, userId int) (err error) {
	if err = db.Where("Id = ?", userId).Delete(model.Users{}).Error; err != nil {
		return err
	}
	return nil
}

//update Users
func UpdateUserInfo(db *gorm.DB, userInfo model.Users) (err error) {
	if err = db.Save(&userInfo).Error; err != nil {
		return err
	}
	return nil
}

// insert Users
func InsertUserInfo(db *gorm.DB,userInfo model.Users)(err error){
	if err = db.Save(&userInfo).Error; err != nil {
		return err
	}
	return nil
}

//update Users password
func UpdateUserInfoToken(db *gorm.DB,  userInfo model.Users) (err error) {
	if err = db.Model(&userInfo).Updates(map[string]interface{}{"Password": userInfo.Password}).Error; err != nil {
		return err
	}
	return nil
}

//insert data into users
func CreateUserInfo(db *gorm.DB, userInfo model.Users) (err error) {
	if err = db.Create(&userInfo).Error; err != nil {
		return err
	}
	return nil
}
