package api

import (
	"net/http"

	"github.com/xuyawen/go-gin-example/pkg/logging"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/xuyawen/go-gin-example/models"
	"github.com/xuyawen/go-gin-example/pkg/e"
	"github.com/xuyawen/go-gin-example/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		if isExist := models.CheckAuth(username, password); isExist {
			token, err := util.GenetateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_ANTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
