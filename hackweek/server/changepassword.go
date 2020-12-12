package server

import (
	"github.com/kashimashino/hackweek/model"
	"github.com/kashimashino/hackweek/utils/errmsg"
)

type EditPwd struct {
	Username string     `json:"username"`
	OldPwd string		`json:"old_pwd"`
	NewPwd string		`json:"new_pwd"`
}
//编辑密码
func EditPassword(id int, data *EditPwd) int {
	var maps = make(map[string]interface{})
	maps["password"] = model.ScryptPw(data.NewPwd)
	//maps["password"] = model.BcryptPw(data.Password)
	err :=model.GetDB().Model(&model.User{}).Where("id=?", id).Updates(maps).Error
	if  err!= nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}


//检查密码
func CheckPassword(data *EditPwd) (code int) {
	var user  model.User
	model.GetDB().Select("id").Where("username=?", data.Username).First(&user)
	if model.ScryptPw(data.OldPwd) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCSE
}

