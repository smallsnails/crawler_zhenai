package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	conn *sql.DB
)

func init() {
	var err error
	conn, err = sql.Open("mysql", "root:root@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func InsertCityList(url, name string) {
	stmt, err := conn.Prepare("insert into zhenai_citylist (`url`,`name`) values (?,?)")
	if err != nil {
		panic(err)
	}
	//result, err := stmt.Exec("http://www.zhenai.com/zhenghun/yilan1", string("宜兰"))
	result, err := stmt.Exec(url, name)
	if err != nil {
		panic(err)
	}
	//rowsAffected, err := result.RowsAffected()
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	//fmt.Println(rowsAffected)
	fmt.Println(lastInsertId)
}

func InsertCity(url, name string) {
	stmt, err := conn.Prepare("insert into zhenai_city (`url`,`name`) values (?,?)")
	if err != nil {
		panic(err)
	}
	//result, err := stmt.Exec("http://www.zhenai.com/zhenghun/yilan1", string("宜兰"))
	result, err := stmt.Exec(url, name)
	if err != nil {
		panic(err)
	}
	//rowsAffected, err := result.RowsAffected()
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	//fmt.Println(rowsAffected)
	fmt.Println(lastInsertId)
}
