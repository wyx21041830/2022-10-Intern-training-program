package model

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	connectDatabase()
	fmt.Println("connected!!")
	//err := DB.AutoMigrate(&Todo{}) // TODO: add table structs here
	if !DB.Migrator().HasTable(&Todo{}) {
		err := DB.Migrator().CreateTable(&User{})
		if err != nil {
			logrus.Error(err)
		}
	}
	if !DB.Migrator().HasTable(&Todo{}) {
		err := DB.Migrator().CreateTable(&Todo{})
		if err != nil {
			logrus.Error(err)
		}
	}

}

func connectDatabase() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./model")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	loginInfo := viper.GetStringMapString("User")
	dbArgs := loginInfo["username"] + ":" + loginInfo["password"] +
		"@(localhost)/" + loginInfo["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
}
