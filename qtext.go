package gosms

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// SingleSend 发送单条短信
func (s *QSender) SingleSend(sign string, countryCode int, mobile string, tplID int, params ...string) (*QSingleResult, error) {
	obj := qSingle{
		Ext:    strconv.Itoa(rand.Int()),
		Params: params,
		Sign:   sign,
		Tel: &qTel{
			Mobile:     mobile,
			Nationcode: strconv.Itoa(countryCode),
		},
		Time:  time.Now().Unix(),
		TplID: tplID,
	}

	strRand := strconv.Itoa(rand.Int())

	sigStr := fmt.Sprintf("appkey=%v&random=%v&time=%v&mobile=%v", s.AppKey, strRand, obj.Time, obj.Tel.Mobile)
	sigBytes := sha256.Sum256([]byte(sigStr))
	obj.Sig = hex.EncodeToString(sigBytes[:])

	jsonBytes, _ := json.Marshal(&obj)

	url := fmt.Sprintf("https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=%v&random=%v", s.AppID, strRand)

	body, err := httpSend(url, jsonBytes)
	if err != nil {
		return nil, err
	}

	result := QSingleResult{}
	json.Unmarshal(body, &result)

	if result.Result != 0 {
		return &result, errors.New(result.ErrMsg)
	}
	return &result, nil
}

// MultiSend 统一国家码群发短信
func (s *QSender) MultiSend(sign string, countryCode int, mobiles []string, tplID int, params ...string) (*QMultiResult, error) {
	telphones := []Telphone{}
	for _, v := range mobiles {
		telphones = append(telphones, Telphone{
			Phone: v,
			CC:    countryCode,
		})
	}
	return s.MultiSendEachCC(sign, telphones, tplID, params...)
}

// MultiSendEachCC 各自国家码群发短信
func (s *QSender) MultiSendEachCC(sign string, telphones []Telphone, tplID int, params ...string) (*QMultiResult, error) {
	obj := qMulti{
		Ext:    strconv.Itoa(rand.Int()),
		Params: params,
		Sign:   sign,
		Time:   time.Now().Unix(),
		TplID:  tplID,
		Tel:    []qTel{},
	}

	strMobile := ""
	for i := range telphones {
		obj.Tel = append(obj.Tel, qTel{
			Mobile:     telphones[i].Phone,
			Nationcode: strconv.Itoa(telphones[i].CC),
		})
		strMobile += telphones[i].Phone + ","
	}
	strMobile = strMobile[:len(strMobile)-1]

	strRand := strconv.Itoa(rand.Int())
	strSig := fmt.Sprintf("appkey=%v&random=%v&time=%v&mobile=%v", s.AppKey, strRand, obj.Time, strMobile)
	sigBytes := sha256.Sum256([]byte(strSig))
	obj.Sig = hex.EncodeToString(sigBytes[:])

	jsonBytes, _ := json.Marshal(&obj)

	url := fmt.Sprintf("https://yun.tim.qq.com/v5/tlssmssvr/sendmultisms2?sdkappid=%v&random=%v", s.AppID, strRand)

	body, err := httpSend(url, jsonBytes)
	if err != nil {
		return nil, err
	}

	result := QMultiResult{}
	json.Unmarshal(body, &result)

	if result.Result != 0 {
		return &result, errors.New(result.ErrMsg)
	}
	return &result, nil
}

