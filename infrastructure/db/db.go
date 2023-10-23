package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/JoaoRafa19/codepix/domain/model"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	// _ "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(filepath.Join(basepath + "/../../.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func ConectDB(env string) *gorm.DB {
	var dns string
	var db *gorm.DB
	var err error

	if env != "test" {

		dns = os.Getenv("dns")
		dbType := os.Getenv("dbType")
		db, err = gorm.Open(dbType, dns)
	} else {
		dns = os.Getenv("dns")
		db, err = gorm.Open(os.Getenv("dbTypeTest"), dns)
	}

	if err != nil {
		log.Fatalf("Error conecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)

	}


	if os.Getenv("AutoMigrateDB") == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db
}
