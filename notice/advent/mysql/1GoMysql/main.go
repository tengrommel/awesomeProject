package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	Name  string  `db:"name"`
	Age   int     `db:"age"`
	Money float64 `db:"rmb"`
}

func main() {
	db, _ := sqlx.Open("mysql", "root:teng@tcp(localhost:3306)/china")
	defer db.Close()
	tx, _ := db.Begin()
	ret1, e1 := tx.Exec("insert into person(name, age, sex) values (?,?,?)", "咸鸭蛋", 20, true)
	ret2, e2 := tx.Exec("delete from person where name not like ?", "%蛋")
	ret3, e3 := tx.Exec("update person set name = ? where name=?", "卤蛋", "双黄蛋")

	if e1 != nil || e2 != nil || e3 != nil {
		fmt.Println("事务执行失败, e1/e2/e3=", e1, e2, e3)
		tx.Rollback()
	} else {
		tx.Commit()
		ra1, _ := ret1.RowsAffected()
		ra2, _ := ret2.RowsAffected()
		ra3, _ := ret3.RowsAffected()
		fmt.Println("事务执行成功影响的列:", ra1+ra2+ra3)
	}
}
