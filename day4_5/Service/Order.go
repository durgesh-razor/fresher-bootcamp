package Service

import (
	"retailer-service/Config"
	"retailer-service/Models"
	"strconv"
)

func PlaceOrder(order *Models.Orders) (err error) {
	mu.Lock()
	var product Models.Product
	err = GetProductbyID(&product, strconv.FormatUint(uint64(order.ProductID), 10))
	if err != nil {
		panic(err)
	}
	product.Quantity -= 1
	Config.DB.Save(&product)
	if err != nil {
		panic(err)
	}
	err = Config.DB.Set("gorm:auto_preload", true).Create(order).Error
	mu.Unlock()
	panic(err)
	if err != nil {
		return err
	}
	return nil
}

func GetOrderByID(order *Models.Orders, id string) (err error) {
	if err = Config.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}
