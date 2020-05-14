# gosms
使用go 实现常用的sms服务SDK(非官方)，目前只实现腾讯云短信功能

已实现功能:
1. 单发短信
2. 统一国家码群发短信
3. 各自国家码群发短信
4. 拉取单个号码短信下发状态
5. 拉取短信下发状态（需要腾讯云开通）
6. 拉取单个号码短信回复
7. 拉取短信回复（需要腾讯云开通）
8. 发送语音验证码
9. 发送语音通知

* 其中5、7、8、9没有测试条件所以暂未测试，如果又问题请联系我，谢谢！

### 使用方法
方法很简单，直接上代码

#### 1. 单发短信

```golang
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

	// 发送短信
	res, err := sender.SingleSend(
		"短信签名",     // 短信签名，此处应填写审核通过的签名内容，非签名 ID，如果使用默认签名，该字段填 ""
		86,            // 国家号
		"13800000000", // 手机号
		10000,         // 短信正文ID
		"123456",      // 参数1
		"5",           // 参数2，后面可添加多个参数
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
```

#### 2. 统一国家码群发短信

所有号码都是同一国家码时可以使用此方法

```golang
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

	// 统一国家码群发短信
	res, err := sender.MultiSend(
		"短信签名", // 短信签名，此处应填写审核通过的签名内容，非签名 ID，如果使用默认签名，该字段填 ""
		86, // 国家号
		[]string{
			"13800000000", // 手机号
			"13800000000", // 手机号
			"13800000000", // 手机号
		},
		10000,    // 短信正文ID
		"123456", // 参数1
		"5",      // 参数2，后面可添加多个参数
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
```

#### 3. 各自国家码群发短信

群发时号码国家码不同时使用此方法

```golang
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

	// 各自国家码群发短信
	res, err := sender.MultiSendEachCC(
		"短信签名", // 短信签名，此处应填写审核通过的签名内容，非签名 ID，如果使用默认签名，该字段填 ""
		[]Telphone{
			Telphone{
				Phone: "13800000000", // 手机号
				CC:    86,            // 国家号
			},
			Telphone{
				Phone: "13800000000", // 手机号
				CC:    86,            // 国家号
			},
		},
		10000,    // 短信正文ID
		"123456", // 参数1
		"5",      // 参数2，后面可添加多个参数
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
```

#### 4. 拉取单个号码短信下发状态

```golang
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

	// 拉取下发状态
	res, err := sender.PullSingleStatus(
		86,                    // 国家码
		"13800000000",         // 号码
		"2019-04-01 00:00:00", // 开始日期，注意格式
		"2019-04-03 00:00:00", // 结束日期，注意格式
		100,                   // 拉取最大条数，最大拉取100条
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
```

#### 5. 拉取短信下发状态

* 此功能需要联系腾讯云开通

```golang
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
	res, err := sender.PullStatus(100) // 最大拉取100条
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
```

#### 6. 拉取单个号码短信回复

```golang
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

	// 拉取短信回复
	res, err := sender.PullSingleReply(
		86,                    // 国家码
		"13800000000",         // 号码
		"2019-04-01 00:00:00", // 开始日期，注意格式
		"2019-04-03 00:00:00", // 结束日期，注意格式
		100,                   // 拉取最大条数，最大拉取100条
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
```

#### 7. 拉取短信回复

* 此功能需要联系腾讯云开通

```golang
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

	// 拉取短信回复 此功能需要联系 qcloud sms helper 开通。
	res, err := sender.PullReply(100) // 最大拉取100条
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
```

### 8. 发送语音验证码

```golang
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

	// 发送语音验证码
	res, err := sender.VoiceSendCaptcha(
		"13800000000", // 手机号
		2,             // 播报次数,最多3次
		"123456",      // 要播报的验证码，仅支持数字（string类型）
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
```

### 9. 发送语音通知

```golang
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

	// 发送语音通知
	res, err := sender.VoiceSendPrompt(
		"13800000000", // 手机号
		2,             // 播报次数,最多3次
		"语音内容文本", // 要播报的语音文本类容
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
```
