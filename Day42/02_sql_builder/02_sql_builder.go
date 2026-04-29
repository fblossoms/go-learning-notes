package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/huandu/go-sqlbuilder"
)

var db2 *sql.DB

func init() {
	var err error
	dsn := "root:@tcp(localhost:3306)/employees_demo"
	db2, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db2.SetConnMaxLifetime(0)
	db2.SetMaxIdleConns(50)
	db2.SetMaxOpenConns(0)
}

type emp2 struct {
	emp_no     int
	first_name string
	last_name  string
}

func main() {
	//query := sqlbuilder.Select("emp_no", "first_name", "last_name").From("employees").Where("emp_no = 10015").String()
	//fmt.Println(query)

	builder := sqlbuilder.Select("emp_no", "first_name", "last_name").From("employees")
	query, args := builder.Where(builder.In("emp_no", 10003, 10010, 10015)).Build()
	fmt.Println(query, args)

	stmt, err := db2.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	//rows, err := stmt.Query()
	rows, err := stmt.Query(args...)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var e emp2
		err = rows.Scan(&e.emp_no, &e.first_name, &e.last_name)
		if err != nil {
			continue
		}
		fmt.Println(e)
	}
}
