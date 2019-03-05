package main

//极客时间42讲,讲的是bufio
// 实现的io操作内置了缓冲区
// 一个读缓冲区一个写缓冲区,接受的是io.Reader或者是io.Writer的类型,然后作为高速缓存
//记住两个方法的底层应用 fill flush
// write方法,如果写入的字节太多,缓冲区容不下,并且缓冲区已经空了, 我就直接写到底层数据
func main() {

}

//两个初始化Reader的函数
func simple() {

	//bufio.Writer
	// 默认4Kb的缓冲区
	//bufio.NewReader()
	//这个需要我自己手动设置缓冲区的大小
	//bufio.NewReaderSize()
}

func jike() {

}
