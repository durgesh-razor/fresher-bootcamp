package Service

import (
	"retailer-service/Config"
	"retailer-service/Models"
)

func CreateUser(user *Models.User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
