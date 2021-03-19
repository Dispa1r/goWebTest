package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goWeb/common"
	"goWeb/model"
	"goWeb/util"
	"log"
	"net/http"
)

func isTelephontExist(db *gorm.DB,telephone string) bool{
	var user model.User
	db.Where("telephone = ?",telephone).First(&user)
	log.Println(user.ID)
	if user.ID !=0{
		return true
	}

	return false
}

func Register(c *gin.Context){
	DB := common.GetDB()
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	if len(telephone)!=11{
		log.Print("mobile error\n")
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"手机号必须为11",
		})
		return
	}
	if len(password)<6{
		log.Print("password error\n")
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"密码不能小于6",
		})
		return
	}
	if len(name)==0{
		name = util.RandomString(10)
	}
	if isTelephontExist(DB,telephone){
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"message":"该用户已存在",
		})
		return
	}
	newUser := model.User{
		Name:name,
		Telephone:telephone,
		Password: password,
	}
	DB.Create(&newUser)
	log.Print(name,telephone,password)
	c.JSON(200,gin.H{
		"message":"注册成功",
	})
}
