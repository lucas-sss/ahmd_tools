package db

import (
	"ahmd_tools/model"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func SearchBindInfo(belongCorp string) (*model.BindInfo, error) {
	var bindInfo model.BindInfo
	result := proxyServerDB.Where("belong_corp = ?", belongCorp).First(&bindInfo)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Printf("not fond bind info for corp: %s\n", belongCorp)
			return &bindInfo, nil
		}
		fmt.Println("search bind info err ", result.Error)
		return nil, result.Error
	}
	data, _ := json.Marshal(&bindInfo)
	fmt.Println("search bind info: ", string(data))
	return &bindInfo, nil
}
