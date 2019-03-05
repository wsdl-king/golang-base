package main

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

//极客时间40节内容,讲的是 io包中的接口和工具(上)
// 这里写了几个基于io.reader和io.writer实现的接口,然后根据不同的接口特性我们可以选用不同的实现
func main() {
	//simpleIO()
	demo82()
}

//简单的demo学习io接口中的API
// 从某处读取数据,就用到了io.reader接口. 写数据到某处,就用到了 io.writer的接口
//如果我要做接口的扩展,那我就需要在接口类型之间进行嵌入
func simpleIO() {
	//3个参数 一个是io.writer,另一个是io.reader,还有一个是src中想要拷贝的长度
	src := strings.NewReader(
		"CopyN copies n bytes (or until an error) from src to dst. " +
			"It returns the number of bytes copied and " +
			"the earliest error encountered while copying.")
	dst := new(strings.Builder)
	written, err := io.CopyN(dst, src, 58)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Written(%d): %q\n", written, dst.String())
	}

}

func simpleIoType() {
	//首先读取,需要有一个N,返回最大N值的可读取量
	//io.LimitedReader{}
}

//基于reader接口的实现
//极客时间demo82
func demo82() {
	comment := "Package io provides basic interfaces to I/O primitives. " +
		"Its primary job is to wrap existing implementations of such primitives, " +
		"such as those in package os, " +
		"into shared public interfaces that abstract the functionality, " +
		"plus some other related primitives."

	// 示例1。
	fmt.Println("New a string reader and name it \"reader1\" ...")
	reader1 := strings.NewReader(comment)
	buf1 := make([]byte, 7)
	n, err := reader1.Read(buf1)
	var offset1, index1 int64
	executeIfNoErr(err, func() {
		fmt.Printf("Read(%d): %q\n", n, buf1[:n])
		offset1 = int64(1)
		//增加写指针的位置
		index1, err = reader1.Seek(offset1, io.SeekCurrent)
	})
	executeIfNoErr(err, func() {
		fmt.Printf("The new index after seeking from current with offset %d: %d\n",
			offset1, index1)
		n, err = reader1.Read(buf1)
	})
	executeIfNoErr(err, func() {
		fmt.Printf("Read(%d): %q\n", n, buf1[:n])
	})
	fmt.Println()

	//// 示例2。
	//reader1.Reset(comment)
	//num1 := int64(7)
	//fmt.Printf("New a limited reader with reader1 and number %d ...\n", num1)
	////限定了 我一次读的buf数量
	//reader2 := io.LimitReader(reader1, 30)
	//buf2 := make([]byte, 10)
	//for i := 0; i < 3; i++ {
	//	n, err = reader2.Read(buf2)
	//	executeIfNoErr(err, func() {
	//		fmt.Printf("Read(%d): %q\n", n, buf2[:n])
	//	})
	//}
	//fmt.Println()

	//
	// 示例3。
	//reader1.Reset(comment)
	//offset2 := int64(0)
	//num2 := int64(81)
	//fmt.Printf("New a section reader with reader1, offset %d and number %d ...\n", offset2, num2)
	//// 只能读出原来数据中的一部分,便宜offset 然后读取num个 返回一个新的buffer.Reader
	//reader3 := io.NewSectionReader(reader1, offset2, num2)
	//buf3 := make([]byte, 20)
	//for i := 0; i < 5; i++ {
	//	// n没有数据会报错,有但不全不会报错
	//	n, err = reader3.Read(buf3)
	//	executeIfNoErr(err, func() {
	//		fmt.Printf("Read(%d): %q\n", n, buf3[:n])
	//	})
	//}
	//fmt.Println()
	//
	//// 示例4。
	//reader1.Reset(comment)
	//writer1 := new(strings.Builder)
	//fmt.Println("New a tee reader with reader1 and writer1 ...")
	//reader4 := io.TeeReader(reader1, writer1)
	//buf4 := make([]byte, 40)
	//for i := 0; i < 8; i++ {
	//	//在reader读取的时候,将w的数据写入r
	//	n, err = reader4.Read(buf4)
	//	executeIfNoErr(err, func() {
	//		fmt.Printf("Read(%d): %q\n", n, buf4[:n])
	//	})
	//}
	//fmt.Println()
	//
	// 示例5。
	//reader5a := strings.NewReader(
	//	"MultiReader returns a Reader that's the logical concatenation of " +
	//		"the provided input readers.")
	//reader5b := strings.NewReader("They're read sequentially.")
	//reader5c := strings.NewReader("Once all inputs have returned EOF, " +
	//	"Read will return EOF.")
	//reader5d := strings.NewReader("If any of the readers return a non-nil, " +
	//	"non-EOF error, Read will return that error.")
	//fmt.Println("New a multi-reader with 4 readers ...")
	////接受若干类型的Reader并返回
	//reader5 := io.MultiReader(reader5a, reader5b, reader5c, reader5d)
	//buf5 := make([]byte, 50)
	//for i := 0; i < 8; i++ {
	//	n, err = reader5.Read(buf5)
	//	executeIfNoErr(err, func() {
	//		fmt.Printf("Read(%d): %q\n", n, buf5[:n])
	//	})
	//}
	//fmt.Println()
	//
	// 示例6。
	fmt.Println("New a synchronous in-memory pipe ...")
	// pipe 是一个代理接口, reader和writer都是从这个代理接口衍生出来的
	pReader, pWriter := io.Pipe()
	_ = interface{}(pReader).(io.ReadCloser)
	_ = interface{}(pWriter).(io.WriteCloser)

	//切片的切片
	comments := [][]byte{
		[]byte("Pipe creates a synchronous in-memory pipe."),
		[]byte("It can be used to connect code expecting an io.Reader "),
		[]byte("with code expecting an io.Writer."),
	}

	// 这里添加这个同步工具纯属为了保证下面示例中的打印语句都能够执行完成。
	// 在实际使用中没有必要这样做。
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, d := range comments {
			time.Sleep(time.Millisecond * 500)
			n, err := pWriter.Write(d)
			if err != nil {
				fmt.Printf("write error: %v\n", err)
				break
			}
			fmt.Printf("Written(%d): %q\n", n, d)
		}
		pWriter.Close()
	}()
	go func() {
		defer wg.Done()
		wBuf := make([]byte, 55)
		for {
			n, err := pReader.Read(wBuf)
			if err != nil {
				fmt.Printf("read error: %v\n", err)
				break
			}
			fmt.Printf("Read(%d): %q\n", n, wBuf[:n])
		}
	}()
	wg.Wait()
}

func executeIfNoErr(err error, f func()) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	f()

}
