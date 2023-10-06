package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Book     []Book `gorm:"foreignKey:Penulis;reference:ID"`
	Blog     []Blog `gorm:"foreignKey:IDUser;reference:ID"`
}

type Blog struct {
	ID        string
	IDUser    uint64 `json:"iduser" form:"iduser"`
	JudulBlog string `json:"judulblog" form:"judulblog"`
	Konten    string `json:"konten" form:"konten"`
}

type Book struct {
	ID       string
	Judul    string `json:"judul" form:"judul"`
	Penulis  uint64 `json:penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
}

var DB *gorm.DB

func InitModel() {
	var dsn = "root:@tcp(localhost:3306)/km_echo?parseTime=true"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Model : cannot connect to database, ", err.Error())
	}

}

func Migrate() {
	DB.AutoMigrate(&User{}, &Blog{}, &Book{})
}
