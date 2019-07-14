package database

import (
	"log"
	// "os"

	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
User ユーザー
	HasMany Posts, UserID(外部キー)
*/
type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);unique_index;not null"`
	Password string `gorm:"type:varchar(60);not null"`
	Posts    []Post
}

// Post 投稿
type Post struct {
	gorm.Model
	URL, Title, Description, Image, Usecase, Genre string
	UserID                                         uint
}

// GormConnect mysqlとの接続
func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "ao6415012"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "clippo"
	// DBMS := os.Getenv("DBMS")
	// USER := "root"
	// PASS := os.Getenv("PASS")
	// PROTOCOL := os.Getenv("PROTOCOL")
	// DBNAME := os.Getenv("DBNAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Println("DBMS = %v", DBMS)
		log.Println("CONNECT = %v", CONNECT)
		log.SetFlags(log.Lshortfile)
		log.Fatalln(err)
	}
	return db
}

// CreateTable テーブル作成
func CreateTable(db *gorm.DB) {
	if db.HasTable("users") {
		db.DropTable("users")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	} else {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	}
	if db.HasTable("posts") {
		db.DropTable("posts")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Post{})
	} else {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Post{})
	}
}
