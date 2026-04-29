package main

import (
	"log"
	"os"
)

func main() {
	// 把日志输出到文件
	// os.Open() // 日志写入一般不用只读
	// name：文件路径，flag：定义进行何种操作，读、写、追加、创建等（可以通过或进行组合使用）
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		panic(err)
	}
	f, err := os.OpenFile(
		"logs/first.log",
		os.O_CREATE| // 若文件不存在则创建
			os.O_WRONLY| // 只写
			os.O_APPEND, // 文件末尾追加
		os.ModePerm,
	)
	if err != nil {
		panic(err)
	}
	defer f.Close() // defer为完成后执行，为了关闭文件，释放资源

	// 日志
	log.Println("abc") // 输出有格式（日期），日志一般称为message
	//log.Fatalf("我类似Printf %v", "xyz") // Fatal代表程序结束，下面的语句无法执行
	//log.Panic("我是错误的日志") // 可以使用recover()拦住，使后续代码继续执行

	// 自定义logger（日志记录器）
	// out：可写对象，prefix：消息前缀（一般作为分割符或区分用途），flag：标记
	//l1 := log.New(os.Stdout, "**", log.Ldate|log.Ltime|log.Lmsgprefix)
	//l1 := log.New(os.Stdout, "Info：", log.Ldate|log.Ltime|log.Lmsgprefix)
	l1 := log.New(f, "Info：", log.Ldate|log.Ltime|log.Lmsgprefix)
	l1.Printf("这是我定义的l1：%v", 111)

	// 自定义log属性
	log.Default()                      // 内建的缺省的不对外的std Logger
	log.Default().SetOutput(os.Stdout) // 改为标准输出
	log.Default().SetPrefix("##")      // 加前缀
	log.Default().SetFlags(log.Ldate | log.Ltime)
	log.Println("def")

}
