package db

import (
	"ahmd_tools/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func QueryDingToken() (*model.DingDingToken, error) {
	var dingToken model.DingDingToken
	result := ahmdDB.Find(&dingToken)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("not fond ding token")
			return &dingToken, nil
		}
		fmt.Println("search ding token err ", result.Error)
		return nil, result.Error
	}
	fmt.Printf("find ding token: %s\n", dingToken.TokenValue)
	return &dingToken, nil
}
