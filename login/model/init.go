package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDb() (err error) {
	driverName := "mysql"
	dataSourceName := "root:123456@tcp(localhost:3306)/login?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		fmt.Printf("connect database failed, err:%v\n", err)
		return
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	return nil
}
