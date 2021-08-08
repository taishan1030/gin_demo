package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func initDB() (err error) {
	driverName := "mysql"
	dataSourceName := "root:123456@tcp(localhost:3306)/mybatis?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	return nil
}

type account struct {
	Id    int64  `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Money int64  `json:"money" form:"money"`
}

func QuerySingleRow() account {
	sqlStr := "select * from account where id = ?"
	var a account
	if err := db.QueryRow(sqlStr, 1).Scan(&a.Id, &a.Name, &a.Money); err != nil {
		log.Printf("scan failed err : %v\n", err)
	}
	defer db.Close()
	log.Println(a.Id, a.Name, a.Money)
	return a
}

func main() {
	if err := initDB(); err != nil {
		panic(err)
	}
	fmt.Println("connect 2 database success")
	//insertRow()

	r := gin.Default()
	r.GET("user", func(c *gin.Context) {
		account := QuerySingleRow()
		c.JSON(200, gin.H{
			"data": account,
		})
	})
	r.GET("users", func(c *gin.Context) {
		var a account
		c.ShouldBind(&a)
		accounts := QueryMuliRows(a.Name)
		c.JSON(200, gin.H{
			"data": accounts,
		})
	})
	r.Run()
	//QuerySingleRow()
}

func QueryMuliRows(name string) []account {
	sql := "select * from account where name = ?"
	stmt, err := db.Prepare(sql)
	accounts := make([]account, 0)
	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()
	rows, _ := stmt.Query(name)
	for rows.Next() {
		var a account
		if err := rows.Scan(&a.Id, &a.Name, &a.Money); err != nil {
			log.Println(err)
		}
		accounts = append(accounts, a)

	}
	return accounts
}

func updateRow() {
	sql := "update account set name = ? where id = ?"
	res, err := db.Exec(sql, "zhangsan", 1)
	if err != nil {
		fmt.Sprintf("update failed err: %v\n", err)
	}
	if n, err := res.RowsAffected(); err != nil {
		fmt.Sprintf("get RowsAffected err:%v\n", err)
	} else {
		fmt.Printf("update success. rows:%d\n", n)
	}
}

func deleteRow() {
	sql := "delete from account where id = ?"
	res, err := db.Exec(sql, 1)
	if err != nil {
		fmt.Sprintf("delete failed err: %v\n", err)
	}
	if n, err := res.RowsAffected(); err != nil {
		fmt.Sprintf("get RowsAffected err:%v\n", err)
	} else {
		fmt.Printf("delete success. rows:%d\n", n)
	}
}

func insertRow() {
	sql := "insert into account (name,money) values (?, ?)"
	res, err := db.Exec(sql, "zhangsan", 3000)
	if err != nil {
		fmt.Sprintf("insert failed err: %v\n", err)
	}
	defer db.Close()
	if n, err := res.LastInsertId(); err != nil {
		fmt.Sprintf("get LastInsertId err:%v\n", err)
	} else {
		fmt.Printf("insert success. id:%d\n", n)
	}

}
