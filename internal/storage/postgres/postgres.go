package postgres

import (
	"fmt"
	"log"
	"sync"

	"github.com/MyProject273/FlowTasker/internal/config"
	"github.com/MyProject273/FlowTasker/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
}

var dbinstance *Database
var dbonce sync.Once

func GetDatabaseConnection() *Database {
	dbonce.Do(func() {
		dbinstance = &Database{
			Connection: dbConnect(),
		}
	})
	return dbinstance
}

func dbConnect() *gorm.DB {
	dsn := getDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database, error: %v", err)
	}

	err = autoMigrate(db)

	if err != nil {
		log.Fatalf("failed to auto migrate, error: %v", err)
	}

	return db
}

func (d *Database) Close() error {
	sqlDB, err := d.Connection.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Task{},
		&entity.Comment{},
		&entity.Project{},
		&entity.TaskAssign{},
		&entity.User{},
	)
}

func getDsn() string {
	cfg := config.GetConfig().DbCfg
	dsn := "host=" + cfg.Host +
		" user=" + cfg.User +
		" password=" + cfg.Password +
		" dbname=" + cfg.DbName +
		" port=" + cfg.Port +
		" sslmode=" + cfg.SSLMode +
		" timezone=" + cfg.TimeZone
	fmt.Println(dsn)
	return dsn
}
