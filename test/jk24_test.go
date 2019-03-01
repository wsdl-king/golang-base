package test

import (
	"errors"
	"fmt"
	"math"
	"testing"
	//"time"
)

//极客时间24讲

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}

// introduce 用于生成介绍内容。
func introduce() string {
	return "Welcome to my Golang column."
}

//测试函数  我执行的命令是 go test . (在当前目录下,反馈给我的结果是)
//ok  	awesomeProject/golang-base/test 	0.001s  是否成功/测试路径/执行时间
// go test 是会缓存结果的 以便在将来的构建中重用 可以通过 go env GOCACHE命令来查看缓存目录的路径
// 一旦变动缓存就会失效 --- go clean  -cache 就可以使缓存失效
//func TestHello(t *testing.T) {
//	var name string
//	greeting, err := hello(name)
//	if err == nil {
//		t.Errorf("The error is nil, but it should not be. (name=%q)",
//			name)
//	}
//	if greeting != "" {
//		t.Errorf("Nonempty greeting, but it should not be. (name=%q)",
//			name)
//	}
//	name = "Robert"
//	greeting, err = hello(name)
//	if err != nil {
//		t.Errorf("The error is not nil, but it should be. (name=%q)",
//			name)
//	}
//	if greeting == "" {
//		t.Errorf("Empty greeting, but it should not be. (name=%q)",
//			name)
//	}
//	expected := fmt.Sprintf("Hello, %s!", name)
//	if greeting != expected {
//		t.Errorf("The actual greeting %q is not the expected. (name=%q)",
//			greeting, name)
//	}
//	t.Logf("The expected greeting is %q.\n", expected)
//}

// 功能测试
//func TestIntroduce(t *testing.T) {
//	intro := introduce()
//	expected := "Welcome to my Golang column."
//	if intro != expected {
//		t.Errorf("The actual introduce %q is not the expected.",
//			intro)
//	}
//	t.Logf("The expected introduce is %q.\n", expected)
//}

//功能测试
func TestFail(t *testing.T) {
	//t.Fail()
	//t.FailNow() // 此调用会让当前的测试立即失败。
	t.Log("Failed.")
}

//此函数测试t.Fail方法
//func Test_Fail(t *testing.T) {
//	time.Sleep(time.Second)
//	//打印并且立即生效 FailNow()
//	//t.Fatal("我手动执行t.Fatal")
//	// 单纯的记录日志
//	t.Log("日你啊 t.log")
//	//测试失败看到测试的数据
//	//t.Error("测试失败看到测试的数据")
//	// 执行完毕返回的数据为:
//	//qws--- FAIL: Test_Fail (1.00s)
//	//jk24_test.go:79: 我手动执行t.Fatal
//	//FAIL
//	//FAIL	awesomeProject/golang-base/test	1.002s
//	// 1 失败结果不会缓存,每次运行最新的
//	// 2  第二行结果 指明了失败的代码行数 还有我自己打印的字符串
//	// 3  t.Fatal方法的内部 其实你可以看到  调用的是
//	//  t.log(fmt.Sprintln(args...))
//	//	t.FailNow()
//	// 先用log打印 再马上退出
//}

// 性能测试的命令执行及解读
//  1 种  go run=***_test.go  -bench=. (只有运行它我才会去运行性能测试) -benchtime="3s" (指定我执行的秒数)
//  2 种  go test -bench=.  (只有运行它我才会去运行性能测试) -run=^$(执行名字为空的功能测试函数)
// go test 不加命令会执行所有功能函数

//结果:
//qiwenshuai@qws:~/go/src/awesomeProject/golang-base/test$  go test -bench=. -run=^$
//qwsgoos: linux
//goarch: amd64
//pkg: awesomeProject/golang-base/test
//BenchmarkCat_Name-4    	qwsqwsqwsqwsqws2000000000	         0.00 ns/op
//BenchmarkGetPrimes-4   	  300000	      4801 ns/op
//PASS
//ok  	awesomeProject/golang-base/test	1.500s

//结果解读 -4就是CPU的数量,但是不是计算机实际的CPU,逻辑核心CPU吧,代表着同时运行goroutine的能力
// 可以通过 -cpu=2来指定同时运行cpu的核数
// b.n默认是1 循环运行一次 它运行几次跟你代码实际执行的时间和外面你指定的运行时间有关
// 补充一下 -count=x x代表运行测试函数的次数
func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(1000)
	}
}

// GetPrimes 用于获取小于或等于参数max的所有质数。
// 本函数使用的是爱拉托逊斯筛选法（Sieve Of Eratosthenes）。
func GetPrimes(max int) []int {
	if max <= 1 {
		return []int{}
	}
	marks := make([]bool, max)
	var count int
	squareRoot := int(math.Sqrt(float64(max)))
	for i := 2; i <= squareRoot; i++ {
		if marks[i] == false {
			for j := i * i; j < max; j += i {
				if marks[j] == false {
					marks[j] = true
					count++
				}
			}
		}
	}
	primes := make([]int, 0, max-count)
	for i := 2; i < max; i++ {
		if marks[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}
