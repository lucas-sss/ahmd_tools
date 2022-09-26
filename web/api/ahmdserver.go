package api

import (
	"ahmd_tools/db"
	"ahmd_tools/model/apiresp"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func PushKeyUpdateNotice(thridobjectid string) (*apiresp.DingtalkApiResp, error) {
	//查询token
	dingToken, err := db.QueryDingToken()
	if err != nil {
		fmt.Printf("query ding token fail, errmsg: %s\n", err.Error())
		return nil, err
	}
	//生成uuid
	uid, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("generate dingtalk api request param[uuid] fail, errmsg: %s\n", err.Error())
		return nil, err
	}

	url := "https://oapi.dingtalk.com/push/to_org_emp?suite_access_token=" + dingToken.TokenValue

	msgBody := fmt.Sprintf("{\"uuid\":\"%s\",\"auth_corpid\":\"%s\",\"appid\":\"1289\",\"data\":\"{\"type\":\"keyupdate\"}\"}", uid, thridobjectid)
	fmt.Printf("json body: %s\n", msgBody)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer([]byte(msgBody)))
	if err != nil {
		fmt.Printf("send key update notice to dingtalk fail, errmsg: %s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read dingtalk response fail!, err: %s\n", err.Error())
		return nil, err
	}
	fmt.Printf("dingtalk response: %s\n", string(result))

	dingtalkApiResp := &apiresp.DingtalkApiResp{}
	err = json.Unmarshal(result, dingtalkApiResp)
	if err != nil {
		fmt.Printf("parse dingtalk response fail!, err: %s\n", err.Error())
		return nil, err
	}
	return dingtalkApiResp, nil
}
