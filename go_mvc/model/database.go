package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {

	err := godotenv.Load("env/.env") //envファイルをロードする
	if err != nil {
		log.Fatal(err)
	}
	var user = "root"
	pass := os.Getenv("PASS") //envファイルから取得
	dbname := os.Getenv("DBNAME")
	var path string = (user + ":" + pass + "@/" + dbname + "?charset=utf8&parseTime=true")
	fmt.Println(path)

	dialector := mysql.Open(path)
	if Db, err = gorm.Open(dialector); err != nil {
		connect(dialector, 100)
	}
	fmt.Println("db connected!!")
	Db.AutoMigrate(&User{}) //tableが作成されてなかったら作成する
	Db.AutoMigrate(&Coffee_recipe{})
	Db.AutoMigrate(&Beans{})

}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
			return
		}
		panic(err.Error())

	}
}
