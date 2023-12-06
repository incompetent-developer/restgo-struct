package cloudsql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (sql *SQL) Connection() error {
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", sql.Username, sql.Password, sql.Host, sql.DatabaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sql.Client = db
	sql.ClientDB = sqlDB

	return nil
}

func (sql *SQL) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", sql.Username, sql.Password, sql.Host, sql.Port, sql.DatabaseName)
	config := &gorm.Config{}

	if sql.IsModeLogEnabled {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
				LogLevel:                  logger.Info,            // Log level
				IgnoreRecordNotFoundError: false,                  // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,                   // Disable color
			},
		)
		config.Logger = newLogger
	}

	db, err := gorm.Open(mysql.Open(dsn), config)

	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sql.Client = db
	sql.ClientDB = sqlDB

	return nil
}
