package models

import (
	"fmt"
	"github.com/featherr-engineering/rest-api/config"
	"github.com/satori/go.uuid"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {

	cfg := config.GetConfig()

	username := cfg.DBUser
	password := cfg.DBPass
	dbName := cfg.DBName
	dbHost := cfg.DBHost
	dbPort := cfg.DBPort
	dbType := cfg.DBType

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	fmt.Println(dbUri)

	conn, err := gorm.Open(dbType, dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.Set("gorm:table_options", "collation_connection=utf8_general_ci")

	db.Debug().AutoMigrate(&User{}, &Vote{}, &Post{}, &Comment{})
}

func GetDB() *gorm.DB {
	return db
}

type GormModel struct {
	ID        string     `gorm:"primary_key;type:varchar(255);" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-";sql:"index"`
}

func (model *GormModel) BeforeCreate(scope *gorm.Scope) error {
	fmt.Println("Base Before Create")
	u1 := uuid.Must(uuid.NewV4(), nil)
	scope.SetColumn("ID", u1.String())
	return nil
}
