package Go_workspace
import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type Users struct {
id uint
name string
password string
}
type todos struct {
id uint
user_id uint
title string
content string
}
func ininMySql(){
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=us789mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(todos), &gorm.Config{})
}
func main() {
	ininMySql()
	db.AutoMigrate(&Users{},&todos{})
	user:={1,"wef","sss"}
	db.Create(user)
	target_user:=db.First(&user)
	db.Model(&user).Update("123","456")
	db.Where(id=11).Delete(&User{})
}