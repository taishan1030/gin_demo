package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

type account struct {
	Id    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Money int64  `json:"money" db:"money"`
}

func initDB() (err error) {
	driverName := "mysql"
	dataSourceName := "root:123456@tcp(localhost:3306)/mybatis?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		fmt.Printf("connect database failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return nil
}

func QuerySingleRow() account {
	sqlStr := "select * from account where id = ?"
	var a account
	if err := db.Get(&a, sqlStr, 1); err != nil {
		log.Printf("query failed err : %v\n", err)
	}
	defer db.Close()
	log.Printf("id:%d, name:%s, money:%d", a.Id, a.Name, a.Money)
	return a
}

func main() {
	if err := initDB(); err != nil {
		panic(err)
	}
	fmt.Println("connect database success")
	QuerySingleRow()
	r := gin.Default()
	r.GET("accounts", func(c *gin.Context) {
		accounts := QueryMultiRows()
		c.JSON(200, gin.H{
			"data": accounts,
		})
	})
	r.Run()
}

func QueryMultiRows() []account {
	sql := "select * from account"
	var a []account
	if err := db.Select(&a, sql); err != nil {
		fmt.Printf("query failed err:%v\n", err)
	}
	return a
}

func UpdateRow() {
	sql := "update account set money= ? where id = ?"
	res, err := db.Exec(sql, 1000, 1)
	if err != nil {
		fmt.Printf("update failed err:%v\n", err)
		return
	}

	// 受影响行数
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed err:%v\n", err)
		return
	}
	fmt.Printf("update success affected rows:%d\n", n)

}

func InsertRow() {
	sql := "insert table account (name,money) values (?,?)"
	res, err := db.Exec(sql, "test", 1000)
	if err != nil {
		fmt.Printf("insert failed err:%v\n", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId err:%v\n", err)
		return
	}
	fmt.Printf("insert success. LastInsertId:%d", id)

}

func DeleteRow() {
	sql := "detele from account where id= ?"
	res, err := db.Exec(sql, 1)
	if err != nil {
		fmt.Printf("detele failed err:%v\n", err)
		return
	}

	// 受影响行数
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed err:%v\n", err)
		return
	}
	fmt.Printf("delete success affected rows:%d\n", n)
}

func selectNameQuery() {
	sql := "select * from account where name=:name"
	rows, err := db.NamedQuery(sql, map[string]interface{}{
		"name": "test",
	})
	if err != nil {
		fmt.Printf("name query failed err:%v\n", err)
		return
	}
	defer db.Close()
	for rows.Next() {
		var u account
		if err := rows.StructScan(&u); err != nil {
			fmt.Printf("StructScan failed err:%v\n", err)
			continue
		}
		fmt.Println(u)
	}
}

// batchInsert 批量插入
func batchInsert() {
	users := []account{
		{Name: "111", Money: 10},
		{Name: "222", Money: 10},
		{Name: "333", Money: 10},
	}
	sqlStr := "insert into account (name,money) values (:name,:money)"
	_, err := db.NamedExec(sqlStr, users)
	if err != nil {
		fmt.Printf("batchInsert failed, err:%v\n", err)
		return
	}
	fmt.Println("success")
}
