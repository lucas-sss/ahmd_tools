package db

import (
	"errors"
	"fmt"

	"ahmd_tools/model"

	"gorm.io/gorm"
)

func SearchObjectAmountByName(name string) (int64, error) {
	var count int64

	queryString := fmt.Sprintf("%%%s%%", name)
	result := ahmdDB.Model(&model.ThridObject{}).Where("objectname like ?", queryString).Count(&count)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("not fond data")
			return 0, nil
		}
		fmt.Println("search err ", result.Error)
		return -1, result.Error
	}
	fmt.Printf("find object amount: %d\n", count)
	return count, nil
}

func SearchObjectByName(name string) ([]model.ThridObject, error) {
	var objects []model.ThridObject
	queryString := fmt.Sprintf("%%%s%%", name)
	result := ahmdDB.Where("objectname like ?", queryString).Find(&objects)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("not fond data")
			return objects, nil
		}
		fmt.Println("search err ", result.Error)
		return nil, result.Error
	}
	fmt.Printf("find object amount: %d\n", result.RowsAffected)
	return objects, nil
}
