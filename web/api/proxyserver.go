package api

import (
	"ahmd_tools/config"
	"ahmd_tools/model/apiresp"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func CheckBoxStatus(thridobjectid string) (*apiresp.BoxStatusResp, error) {
	ip := config.AppConf.ProxyServer.Ip
	url := "http://" + ip + ":8899/proxyserver/business"

	msgBody := fmt.Sprintf("{\"type\":\"BoxStatusReq\",\"corpId\":\"%s\",\"ackCheck\":false,\"cipher\":null,\"cipherType\":null,\"ukeyDamage\":false,\"macAddress\":null}", thridobjectid)
	formParam := fmt.Sprintf("msgType=BoxStatusReq&corpId=%s&msgBody=%s", thridobjectid, msgBody)

	fmt.Printf("proxy server request form param: %s\n", formParam)
	payload := strings.NewReader(formParam)
	client := &http.Client{Timeout: 5 * time.Second}
	response, err := client.Post(url, "application/x-www-form-urlencoded", payload)
	if err != nil {
		fmt.Printf("send request to proxy server fail!, err: %s\n", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("read response fail!, err: %s\n", err.Error())
		return nil, err
	}
	fmt.Printf("proxy server  response: %s\n", string(result))

	boxStatusResp := &apiresp.BoxStatusResp{}
	err = json.Unmarshal(result, boxStatusResp)
	if err != nil {
		fmt.Printf("parse response fail!, err: %s\n", err.Error())
		return nil, err
	}
	return boxStatusResp, nil
}
