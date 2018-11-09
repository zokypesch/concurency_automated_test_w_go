package config

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"

func GetConn() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 sslmode=disable user=kiper_portal dbname=db_go password=kiper_portal")

	return db, err
}
