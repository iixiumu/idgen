package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MyDB *gorm.DB

type MyID struct {
	ID   uint64 `gorm:"Column:id;Type:BIGINT UNSIGNED;NOT NULL;PRIMARY_KEY;AUTO_INCREMENT;" json:"id"`
	MyID uint64 `gorm:"Column:my_id;Type:BIGINT UNSIGNED;NOT NULL" json:"my_id"`
}

func (MyID) TableName() string {
	return "my_id"
}

func init() {
	db, err := gorm.Open("mysql", "xiumu:123@abcd@tcp(127.0.0.1:3306)/xiumu?allowNativePasswords=true&charset=utf8&parseTime=true&loc=Asia%2FChongqing")
	if err != nil {
		panic("mysql connect error: " + err.Error())
	}
	db.DB().Ping()
	db.DB().SetMaxIdleConns(128)
	db.DB().SetMaxOpenConns(1024)
	// db.LogMode(true)
	db.AutoMigrate(&MyID{})

	ins := &MyID{
		ID:   1,
		MyID: 0,
	}

	for _, id := range []int{1, 2, 3, 4} {
		ins.ID = uint64(id)
		db.Save(ins)
	}

	MyDB = db
}
