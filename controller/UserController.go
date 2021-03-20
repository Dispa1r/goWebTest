package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goWeb/common"
	"goWeb/model"
	"goWeb/util"
	"golang.org/x/crypto/bcrypt"
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

func Login(c *gin.Context){
	//获取参数
	//数据验证
	DB := common.GetDB()
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
	var user model.User
	DB.Where("telephone = ?",telephone).First(&user)
	log.Println(user.ID)
	if user.ID ==0{
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"msg":"该用户不存在",
		})
		return
	}
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"msg":"密码错误",
		})
		return
	}
	token,err := common.ReleaseToken(user)
	if err!=nil{
		log.Println("token generate error")
		if err != nil{
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"code":500,
				"message":"系统错误",
			})
			return
		}
	}
	c.JSON(200,gin.H{
		"code":200,
		"message":"登陆成功",
		"data":gin.H{
			"token":token,
		},
	})


}
func Info(c *gin.Context){
	user,_ := c.Get("user")
	c.JSON(200,gin.H{
		"code":200,
		"data":gin.H{
			"user":user,
		},
	})
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
	hashedPassword,err :=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":500,
			"message":"系统错误",
		})
		return
	}
	newUser := model.User{
		Name:name,
		Telephone:telephone,
		Password: string(hashedPassword),
	}
	DB.Create(&newUser)
	log.Print(name,telephone,password)
	c.JSON(200,gin.H{
		"message":"注册成功",
	})
}
