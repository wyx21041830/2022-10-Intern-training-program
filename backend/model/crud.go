package model

import (
	"github.com/sirupsen/logrus"
)

func Add(todos *Todos) {
	tmp := *todos
	err := DB.Debug().Create(&tmp).Error
	if err != nil {
		logrus.Error(err)
	}
}

func Update(todos *Todos) {
	tmp := *todos
	err := DB.Debug().Updates(&tmp).Error
	if err != nil {
		logrus.Error(err)
	}
}

func Query1(id uint) Todos { //根据 id查找
	tmp := Todos{}
	tmp.ID = id
	err := DB.Debug().Find(&tmp).Error
	if err != nil {
		logrus.Error(err)
	}
	return tmp
}
func Query2(userId uint) Todos { //根据 userid查找
	tmp := Todos{}
	tmp.UserId = userId
	err := DB.Debug().Find(&tmp).Error
	if err != nil {
		logrus.Error(err)
	}
	return tmp
}

func Del(todos *Todos) {
	tmp := *todos
	err := DB.Debug().Delete(&tmp).Error
	if err != nil {
		logrus.Error(err)
	}
}

func Dels(id uint) {
	tmp := Query2(id)
	err := DB.Debug().Delete(&tmp).Error
	if err != nil {
		logrus.Error(err)
	}
}
