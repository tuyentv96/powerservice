package service

import (
	"gopkg.in/gin-gonic/gin.v1"
	model "power/models"
	"strconv"
)

func GetRankingDevicePowerInMonth(c *gin.Context) {
	ret_val:= model.Response{Status:true,Rcode:200}
	ret_val.Message="Success"

	//token_str:=c.Request.Header.Get("auth")

	var hid string
	var time int64

	time,_=strconv.ParseInt(c.Query("time"),0,64)
	hid=c.Query("hid")
	println("DID:",hid)

	ldevice,err := model.GetRankingDevicePowerInMonth(hid,time)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		c.JSON(200, ret_val)
		return
	}
	print(ldevice)

	data:= model.PowerOutput{}

	data.Hid=hid
	data.Devices=ldevice

	ret_val.Data=data

	c.JSON(200, ret_val)
}