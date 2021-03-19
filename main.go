package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"goWeb/common"
)





func main(){
	DB := common.InitDB()
	defer DB.Close()
	r := gin.Default()
	r=collectRoute(r)
	r.Run(":8080")
}