package Service

import (
	"retailer-service/Config"
	"retailer-service/Models"
)

func Transaction(order *[]Models.Orders) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}