// PullSingleStatus  拉取单个号码短信下发状态
func (s *QSender) PullSingleStatus(countryCode int, moblie string, beginTimeStr string, endTimeStr string, max int) (*QPullStatusResult, error) {

	if max > 100 {
		return nil, errors.New("最多拉取100条数据")
	}
	beginTime, err := time.Parse("2006-01-02 15:04:05", beginTimeStr)
	if err != nil {
		return nil, errors.New("beginTimeStr 格式不正确 e.g. \"2006-01-02 15:04:05\"")
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", endTimeStr)
	if err != nil {
		return nil, errors.New("endTimeStr 格式不正确 e.g. \"2006-01-02 15:04:05\"")
	}
	obj := qPullSingleInfo{
		BeginTime:  beginTime.Unix(),
		EndTime:    endTime.Unix(),
		Mobile:     moblie,
		Nationcode: strconv.Itoa(countryCode),
		Max:        max,
		Time:       time.Now().Unix(),
		Type:       0, // Enum{0: 短信下发状态, 1: 短信回复}
	}

	strRand := strconv.Itoa(rand.Int())
	strSig := fmt.Sprintf("appkey=%v&random=%v&time=%v", s.AppKey, strRand, obj.Time)
	sigBytes := sha256.Sum256([]byte(strSig))
	obj.Sig = hex.EncodeToString(sigBytes[:])

	jsonBytes, _ := json.Marshal(&obj)

	url := fmt.Sprintf("https://yun.tim.qq.com/v5/tlssmssvr/pullstatus4mobile?sdkappid=%v&random=%v", s.AppID, strRand)

	body, err := httpSend(url, jsonBytes)
	if err != nil {
		return nil, err
	}

	result := QPullStatusResult{}

	json.Unmarshal(body, &result)

	if result.Result != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}

// PullSingleReply  拉取单个号码短信回复
func (s *QSender) PullSingleReply(countryCode int, moblie string, beginTimeStr string, endTimeStr string, max int) (*QPullReplyResult, error) {

	if max > 100 {
		return nil, errors.New("最多拉取100条数据")
	}
	beginTime, err := time.Parse("2006-01-02 15:04:05", beginTimeStr)
	if err != nil {
		return nil, errors.New("beginTimeStr 格式不正确 e.g. \"2006-01-02 15:04:05\"")
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", endTimeStr)
	if err != nil {
		return nil, errors.New("endTimeStr 格式不正确 e.g. \"2006-01-02 15:04:05\"")
	}
	obj := qPullSingleInfo{
		BeginTime:  beginTime.Unix(),
		EndTime:    endTime.Unix(),
		Mobile:     moblie,
		Nationcode: strconv.Itoa(countryCode),
		Max:        max,
		Time:       time.Now().Unix(),
		Type:       1, // Enum{0: 短信下发状态, 1: 短信回复}
	}

	strRand := strconv.Itoa(rand.Int())
	strSig := fmt.Sprintf("appkey=%v&random=%v&time=%v", s.AppKey, strRand, obj.Time)
	sigBytes := sha256.Sum256([]byte(strSig))
	obj.Sig = hex.EncodeToString(sigBytes[:])

	jsonBytes, _ := json.Marshal(&obj)

	url := fmt.Sprintf("https://yun.tim.qq.com/v5/tlssmssvr/pullstatus4mobile?sdkappid=%v&random=%v", s.AppID, strRand)

	body, err := httpSend(url, jsonBytes)
	if err != nil {
		return nil, err
	}

	result := QPullReplyResult{}

	json.Unmarshal(body, &result)

	if result.Result != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}

// PullStatus  拉取短信下发状态
func (s *QSender) PullStatus(max int) (*QPullStatusResult, error) {

	if max > 100 {
		return nil, errors.New("最多拉取100条数据")
	}
	obj := qPullInfo{
		Max:  max,
		Time: time.Now().Unix(),
		Type: 0, // Enum{0: 短信下发状态, 1: 短信回复}
	}

	strRand := strconv.Itoa(rand.Int())
	strSig := fmt.Sprintf("appkey=%v&random=%v&time=%v", s.AppKey, strRand, obj.Time)
	sigBytes := sha256.Sum256([]byte(strSig))
	obj.Sig = hex.EncodeToString(sigBytes[:])

	jsonBytes, _ := json.Marshal(&obj)

	url := fmt.Sprintf("https://yun.tim.qq.com/v5/tlssmssvr/pullstatus?sdkappid=%v&random=%v", s.AppID, strRand)

	body, err := httpSend(url, jsonBytes)
	if err != nil {
		return nil, err
	}

	result := QPullStatusResult{}

	json.Unmarshal(body, &result)

	if result.Result != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}

// PullReply  拉取短信回复
func (s *QSender) PullReply(max int) (*QPullReplyResult, error) {

	if max > 100 {
		return nil, errors.New("最多拉取100条数据")
	}
	obj := qPullInfo{
		Max:  max,
		Time: time.Now().Unix(),
		Type: 1, // Enum{0: 短信下发状态, 1: 短信回复}
	}

	strRand := strconv.Itoa(rand.Int())
	strSig := fmt.Sprintf("appkey=%v&random=%v&time=%v", s.AppKey, strRand, obj.Time)
	sigBytes := sha256.Sum256([]byte(strSig))
	obj.Sig = hex.EncodeToString(sigBytes[:])

	jsonBytes, _ := json.Marshal(&obj)

	url := fmt.Sprintf("https://yun.tim.qq.com/v5/tlssmssvr/pullstatus?sdkappid=%v&random=%v", s.AppID, strRand)

	body, err := httpSend(url, jsonBytes)
	if err != nil {
		return nil, err
	}

	result := QPullReplyResult{}

	json.Unmarshal(body, &result)

	if result.Result != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
