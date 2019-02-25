package main

import "fmt"

//极客时间第九章练习
//对于hashMap这样的结构来说,它的键不能为 函数类型,切片类型,字典类型
//对于hashMap这样的结构来讲,我们定位到具体的桶以后,还需要我们进行key的比对(reason hash impact)因为哈希碰撞
//所以,我还需要拿key进行== !=这样的操作,但是呢 函数类型,切片类型,字典类型是不可以进行比较的
func main() {

	//这种声明方式出来的map是nil,在nil的map上试图添加键值对的时候,会引发错误！初次之外其他的操作不会触发panic
	var badMap2 map[int]int
	//这种声明方式出来的map不是nil,可以对map进行随意的操作了,类似java 中的 new
	var dd = make(map[int]int)
	//或者说你声明的时候 使用这种操作 有一个代码块就是空值
	//但是你不能 var qq map[ints]ints{} 这样的声明，这样的弱智错误不要犯了 啊哈哈 var name Type 不可以加{}
	var qq = map[int]int{}
	qq[0] = 0
	fmt.Println(qq)
	dd[0] = 9
	fmt.Println(dd)
	myInt, ok := interface{}(badMap2).(map[int]int)
	if ok {
		fmt.Println(myInt)
	}
	fmt.Print(badMap2[0])
	fmt.Println(len(badMap2))
}
