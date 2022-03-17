package database

import (
	"HostSec/config"
	"HostSec/operations/log"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB
var RecoveryDB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", config.HipsDBName)
	if err != nil {
		panic(err)
	}
	RecoveryDB, err = gorm.Open("sqlite3", "recovery.db")
	if err != nil {
		panic(err)
	}
	RecoveryDB.AutoMigrate(&log.RecordData{})
}
