package Go_workspace

// orm crud 记录
import (
	//"database/sql"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	id       uint
	name     string
	password string
}
type todos struct {
	id      uint
	user_id uint
	title   string
	content string
}

func main() {
	// 连接sql数据库

	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
	}
	//
	db.Migrator().DropTable()

	db.AutoMigrate(&Users{}, &todos{})
	//db.Migrator().CreateTable(&Users{});
	//crud
	user := Users{1, "name", "passwd"}
	db.Create(user)

	fisrtUser := db.First(&user) // 查第一个
	k := 5                       //第k个
	target_user := Users{}
	target_user.id = uint(k)
	db.Find(&target_user)

	//update 和updates 的struct 更新
	db.Model(&target_user).Update("id", 1)                                     // set id->1
	db.Model(&target_user).Updates(Users{id: 1, name: "wyx", password: "psd"}) // 整体update
	//更新选定字段   select()
	db.Model(&user).Select("id", "name").Updates(Users{Name: "new_name", Age: 0})

	db.Delete(&fisrtUser)
	db.Delete(&target_user)
}
