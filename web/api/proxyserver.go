package api

import (
	"ahmd_tools/model/apiresp"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func CheckBoxStatus(thridobjectid string) (*apiresp.BoxStatusResp, error) {
	url := "http://10.1.202.250:8899/proxyserver/business"

	msgBody := fmt.Sprintf("{\"type\":\"BoxStatusReq\",\"corpId\":\"%s\",\"ackCheck\":false,\"cipher\":null,\"cipherType\":null,\"ukeyDamage\":false,\"macAddress\":null}", thridobjectid)
	formParam := fmt.Sprintf("msgType=BoxStatusReq&corpId=%s&msgBody=%s", thridobjectid, msgBody)

	fmt.Printf("form param: %s\n", formParam)
	payload := strings.NewReader(formParam)
	response, err := http.Post(url, "application/x-www-form-urlencoded", payload)
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

	boxStatusResp := &apiresp.BoxStatusResp{}
	err = json.Unmarshal(result, boxStatusResp)
	if err != nil {
		fmt.Printf("parse response fail!, err: %s\n", err.Error())
		return nil, err
	}
	return boxStatusResp, nil
}
