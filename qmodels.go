package gosms

// QSender 发送模型
type QSender struct {
	AppID  string
	AppKey string
}

// Tel 电话
type qTel struct {
	Mobile     string `json:"mobile"`     // 手机号码
	Nationcode string `json:"nationcode"` // 国家码
}

/* 短信 Start */

// single 单发短信
type qSingle struct {
	Ext    string   `json:"ext"`
	Extend string   `json:"extend"`
	Params []string `json:"params"`
	Sig    string   `json:"sig"`
	Sign   string   `json:"sign"`
	Tel    *qTel    `json:"tel"`
	Time   int64    `json:"time"`
	TplID  int      `json:"tpl_id"`
}

// QSingleResult 单发回复
type QSingleResult struct {
	Result int    `json:"result"`
	ErrMsg string `json:"errmsg"`
	Ext    string `json:"ext"`
	Fee    int    `json:"fee"`
	Sid    string `json:"sid"`
}

// multi 群发短信
type qMulti struct {
	Ext    string   `json:"ext"`
	Extend string   `json:"extend"`
	Params []string `json:"params"`
	Sig    string   `json:"sig"`
	Sign   string   `json:"sign"`
	Tel    []qTel   `json:"tel"`
	Time   int64    `json:"time"`
	TplID  int      `json:"tpl_id"`
}

// QMultiResult 群发回复
type QMultiResult struct {
	Result int                  `json:"result"`
	ErrMsg string               `json:"errmsg"`
	Ext    string               `json:"ext"`
	Detail []qMultiResultDetail `json:"detail"`
	Sid    string               `json:"sid"`
}

// 群发回复详情
type qMultiResultDetail struct {
	Result     int    `json:"result"`
	ErrMsg     string `json:"errmsg"`
	Fee        int    `json:"fee"`
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
	Sid        string `json:"sid"`
}

// 拉取短信状态请求
type qPullInfo struct {
	Max  int    `json:"max"`
	Sig  string `json:"sig"`
	Time int64  `json:"time"`
	Type int    `json:"type"`
}

// 拉取单个手机短信状态
type qPullSingleInfo struct {
	BeginTime  int64  `json:"begin_time"`
	EndTime    int64  `json:"end_time"`
	Max        int    `json:"max"`
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
	Sig        string `json:"sig"`
	Time       int64  `json:"time"`
	Type       int    `json:"type"`
}

// QPullStatusResult 拉取短信状态结果
type QPullStatusResult struct {
	Count  int           `json:"count"`
	Data   []QSendStatus `json:"data"`
	ErrMsg string        `json:"errmsg"`
	Result int           `json:"result"`
}

// QPullReplyResult 拉取短信状态结果
type QPullReplyResult struct {
	Count  int             `json:"count"`
	Data   []QReplyMessage `json:"data"`
	ErrMsg string          `json:"errmsg"`
	Result int             `json:"result"`
}

// QSendStatus 短信状态
type QSendStatus struct {
	UserReceiveTime string `json:"user_receive_time"`
	Nationcode      string `json:"nationcode"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
	Errmsg          string `json:"errmsg"`
	Description     string `json:"description"`
	Sid             string `json:"sid"`
}

// QReplyMessage 回复消息
type QReplyMessage struct {
	Extend     string `json:"extend"`
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
	Sign       string `json:"sign"`
	Text       string `json:"text"`
	Time       int64  `json:"time"`
}

/* 短信 End */

/* 语音 Start */

// 语音验证码
type qVoiceCaptcha struct {
	Ext       string `json:"ext"`
	Msg       string `json:"msg"`
	PlayTimes int    `json:"playtimes"`
	Sig       string `json:"sig"`
	Tel       *qTel  `json:"tel"`
	Time      int64  `json:"time"`
}

// 语音通知
type qVoicePrompt struct {
	Ext        string `json:"ext"`
	Promptfile string `json:"promptfile"`
	Prompttype int    `json:"prompttype"`
	PlayTimes  int    `json:"playtimes"`
	Sig        string `json:"sig"`
	Tel        *qTel  `json:"tel"`
	Time       int64  `json:"time"`
}

// QVoiceResult 语音验证码和通知返回
type QVoiceResult struct {
	Result int    `json:"result"`
	ErrMsg string `json:"errmsg"`
	CallID string `json:"callid"`
	Ext    string `json:"ext"`
}

/* 语音 End */
