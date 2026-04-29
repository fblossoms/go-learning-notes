package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "root:@tcp(localhost:3306)/employees_demo?parseTime=true&loc=Local" // 不是懒连接。parseTime=true用于转换*time.Time为数据库的时间类型
	mysqlDialector := mysql.Open(dsn)

	db, err = gorm.Open(mysqlDialector, &gorm.Config{
		// gorm.Open -> dialector.Initialize() -> sql.Open("mysql", dsn)
		Logger: logger.Default.LogMode(logger.Info),
		//DryRun: true, // 干运行：所有操作都不应用到数据上
	}) // 参数1为驱动对象，参数2为打开选项
	if err != nil {
		log.Fatal(err, "!!!")
	}
}

type Employee struct { // 若属性字段未写完整会被GORM自动丢弃（不报错），区别于先前的会报错要求填完所有字段
	EmpNo     int `gorm:"primaryKey"`
	BirthDate string
	//BD        string `gorm:"column:birth_date"` // 使用标签自定义指定命名，不建议使用，能不用就不用，也可以用来做注释
	FirstName string
	LastName  string
	Email     string
	Gender    string
	HireDate  string
}

type Emp struct { // 若属性字段未写完整会被GORM自动丢弃（不报错），区别于先前的会报错要求填完所有字段
	EmpNo     int `gorm:"primaryKey"`
	BirthDate string
	FirstName string
	LastName  string
	Email     string
	Gender    byte
	HireDate  string
}

func (Employee) TableName() string { // 自定义表名
	return "employees"
}

type Student struct {
	ID       int        `gorm:"primaryKey"`
	Name     string     `gorm:"size:48;not null"` // 若不指定，默认走longtext。size:48相当于varcher(48)
	Age      int        `gorm:"type:tinyint"`
	Birthday *time.Time // *time.Time
	Gender   int        `gorm:"size:1"` // uint8
}

func (s *Student) String() string {
	return fmt.Sprintf("<Stu id:%d, %s, %d>", s.ID, s.Name, s.Age)
}

func main() {
	//var instance Employee                          // 查询时默认时类型名+s当作表名
	//result := db.Take(&instance)                   // Take：拿一条回来
	//fmt.Println(result.Error, result.Error == nil) // 属性
	//fmt.Printf("%#v\n", instance)

	////if config.Debug == true {
	////	db.Debug()
	////}
	//db = db.Debug()
	////db.Migrator().DropTable(&Student{})   // 删除，不要轻易使用
	//db.Migrator().CreateTable(&Student{}) // 迁移器，CreateTable会获取Student的实例的类型（取名称的小写+s作为表名）

	// 增删改查
	// 增加（插入）
	//s1 := Student{ID: 20234332916, Name: "Feifei", Age: 22} // ID可不写，因为会默认从1往后逐个自增1
	//fmt.Println(s1)
	//t1 := time.Now()
	//t2, _ := time.Parse("2006/01/02", "2010/03/15")
	//s2 := Student{Name: "Zhizhi", Age: 21, Birthday: &t1} // 主键默认增1，有删除也不会删除，依旧增1
	//s3 := Student{Name: "Nannan", Age: 21, Birthday: &t2}
	//result := db.Create(&s1) // 落盘了，ACID D持久化
	//result = db.Create([]*Student{&s2, &s3})
	//if result.Error != nil {
	//	log.Fatal(result.Error)
	//}
	//fmt.Println(s1, "!!!")

	// 查询1
	//var students []*Student
	//db.Find(&students)
	//db.Distinct("name").Find(&students)
	//db.Select("id", "name", "age").Limit(2).Offset(3).Find(&students)
	//db.Where("id=20234332916").Find(&students)                             // 推荐
	//db.Where("name = ?", "Feifei").Find(&students)                         // 推荐
	//db.Where("id = ? and name = ?", 20234332916, "Feifei").Find(&students) // 推荐
	//db.Where(&Student{ID: 20234332916}).Find(&students) // {}里面不写默认全表查询
	//db.Where(map[string]interface{}{"id": 20234332916, "name": "Feifei"}).Find(&students)
	//db.Where(&Student{ID: 20234332916}).Or(&Student{ID: 20234332917}).Find(&students)
	//db.Where("id >= 101").Order("name, id desc").Find(&students)
	//db.Group("id").Find(&students) // 按主键分组，不建议
	//db.Group("gender").Find(&students) // 非分组字段显示没有意义
	//if db.Error != nil {
	//	log.Fatal(db.Error)
	//}
	//fmt.Println(students)

	// 查询2：自由度高
	//type Result struct {
	//	name  string
	//	count int
	//}
	//rows, err := db.Table("students").Select("name", "COUNT(id) AS count").Group("name").Rows()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for rows.Next() {
	//	var r Result
	//	err = rows.Scan(&r.name, &r.count)
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//	fmt.Println(r)
	//}

	// 查询3：自由度高，比查询2简洁一些
	//type Result struct {
	//	Name  string // 注意大写
	//	Count int
	//}
	//var rows []*Result
	//// Find(&rows)可以换成Scan(&rows)，Table("students")可以换成Model(&Student{})
	//db.Select("name", "COUNT(id) AS count").Group("name").Table("students").Find(&rows)
	//if db.Error != nil {
	//	log.Fatal(db.Error)
	//}
	//fmt.Println(rows)
	//for _, v := range rows {
	//	fmt.Println(v)
	//}

	// JOIN连接
	//type Result struct {
	//	EmpNo     int
	//	FirstName string
	//	LastName  string
	//	Salary    int
	//}
	//var rs []*Result
	//db.Select("e.emp_no, e.first_name, e.last_name", "s.salary").Table("employees AS e").Joins("JOIN salaries AS s ON e.emp_no = s.emp_no").Scan(&rs)
	//for _, v := range rs {
	//	fmt.Println(v)
	//}

	// LEFT JOIN 左连接，右连接同理
	//type Result struct {
	//	Emp
	//	Salary int
	//}
	//var rs []*Result
	//db.Select("e.*", "s.salary").Table("employees as e").Joins("LEFT JOIN salaries AS s on e.emp_no = s.emp_no").Scan(&rs)
	//for _, v := range rs {
	//	fmt.Println(v)
	//}

	// 更新
	//db.Model(&Student{}).Where("id=20234332916").Update("age", 33) // 更新必须指定更新哪一个实例，不然默认全表更新，非常危险
	//db.Model(&Student{ID: 20234332916}).Update("age", 23)
	//db.Model(&Student{}).Where("id > 20234332916").Updates(&Student{Age: 18})         // 建议做等值条件
	//db.Model(&Student{}).Where("id > 20234332916").Updates(map[string]any{"age": 20}) // 建议做等值条件
	//var s Student
	//fmt.Println(s)
	//db.Take(&s, 20234332925)
	//fmt.Println(s)
	//if s.ID > 0 {
	//	s.Name = "Jingjing"
	//	s.Age += 18
	//}
	//fmt.Println(s)
	//db.Save(&s)
	//fmt.Println(s)

	// 删除，慎用，一定要加条件
	//db.Where("id = 20234332920").Delete(&Student{})
	db.Delete(&Student{}, 20234332921, 20234332922) // IN (20234332921, 20234332922)

	//var s Student
	//db.Take(&s)
	//fmt.Println(s)
	//s = Student{}            // 清理操作
	//db.Take(&s, 20234332918) // 第二值加上主键就查该主键对应的数据。若前面的&变量未清理，会做AND运算
	//fmt.Println(s)
	//if db.Error != nil {
	//	log.Fatal(db.Error, "!!!")
	//}
}
