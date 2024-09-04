package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(DBUser string, DBPass string, DBHost string, DBPort string, DBName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPass, DBHost, DBPort, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
