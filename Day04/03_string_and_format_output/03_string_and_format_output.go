package main

import "fmt"

func main() {
	var s = "    " // 虽然能打印，但是不建议，可以使用\t实现
	fmt.Print(s)

	var s2 = `
	""`
	fmt.Print(s2) // 等价于\n，里面可以写""

	var s3 = 'a'
	fmt.Printf("\n%#x", s3) // %x：按十六进制打印，#显示进制

	var f = 150.11434
	fmt.Printf("\n%f", f)
	d := fmt.Sprintf("\n%f", f) // 接收字符串进入变量标识符
	fmt.Print(d)
	d = fmt.Sprintln("\nollama")
	fmt.Printf("%s\n%[1]q", d) // 由于是ln所以会打印出自带的\n换行符
}
