package main

import "fmt"

func main() {

	// 示例1。
	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	//fmt.Printf("a1: %v (len: %d, cap: %d)\n",
	//	a1, len(a1), cap(a1))
	//s9 := a1[1:4]
	//i:=10
	//s10 :=a1[1:i]
	//fmt.Printf("s10: %v (len: %d, cap: %d)\n",
	//	s10, len(s10), cap(s10))
	////s9[0] = 1
	//fmt.Printf("s9: %v (len: %d, cap: %d)\n",
	//	s9, len(s9), cap(s9))
	//for i := 1; i <= 5; i++ {
	//	s9 := append(s9, i)
	//	fmt.Printf("s9(%d): %v (len: %d, cap: %d)\n",
	//		i, s9, len(s9), cap(s9))
	//}
	//fmt.Printf("a1: %v (len: %d, cap: %d)\n",
	//	a1, len(a1), cap(a1))
	////fmt.Println()
	//fmt.Printf("s9: %v (len: %d, cap: %d)\n",
	//	 s9, len(s9), cap(s9))
	//qiepian(a1)
	//fmt.Print(a1)

	//此切片的例子，a1的长度是7,a1的容量是7
	//我进行a1[:5]以后,可以发现,长度是5，容量是7 --a2
	a2 := append(a1[:6])
	//在次我进行切片的扩容 生成的a3
	a3 := append(a1[:5], 1, 2)
	//可以发现添加的两个元素，那么长度就是7，容量也是7，不过由于添加的两个元素没有超过a1原始的切片长度
	//于是乎你就可以发现，原来的a1也发生了变化，后两个元素进行了替换
	fmt.Println(a1)
	fmt.Println(a3, len(a3), cap(a3))
	//此处a2的切片，因为是和a1共同的底层数组，所以呢，你就可以发现，所有的切片对应统一的底层数组
	fmt.Print(a2, len(a2), cap(a2))
	//这次我们扩容超过 a1的长度7,则你会发现，我产生的新切片对应的是新的数组,我只不过是借用了原来切片对应数组的数据而已
	a9 := append(a1[:5], 1, 2, 3)
	fmt.Println(a9)
	//扩容策略，小于2倍直接2倍扩容，(1024以后，扩容是1.25倍)大于2倍，直接新的容量.
	//如果扩展缩容的话，那我就只能copy一个新的数组了，这样不会影响原来的gc

	//下面是基于数组创建切片
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //数组
	slice1 := arr[0:5]                             //基于数组切片1，左闭右开
	arr[0] = 2
	fmt.Print(slice1)
}
func qiepian(s [7]int) {
	ints := s[1:2]
	fmt.Print(ints)
}
