package Service

import (
	"fmt"
	"retailer-service/Config"
	"retailer-service/Models"
	"strconv"
	"time"
)

func checkCoolDown(order *Models.Orders) bool {
	lastOrderTime := order.CreatedAt
	lastOrderTime = lastOrderTime.Add(5 * time.Minute)
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(lastOrderTime)
	return currentTime.Before(lastOrderTime)
}

func PlaceOrder(order *Models.Orders) (err error) {
	mu.Lock()
	var _order Models.Orders
	fmt.Println(order.UserID)

	if err = Config.DB.Where("user_id = ?", order.UserID).Last(&_order).Error; err != nil {
		fmt.Println("hihi")
	}

	if checkCoolDown(&_order) {
		mu.Unlock()
		order.Status = "Failed"
		return nil
	}
	var product Models.Product
	err = GetProductbyID(&product, strconv.FormatUint(uint64(order.ProductID), 10))
	if err != nil {
		mu.Unlock()
		panic(err)
	}
	product.Quantity -= 1
	Config.DB.Save(&product)
	if err != nil {
		mu.Unlock()
		panic(err)
	}
	if err = Config.DB.Set("gorm:auto_preload", true).Create(order).Error; err != nil {
		mu.Unlock()

		return err
	}
	mu.Unlock()
	return nil
}

func GetOrderByID(order *Models.Orders, id string) (err error) {
	if err = Config.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}
