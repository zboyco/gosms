package qcloudsms

// Sender 发送模型
type Sender struct {
	AppID  string
	AppKey string
}

// Tel 电话
type tel struct {
	Mobile     string `json:"mobile"`     // 手机号码
	Nationcode string `json:"nationcode"` // 国家码
}

/* 短信 Start */

// single 单发短信
type single struct {
	Ext    string   `json:"ext"`
	Extend string   `json:"extend"`
	Params []string `json:"params"`
	Sig    string   `json:"sig"`
	Sign   string   `json:"sign"`
	Tel    *tel     `json:"tel"`
	Time   int64    `json:"time"`
	TplID  int      `json:"tpl_id"`
}

// SingleResult 单发回复
type SingleResult struct {
	Result int    `json:"result"`
	ErrMsg string `json:"errmsg"`
	Ext    string `json:"ext"`
	Fee    int    `json:"fee"`
	Sid    string `json:"sid"`
}

// multi 群发短信
type multi struct {
	Ext    string   `json:"ext"`
	Extend string   `json:"extend"`
	Params []string `json:"params"`
	Sig    string   `json:"sig"`
	Sign   string   `json:"sign"`
	Tel    []tel    `json:"tel"`
	Time   int64    `json:"time"`
	TplID  int      `json:"tpl_id"`
}

// MultiResult 群发回复
type MultiResult struct {
	Result int                 `json:"result"`
	ErrMsg string              `json:"errmsg"`
	Ext    string              `json:"ext"`
	Detail []multiResultDetail `json:"detail"`
	Sid    string              `json:"sid"`
}

// 群发回复详情
type multiResultDetail struct {
	Result     int    `json:"result"`
	ErrMsg     string `json:"errmsg"`
	Fee        int    `json:"fee"`
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
	Sid        string `json:"sid"`
}

// 拉取短信状态请求
type pullInfo struct {
	Max  int    `json:"max"`
	Sig  string `json:"sig"`
	Time int64  `json:"time"`
	Type int    `json:"type"`
}

// 拉取单个手机短信状态
type pullSingleInfo struct {
	BeginTime  int64  `json:"begin_time"`
	EndTime    int64  `json:"end_time"`
	Max        int    `json:"max"`
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
	Sig        string `json:"sig"`
	Time       int64  `json:"time"`
	Type       int    `json:"type"`
}

// PullStatusResult 拉取短信状态结果
type PullStatusResult struct {
	Count  int          `json:"count"`
	Data   []SendStatus `json:"data"`
	ErrMsg string       `json:"errmsg"`
	Result int          `json:"result"`
}

// PullReplyResult 拉取短信状态结果
type PullReplyResult struct {
	Count  int            `json:"count"`
	Data   []ReplyMessage `json:"data"`
	ErrMsg string         `json:"errmsg"`
	Result int            `json:"result"`
}

// SendStatus 短信状态
type SendStatus struct {
	UserReceiveTime string `json:"user_receive_time"`
	Nationcode      string `json:"nationcode"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
	Errmsg          string `json:"errmsg"`
	Description     string `json:"description"`
	Sid             string `json:"sid"`
}

// ReplyMessage 回复消息
type ReplyMessage struct {
	Extend     string `json:"extend"`
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
	Sign       string `json:"sign"`
	Text       string `json:"text"`
	Time       int64  `json:"time"`
}

/* 短信 End */

/* 语音 Start */
type voiceCaptcha struct {
	Ext       string `json:"ext"`
	Msg       string `json:"msg"`
	PlayTimes int    `json:"playtimes"`
	Sig       string `json:"sig"`
	Tel       *tel   `json:"tel"`
	Time      int64  `json:"time"`
}

type voiceResult struct {
	Result int    `json:"result"`
	ErrMsg string `json:"errmsg"`
	CallID string `json:"callid"`
	Ext    string `json:"ext"`
}

/* 语音 End */
