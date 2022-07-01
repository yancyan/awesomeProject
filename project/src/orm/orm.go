package orm

import (
	oracle "github.com/wdrabbit/gorm-oracle"
	"gorm.io/gorm"
	"log"
	. "project/src/config"
	db2 "project/src/db"
)

func TestOrm() {
	InitConfig("dev-f1")

	db, err := gorm.Open(oracle.Open(db2.GetDbUrl()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Insert
	//datas := []TestBean{
	//	{"a", "a", "a", "01", time.Now()},
	//	{"b", "a", "a", "01", time.Now()},
	//}
	//db = db.Create(&datas)
	//db.Commit()

	////Update
	//db = db.Debug().Where("id=?", "a").Model(&bean.TestBean{}).Update("state", "02")
	//
	////Delete
	//db := db.Where("id = ?", "a").Delete(&bean.TestBean{})

	//Select
	var row []TestBean
	db = db.First(&row)
	log.Printf("%+v \n", row)
}

type TestBean struct {
	Id         string `gorm:"column:ID;not null;primaryKey;size:36"`
	CustomerId int    `gorm:"column:CUSTOMER_ID;not null;primaryKey;size:36"`
}

func (s *TestBean) TableName() string {
	return "TEST123"
}
