package controller

import (
	"github.com/ant0ine/go-json-rest/rest"
	"strconv"
	"userMgmtDemo/db"
	"log"
	"userMgmtDemo/model"
)

type UserController struct {
}

//query all data from user_info
func (userController *UserController) QueryAllUser(w rest.ResponseWriter, r *rest.Request) {
	var (
		err               error
		baseGorm          *db.BaseGorm
		userInfo          []model.Users
	)
	returnJson := make(map[string]interface{})

	//operate table(query):user_info
	userInfo, err = baseGorm.QueryAllUserInfo()
	if err != nil {
		log.Fatal(err)
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		return
	}

	returnJson["code"] = 0
	returnJson["msg"] = "query * from user_info successfully!"
	returnJson["userInfo"] = userInfo
	w.WriteJson(returnJson)
}


func (userController *UserController) QueryUserById(w rest.ResponseWriter, r *rest.Request) {
	returnJson := make(map[string]interface{})

	//解析获取数据
	id := r.PathParam("id")
	userId, _ := strconv.Atoi(id)

	//通过gorm操作数据库
	baseGorm:=db.BaseGorm{}
	userInfo, err := baseGorm.QueryUserInfoById(userId)
	if err != nil {
		log.Fatal(err)
		//操作失败返回结果
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		return
	}

	//操作成功返回结果
	returnJson["code"] = 0
	returnJson["msg"] = "query userInfo success!"
	returnJson["user"] = userInfo
	w.WriteJson(returnJson)
}

//delete
func (userController *UserController) DeleteUser(w rest.ResponseWriter, r *rest.Request) {
	returnJson := make(map[string]interface{})

	id := r.PathParam("id")
	userId, _ := strconv.Atoi(id)

	//开启事务
	baseGorm:=db.BaseGorm{}
	tx := baseGorm.GetDB().Begin()
	err := db.DeleteUserInfo(tx, userId)
	if err != nil {
		log.Fatal(err)
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		tx.Rollback()
		return
	}
	tx.Commit()

	returnJson["code"] = 0
	returnJson["msg"] = "delete user success!"
	w.WriteJson(returnJson)
}
// insert
func (userController *UserController) InsertUser(w rest.ResponseWriter, r *rest.Request) {
	var (
		user,userInfo    model.Users
		err         error
		baseGorm    *db.BaseGorm
	)
	returnJson := make(map[string]interface{})
	err = r.DecodeJsonPayload(&user)
	if err != nil {
		log.Fatal(err)
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		return
	}

	//开启事务
	tx := baseGorm.GetDB().Begin()
	userInfo.Name = user.Name
    userInfo.Password =user.Password
	userInfo.Status = user.Status


	err = db. InsertUserInfo(tx, userInfo)
	if err != nil {
		log.Fatal(err)
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		tx.Rollback()
		return
	}

	tx.Commit()

	returnJson["code"] = 0
	returnJson["msg"] = "user insert success!"
	w.WriteJson(returnJson)
}

//update
func (userController *UserController) UpdateUser(w rest.ResponseWriter, r *rest.Request) {
	var (
		user,userInfo    model.Users
		err         error
		baseGorm    *db.BaseGorm
	)
	returnJson := make(map[string]interface{})
	err = r.DecodeJsonPayload(&user)
	if err != nil {
		log.Fatal(err)
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		return
	}

	//开启事务
	tx := baseGorm.GetDB().Begin()
	userInfo.Id = user.Id
	userInfo.Name = user.Name

	userInfo.Password = user.Password
	// 查询该用户是否存在
	_, err = baseGorm.QueryUserInfoById(user.Id)
	if err != nil {
		log.Fatal(err)
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		tx.Rollback()
		return
	}


	err = db.UpdateUserInfo(tx, userInfo)
	if err != nil {
		log.Fatal(err)
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		tx.Rollback()
		return
	}

	tx.Commit()

	returnJson["code"] = 0
	returnJson["msg"] = "user update success!"
	w.WriteJson(returnJson)
}

