package main

import (
	"fmt"

	"github.com/zboyco/gosms"
)

func main() {
	// 创建Sender
	sender := &gosms.QSender{
		AppID:  "1234567890",                       // appid
		AppKey: "12345678901234567890123456789000", // appkey
	}

	// 拉取下发状态 此功能需要联系 qcloud sms helper 开通。
	obj, err := sender.SingleSend(
		"短信签名",        // 短信签名
		86,            // 国家码
		"13800000000", // 号码
		10000,         // 短信正文ID
		"100000",      // 参数1
		"5",           // 参数2
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*obj)
}
