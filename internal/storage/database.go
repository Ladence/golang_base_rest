package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

const (
	PostgreSQL = "postgres"
	SQLite = "sqlite"
	MySQL = "mysql"
)

func Connect(c *Config) (*gorm.DB, error) {
	return connectLoop(c, 0)
}

func connectLoop(c *Config, count int) (*gorm.DB, error) {
	db, err := attemptConnect(c)
	if err != nil && count < c.MaxConnectAttempts {
		time.Sleep(1 * time.Second)
		return connectLoop(c, count + 1)
	}

	return db, err
}

func attemptConnect(c *Config) (*gorm.DB, error) {
	switch c.Dialect {
	case SQLite:
		return openSqlite(c.Name)
	case PostgreSQL:
		return openPostgreSql(c)
	case MySQL:
		panic("not implemented")
	default:
		return openSqlite("") // on unrecognized dialect, open sqlite in-memory db
	}
}

func openSqlite(dbName string) (*gorm.DB, error) {
	dsn := "file::memory:?cache=shared"
	if dbName != "" {
		dsn = fmt.Sprintf("%s.db", dbName)
	}
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}

func openPostgreSql(c *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", c.Host, c.Username, c.Password, c.Name, c.Port)
	return gorm.Open(postgres.Open(dsn))
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		// todo: here be dragons
		)
}