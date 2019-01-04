package main

import (
	"encoding/xml"
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

type outputXML struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}

func main() {
	router := gin.Default()

	router.POST("wechatreceiver", gettingDataReceiver)

	router.Run()
}

func gettingDataReceiver(c *gin.Context) {
	var receivedData wechatReceiveData
	if err := c.ShouldBindXML(&receivedData); err == nil {
		log.Println(receivedData.FromUserName)
	} else {
		log.Println("Error occurred")
	}

	c.XML(http.StatusOK, outputXML{
		ToUserName:   "tester-suh",
		FromUserName: "from Tester",
		CreateTime:   "2018-01-04",
		MsgType:      "text",
		Content:      "Hello " + receivedData.FromUserName})
}
