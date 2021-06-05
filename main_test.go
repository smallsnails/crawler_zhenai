package main

import (
	"crawler_zhenai/zhenai/parser"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func Test(t *testing.T) {
	conn, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	stmt, err := conn.Prepare("insert into zhenai_citylist (`url`,`name`) values (?,?)")
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec("http://www.zhenai.com/zhenghun/yilan1", string("宜兰"))
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

func a() {
	fmt.Println("123")
}

func Test_2(t *testing.T) {
	n := runtime.FuncForPC(reflect.ValueOf(parser.ParseCitylist).Pointer()).Name()
	fmt.Println(n)
}

func Test_3(t *testing.T) {
	s := "crawler_zhenai/zhenai/parser.ParseCitylist"
	b := strings.Contains(s, "Parsecitylist")
	fmt.Println(b)
}

func Test_5(t *testing.T) {

	conn, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	stmt, err := conn.Prepare("insert into go_test_data (`desc`,`time`) values (?,?)")
	if err != nil {
		panic(err)
	}

	format := "2006-01-02 15:04:05"
	timestamp := time.Now().Add(-48 * time.Hour).Unix()
	for i := 0; i < 50; i++ {
		timestamp += int64(i)
		datetime := time.Unix(timestamp, 0).Format(format)
		result, err := stmt.Exec("", datetime)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result.LastInsertId())
	}
}

func Test_6(t *testing.T) {
	conn, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	stmt, err := conn.Prepare("delete from go_test_data where `time` <= ?")
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec("2021-06-04 01:00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}
