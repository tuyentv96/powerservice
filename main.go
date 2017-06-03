package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	api "power/service"
)


func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/getDevice",api.GetDevicePower)
	r.GET("/getPowerDeviceLimit",api.GetDevicePowerLimit)
	r.GET("/getPowerDeviceByTime",api.GetDevicePowerByTime)

	r.GET("/getDevicePowerInMonth",api.GetDevicePowerInMonth)
	r.GET("/getDevicePowerInYear",api.GetDevicePowerInYear)

	r.GET("/getHomePowerInYear",api.GetHomePowerInYear)
	r.GET("/getHomePowerInMonth",api.GetHomePowerInMonth)

	r.GET("/getRankingPowerInMonth",api.GetRankingDevicePowerInMonth)
	r.GET("/getRankingPowerInYear",api.GetRankingDevicePowerInYear)
	r.GET("/getRankingPowerAll",api.GetRankingDevicePowerAll)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080
}