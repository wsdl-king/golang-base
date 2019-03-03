package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

//极客时间33讲,讲述的是临时变量池--通常用来做缓存
// 笔记: 临时变量池的结构, 顶层是一个跟P数量一样的本地池列表,其实就是一个数组,存储好多好多本地池
// 每个本地池的结构: 存储私有临时对象的字段private 代表了临时对象列表的字段shared,以及一个sync.mutex互斥锁
//我put的话会先put到与我当前goroutine的本地池内,如果我当前goroutine的private私有临时变量已经有值了,那么
// 我继续的put就会put到shared共享池中,需要加锁,且放在共享临时对象列表的末尾 这个共享池可以被任何的goroutine访问到.
//GET流程, 我先拿我自己的private对象,拿不到就在本地池列表遍历拿shared共享对象,最后都没有我就宣告GG,老子直接New
// private存在的意义是 有利于临时变量的快速存取, shared存在的意义是:用于临时对象的池内共享
//讲一下我的销毁流程,在GC以后会销毁临时变量池中的值,具体是:
// sync包初始化的时候就会向GO运行时的系统注册一个函数,这个函数的功能就是清除临时变量池中的值(池清理函数)
// 每次Gc的时候都会调用这个函数
// sync包中还有一个包级私有的全局变量,这个变量代表了当前的程序中使用的所有的临时对象池的汇总.称为 池汇总列表
// 在GET-PUT方法使用的时候这个池对象就会被添加到池汇总列表,正因为如此,清理函数可以访问所有真正被使用的临时对象的值
// 清理临时变量池的时候会将 private shared 置位nil,然后再把池中的所有本地池列表销毁掉,最后池汇总列表会被清理函数置位空切片
// 对于sync.pool的话 如果没有外层引用 就会清除哒
func main() {

	//simplePool()
	jkdemo27()
}

//临时对象池的简单应用
// 临时对象池,它的值可以被用来存储临时的对象, pool属于结构类型,值被使用以后,就不应该再被复制
// 满足条件: 不需要持久的使用,可有可无,但是有的话更好,创建销毁可以在任何时候发生,并且不影响程序的功能
// 它们也应该是无需区分的,其中的任何值都可以代替另一个.
// GET拿到值以后不会存储,会返回给调用方
func simplePool() {
	// 初始化的时候就应该给,否则GET可能获得到nil,不是开箱即用的
	var pool = sync.Pool{
		New: func() interface{} {
			return "1"
		},
	}
	// 这里的做法是不可取的,我取值和存值都应该是一种数据类型的
	pool.Put(1)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}

//极客时间对应的练习demo27

// bufPool 代表存放数据块缓冲区的临时对象池。
var bufPool sync.Pool

// Buffer 代表了一个简易的数据块缓冲区的接口。
type Buffer interface {
	// Delimiter 用于获取数据块之间的定界符。
	Delimiter() byte
	// Write 用于写一个数据块。
	Write(contents string) (err error)
	// Read 用于读一个数据块。
	Read() (contents string, err error)
	// Free 用于释放当前的缓冲区。
	Free()
}

// myBuffer 代表了数据块缓冲区一种实现。
//结构体+实现接口我们常用的方式
type myBuffer struct {
	buf       bytes.Buffer
	delimiter byte
}

func (b *myBuffer) Delimiter() byte {
	return b.delimiter
}

func (b *myBuffer) Write(contents string) (err error) {
	if _, err = b.buf.WriteString(contents); err != nil {
		return
	}
	return b.buf.WriteByte(b.delimiter)
}

func (b *myBuffer) Read() (contents string, err error) {
	return b.buf.ReadString(b.delimiter)
}

func (b *myBuffer) Free() {
	bufPool.Put(b)
}

// delimiter 代表预定义的定界符。
var delimiter = byte('\n')

func jkdemo27() {
	buf := GetBuffer()
	defer buf.Free()
	buf.Write("A Pool is a set of temporary objects that" +
		"may be individually saved and retrieved.")
	buf.Write("A Pool is safe for use by multiple goroutines simultaneously.")
	buf.Write("A Pool must not be copied after first use.")

	fmt.Println("The data blocks in buffer:")
	for {
		block, err := buf.Read()
		if err != nil {
			if err == io.EOF {
				//读方法 最后会报一个end of,告诉我们,小子你已经读完了
				//具体的 我还要io-Api学习的时候再详细笔记
				fmt.Println("读完了")
				break
			}
			panic(fmt.Errorf("unexpected error: %s", err))
		}
		fmt.Print(block)
	}
}
func init() {
	bufPool = sync.Pool{
		//初始化定义了NEW函数返回的是指针,定义了间隔符,我们一般在init中进行初始化New函数的调用
		// 如果你不New 你pool.get可能得到的就是nil
		New: func() interface{} {
			return &myBuffer{delimiter: delimiter}
		},
	}
}

// GetBuffer 用于获取一个数据块缓冲区。
func GetBuffer() Buffer {
	return bufPool.Get().(Buffer)
}
