package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SQLHandler interface
type SQLHandler interface {
	Connect() *gorm.DB
}

func initDB(host, port, user, pass, name string) *gorm.DB {
	teplate := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(teplate, user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	return db
}

type sqlHandler struct {
	DB *gorm.DB
}

// NewSQLHandler method
func NewSQLHandler() SQLHandler {
	c := newConfig()
	db := initDB(
		c.DB.Host,
		c.DB.Port,
		c.DB.User,
		c.DB.Pass,
		c.DB.Name,
	)
	return &sqlHandler{DB: db}
}

// Connect method
func (db *sqlHandler) Connect() *gorm.DB {
	return db.DB
}
