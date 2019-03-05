package main

import (
	"io"
	"strings"
)

//极客时间41章练习,io包中的接口和工具
// 笔记: 根据定义的方法数量以及是否有接口的嵌入,把io包中的接口分为了简单接口和扩展接口.
// 同时  根据简单接口的扩展接口和实现类型的数量级,分为了核心接口和非核心接口
// io包中的核心接口 io.Reader io.Writer io.Closer 基于这些接口的实现类型都在200个以上.
// 简单接口有4大类  读 写 关闭 seek
// 简单接口有11个 读5写4关1设1
func main() {

	//Reader接口的实现
	//io.ByteReader()
	//这个其实就是可以回退的
	//io.ByteScanner()
	//读取指定位置的值,不去修改Reader
	//io.ReaderAt()
	// 这里是我的值读出来 写到参数里
	//io.WriterTo()
	// 这里是我参数的值读出来,写到我自己
	//io.ReaderFrom()
	// 一个同步内存管道核心实现的io.pipe类型
	//io.Pipe()
}

//demo83 极客时间
func write() {

	comment := "Because these interfaces and primitives wrap lower-level operations with various implementations, " +
		"unless otherwise informed clients should not assume they are safe for parallel execution."
	basicReader := strings.NewReader(comment)
	basicWriter := new(strings.Builder)

	// 示例1。
	reader1 := io.LimitReader(basicReader, 98)
	_ = interface{}(reader1).(io.Reader)

	// 示例2。
	reader2 := io.NewSectionReader(basicReader, 98, 89)
	_ = interface{}(reader2).(io.Reader)
	_ = interface{}(reader2).(io.ReaderAt)
	_ = interface{}(reader2).(io.Seeker)

	// 示例3。
	reader3 := io.TeeReader(basicReader, basicWriter)
	_ = interface{}(reader3).(io.Reader)

	// 示例4。
	reader4 := io.MultiReader(reader1)
	_ = interface{}(reader4).(io.Reader)

	// 示例5。
	writer1 := io.MultiWriter(basicWriter)
	_ = interface{}(writer1).(io.Writer)

	// 示例6。
	pReader, pWriter := io.Pipe()
	_ = interface{}(pReader).(io.Reader)
	_ = interface{}(pReader).(io.Closer)
	_ = interface{}(pWriter).(io.Writer)
	_ = interface{}(pWriter).(io.Closer)
}
