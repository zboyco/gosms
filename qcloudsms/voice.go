package qcloudsms

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

// VoiceSendCaptcha 发送语音验证码
// 给国内用户发语音验证码（仅支持数字）
func (s *Sender) VoiceSendCaptcha(mobile string, playTimes int, code string) (*VoiceResult, error) {
	obj := voiceCaptcha{
		Ext:       strconv.Itoa(rand.Int()),
		Msg:       code,
		PlayTimes: playTimes,
		Tel: &tel{
			Mobile:     mobile,
			Nationcode: "86",
		},
		Time: time.Now().Unix(),
	}

	strRand := strconv.Itoa(rand.Int())

	sigStr := fmt.Sprintf("appkey=%v&random=%v&time=%v&mobile=%v", s.AppKey, strRand, obj.Time, obj.Tel.Mobile)
	sigBytes := sha256.Sum256([]byte(sigStr))
	obj.Sig = hex.EncodeToString(sigBytes[:])

	jsonBytes, _ := json.Marshal(&obj)

	url := fmt.Sprintf("https://cloud.tim.qq.com/v5/tlsvoicesvr/sendcvoice?sdkappid=%v&random=%v", s.AppID, strRand)

	body, err := httpSend(url, jsonBytes)
	if err != nil {
		return nil, err
	}

	result := VoiceResult{}
	json.Unmarshal(body, &result)

	if result.Result != 0 {
		return &result, errors.New(result.ErrMsg)
	}
	return &result, nil
}

// VoiceSendPrompt 发送语音通知
// 给国内用户发语音通知（支持中文、英文字母、数字及组合，内容长度不超过100字）
func (s *Sender) VoiceSendPrompt(mobile string, playTimes int, promptfile string) (*VoiceResult, error) {
	obj := voicePrompt{
		Ext:        strconv.Itoa(rand.Int()),
		Promptfile: promptfile,
		PlayTimes:  playTimes,
		Prompttype: 2,
		Tel: &tel{
			Mobile:     mobile,
			Nationcode: "86",
		},
		Time: time.Now().Unix(),
	}

	strRand := strconv.Itoa(rand.Int())

	sigStr := fmt.Sprintf("appkey=%v&random=%v&time=%v&mobile=%v", s.AppKey, strRand, obj.Time, obj.Tel.Mobile)
	sigBytes := sha256.Sum256([]byte(sigStr))
	obj.Sig = hex.EncodeToString(sigBytes[:])

	jsonBytes, _ := json.Marshal(&obj)

	url := fmt.Sprintf("https://cloud.tim.qq.com/v5/tlsvoicesvr/sendvoiceprompt?sdkappid=%v&random=%v", s.AppID, strRand)

	body, err := httpSend(url, jsonBytes)
	if err != nil {
		return nil, err
	}

	result := VoiceResult{}
	json.Unmarshal(body, &result)

	if result.Result != 0 {
		return &result, errors.New(result.ErrMsg)
	}
	return &result, nil
}
