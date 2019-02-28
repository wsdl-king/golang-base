package test

import (
	"fmt"
	"testing"
)

//极客时间23讲 --go 程序测试基础知识
//函数的命名需要_test.go
// 测试分为三种  功能测试 基准测试 示例测试
// 说一句至理名言吧
// 人是否会进步以及进步得有多快,依赖的恰恰就是对自我的否定,这包括否定的深刻与否,
// 以及否定自我的频率如何. 这其实就是"不破不立" 这个词表达的含义
func main() {

	//这里我就写测试函数吧

}

type Cat2 struct {
	name           string // 名字。
	scientificName string // 学名。
	category       string // 动物学基本分类。
}

//错误示例
//测试函数需要以Test为前缀
//这里 这个函数 两个错误 函数的名称不是Test前缀 函数的参数 没有*testing.T
// 当我运行 go test 的时候 会告诉我 检查不到test函数
//go test jk23_test.go
//ok  	command-line-arguments	0.001s [no tests to run]
func Cat_Name() {
	var cat = Cat2{"qws", "qws1", "qws2"}
	fmt.Print(cat.name)
}

//正确示例 这个是 功能测试函数  go test 可以检查到 也可以执行
func TestCat_Name(t *testing.T) {
	cat := Cat2{"qws", "qws1", "qws2"}
	fmt.Print(cat.name)
}

// 正确示例 这个是性能测试 go test可以检查到 也可以执行
// 性能测试的命令
// 解释一下 -4代表4个cpu线程执行.在3秒执行了5000000000次, v1 ns/op ,每次执行耗费v1纳秒
// v2 B/op 每次执行分配了v2字节的内存 . v3 allocs/op 每次执行分配了v3次对象
//qiwenshuai@qws:~/go/src/awesomeProject/golang-base/mypackage$
// go test -run=jk23_test.go -bench=.  -benchtime="3s"
//qwsgoos: linux
//goarch: amd64
//pkg: awesomeProject/golang-base/mypackage
//BenchmarkCat_Name-4   	qwsqwsqwsqwsqws5000000000	         0.00 ns/op  ... B/op  ...allocs/op
//PASS
//ok  	awesomeProject/golang-base/mypackage	0.003s
func BenchmarkCat_Name(b *testing.B) {
	cat := &Cat2{"qws", "qws1", "qws2"}
	fmt.Print(cat.name)
}
