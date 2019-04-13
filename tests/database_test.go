package test

import (
	"fmt"
	"github.com/featherr-engineering/rest-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"testing"
)

func TestDatabase(t *testing.T) {
	cfg := config.GetConfig()

	username := cfg.DBUser
	password := cfg.DBPass
	dbName := cfg.DBName
	dbHost := cfg.DBHost
	dbPort := cfg.DBPort
	dbType := cfg.DBType

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)

	_, err := gorm.Open(dbType, dbUri)
	if err != nil {
		log.Fatalln(err)
	}

}
