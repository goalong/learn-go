package api


import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/goalong/learn-go/gin_blog/pkg/err"
	"github.com/goalong/learn-go/gin_blog/models"
	"github.com/goalong/learn-go/gin_blog/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}


func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username:username, Password:password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := err.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, _err := util.GenerateToken(username, password)
			if _err != nil {
				code = err.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = err.SUCCESS
			}
		} else {
			code = err.ERROR_AUTH
		}

	} else {
		for _, _err := range valid.Errors {
			log.Println(_err.Key, _err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": err.GetMsg(code),
		"data": data,
	})
}


