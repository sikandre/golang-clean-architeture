package gorm

import (
	"cleanArch/config"
	"cleanArch/internal/domain/model"
	"cleanArch/pkg/datastore"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func NewDB() datastore.Database {
	DBMS := "mysql"

	mySqlConfig := &mysql.Config{
		User:                 config.Configuration.Database.User,
		Passwd:               config.Configuration.Database.Password,
		Net:                  config.Configuration.Database.Net,
		Addr:                 config.Configuration.Database.Addr,
		DBName:               config.Configuration.Database.DBName,
		AllowNativePasswords: config.Configuration.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.Configuration.Database.Params.ParseTime,
		},
	}

	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.User{})

	//db.Create(&model.User{
	//	Name:      "Bob",
	//	Age:       "10",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//})
	db.LogMode(true)

	return &gormDB{DB: db}
}
