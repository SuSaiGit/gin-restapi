package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type wechatReceiveData struct {
	ToUserName   string `form:"ToUserName"`
	FromUserName string `form:"FromUserName"`
	CreateTime   string `form:"CreateTime"`
	MsgType      string `form:"MsgType"`
	Content      string `form:"Content"`
	MsgID        int    `form:"MsgId"`
}

func main() {
	router := gin.Default()

	router.POST("wechatreceiver", gettingDataReceiver)

	router.Run()
}

func gettingDataReceiver(c *gin.Context) {
	log.Println("=======gettingDataReceiver started=========")
	var receivedData wechatReceiveData
	if err := c.ShouldBindXML(&receivedData); err == nil {
		log.Println("No error")
		log.Println(receivedData.ToUserName)
	} else {
		log.Println("Error occurred")
		log.Println(err.Error())
	}

	c.XML(http.StatusOK, gin.H{
		"ToUserName":   "tester-suh",
		"FromUserName": "from Tester",
		"CreateTime":   "2018-01-04",
		"MsgType":      "text",
		"Content":      "Recived data" + receivedData.ToUserName})
}
