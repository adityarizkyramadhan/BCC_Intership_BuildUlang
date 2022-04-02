package config

import (
	"BCC_Intership_BuildUlang/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabases() (*gorm.DB, error) {
	//dnsBCC := "admin:HnVXVx8rF4G3YjS3nKuQrKVS7apg4Vzt@tcp(13.212.140.154:3306)/intern_bcc_7?parseTime=true"
	dnsXAMPP := "root:@tcp(127.0.0.1:3306)/klinikhewanrebuild?parseTime=true"
	db, err := gorm.Open(mysql.Open(dnsXAMPP), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&domain.Clinic{}, &domain.User{})
	if err != nil {
		panic(err)
	}
	return db, err
}
