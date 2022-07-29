package model

import (
	"go_mvc/my-app/src/service"
	"log"
)

func GetAll() (datas []Coffee_recipe) {
	result := Db.Find(&datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOne(id int) (data Coffee_recipe) {
	result := Db.First(&data, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (b *Coffee_recipe) Create() {
	result := Db.Create(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

// func (b *Coffee_recipe) Update() {
// 	result := Db.Save(b)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// }

// func (b *Coffee_recipe) Delete() {
// 	result := Db.Delete(b)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// }

func CreateUser(username string, password string) error {

	passwordEncrypt, _ := service.PasswordEncrypt(password)

	result := Db.Create(&User{Username: username, Password: passwordEncrypt})
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func GetUser(username string) User {
	var user User
	log.Println(username)
	Db.First(&user, "username = ?", username)
	return user
}

func RegisterBeans(country, quantity string) error {
	result := Db.Create(&Beans{Country: country, Quantity: quantity})
	if result != nil {
		log.Print(result.Error)
		return result.Error
	}

	return nil
}

func GetAllData() (datas []Coffee_recipe) {
	result := Db.Find(&datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
