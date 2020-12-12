package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kashimashino/hackweek/middleware"
	"github.com/kashimashino/hackweek/model"
	"github.com/kashimashino/hackweek/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context)  {
	var data model.User
	var token string

	_=c.ShouldBindJSON(&data)

	code ,id:= model.CheckLogin(data.Username,data.Password)
	if code == errmsg.SUCCSE {
		token,code= middleware.SetToken(data.Username)
	}

	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"msg" : errmsg.GetErrMsg(code),
		"token" : token,
		"id" : id,
	})
}
