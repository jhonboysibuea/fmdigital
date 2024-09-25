package database

import (
	"crud-app/config"
	"fmt"
	"log"

	"crud-app/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func Connect() *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.Db().User, config.Db().Pass, config.Db().Host, config.Db().Port, config.Db().Schema)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	print(config.Db().EnableMigration)
	// Migrate the database schema
	if config.Db().EnableMigration {
		db.AutoMigrate(&model.User{}) // Migrate all your models
	}
	Database = db
	if err != nil {
		log.Default().Print(err)
		return nil
	} else {
		// logger.DefaultLogger().Info("Successfully connected to the database")
		return db
	}

}

func ConnectH2() *gorm.DB {
	var err error

	// Use SQLite for an in-memory database
	// ":memory:" ensures the database is created in memory (similar to H2)
	dsn := ":memory:"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// Migrate the database schema
	if config.Db().EnableMigration {
		db.AutoMigrate(
			&model.User{},
		) // Migrate all your models
	}

	Database = db
	if err != nil {
		// logger.DefaultLogger().Error(err)
		return nil
	} else {
		// logger.DefaultLogger().Info("Successfully connected to the in-memory SQLite database")
		return db
	}
}
