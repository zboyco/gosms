package gosms

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

func httpSend(url string, jsonBytes []byte) ([]byte, error) {
	tr := &http.Transport{ //解决x509: certificate signed by unknown authority
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(url,
		"application/json",
		bytes.NewReader(jsonBytes)) // Content-Type post请求必须设置
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
