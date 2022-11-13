package model

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORM 允许通过一个现有的数据库连接来初始化 *gorm.DB
var DB *gorm.DB

func Init() {
	connectDatabase()
	err := DB.AutoMigrate(&Users{}, &Todos{}) // TODO: add table structs here
	if err != nil {
		logrus.Fatal(err)
	}
}

func connectDatabase() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./model")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	loginInfo := viper.GetStringMapString("Users")
	dbArgs := loginInfo["username"] + ":" + loginInfo["password"] +
		"@(localhost)/" + loginInfo["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
}
