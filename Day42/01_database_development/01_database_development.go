package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql 注册了，匿名导入，相当于初始化init()
)

var db *sql.DB

func init() {
	var err error
	dsn := "root:@tcp(localhost:3306)/employees_demo"
	// 格式：[username（用户名）[:password（密码）]@][protocol（协议）[(address（IP地址，若端口不是默认端口需要写端口）)]]/dbname（数据库名称）[?param1=value1&...&paramN=valueN]
	// 懒连接（并未真正连接，用到时才真正连接）
	db, err = sql.Open("mysql", dsn) // db类型是*sql.DB，是一个操作数据库的句柄，底层是一个多协程安全的连接池
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(0) // 设置连接最大存活时间，小于等于0为永不关闭
	db.SetMaxIdleConns(50)   // 最大闲置连接数
	db.SetMaxOpenConns(50)   // 最大打开连接数，小于等于0为无限

	//fmt.Println(db)
	//err = db.Ping() // 认证（username/pwd）authenticate（认证），不是authorization（授权）
	//if err != nil {
	//	log.Fatal(err, "###")
	//}
}

type emp struct {
	emp_no                                                      int
	birth_date, first_name, last_name, email, gender, hire_date string
}

func main() {
	s := "10015 or 1=1" // 模拟客户端传回来的值，1=1为恒成立，全表遍历（拖库）
	//query := "SELECT * FROM employees WHERE emp_no > " + s // 不写=会得到多条，生产中必须明确指出使用哪些字段，禁止写*
	query := "SELECT * FROM employees WHERE emp_no > ?" // 预防注入攻击
	stmt, err := db.Prepare(query)                      // 预编译
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(s)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var e emp
		err = rows.Scan(&e.emp_no, &e.birth_date, &e.first_name, &e.last_name, &e.email, &e.gender, &e.hire_date)
		if err != nil {
			continue
		}
		fmt.Println(e)
	}

	//rows, err := db.Query(query, s) // 有几个?就加多少个变量
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(rows)
	//for rows.Next() { // 遍历
	//	// 对当前行
	//	var e emp
	//	err = rows.Scan(&e.emp_no, &e.birth_date, &e.first_name, &e.last_name, &e.email, &e.gender, &e.hire_date)
	//	if err != nil {
	//		continue
	//	}
	//	fmt.Println(e)
	//}

	//row := db.QueryRow(query) // 只能取到一条
	//if row.Err() != nil {
	//	log.Fatal(row.Err().Error())
	//}
	//var emp_no int
	//var birth_date, first_name, last_name, email, gender, hire_date string

	//var e emp
	//err := row.Scan(&e.emp_no, &e.birth_date, &e.first_name, &e.last_name, &e.email, &e.gender, &e.hire_date) // 行，Scan，提取字段值
	//if err != nil {
	//	log.Fatal(err, "###")
	//}
	//fmt.Println(e.emp_no, e.birth_date, e.first_name, e.last_name, e.email, e.gender, e.hire_date)
	//fmt.Printf("%T %[1]v\n", e.birth_date)
	//fmt.Println(time.Parse("2006-01-02", e.birth_date)) // 操作时间函数
}
