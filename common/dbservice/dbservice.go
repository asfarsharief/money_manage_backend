package dbservice

import (
	"fmt"
	"sync"

	log "github.com/asfarsharief/money_management_backend/common/logingservice"
	"github.com/asfarsharief/money_management_backend/lib/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB
var dberr error
var once sync.Once

func GetDbInstance(cfg *config.DatabaseConfiguration) (*gorm.DB, error) {
	once.Do(func() {
		Db, dberr = newDbConnection(cfg)
	})
	err1 := Db.DB().Ping()
	if err1 != nil {
		Db, dberr = newDbConnection(cfg)
	}
	if dberr != nil {
		return nil, dberr
	}
	return Db, nil
}

func newDbConnection(cfg *config.DatabaseConfiguration) (*gorm.DB, error) {
	var connectionString string
	connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, "public", "disable")

	Db, Err := gorm.Open("postgres", connectionString)

	if Err != nil {
		log.Error("Failed to open Database Connection", Err)
		return nil, Err
	}
	log.Info("Db connection created...")
	return Db, nil

}

// CloseDbConnection : close the database connection
func CloseDbConnection(Db *gorm.DB) {
	Db.Close()
	log.Info("Db connection closed...")
}
