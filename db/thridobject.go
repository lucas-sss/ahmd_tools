package db

import (
	"errors"
	"fmt"

	"ahmd_tools/model"

	"gorm.io/gorm"
)

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
	fmt.Printf("find object count: %d\n", result.RowsAffected)
	return objects, nil
}
