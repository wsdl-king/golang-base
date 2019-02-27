package main

import "fmt"

//极客时间第六讲 对应的main方法
//我所理解的 interface{} (接口类型) 就像java中的Object
// 注意注意 在Go语言中,interface{}代表空接口,任何类型都是它的实现类型
// 然后 这里的类型断言表达式为任意接口类型.(type)

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	a, ok1 := interface{}(container).(interface{})
	i := interface{}(a)
	fmt.Print(len("我"))
	fmt.Println(a)
	fmt.Print(i)
	_, ok2 := interface{}(container).(map[int]string)
	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported container type: %T\n", container)
		return
	}
	dd := "123"
	//interface{}(value).(Type) 不会检查Type的类型,我这里做个更正完事了
	s := interface{}(dd).(string)
	fmt.Printf("%s\n", s)
	switch interface{}(container).(type) {
	case map[int]string:
		fmt.Print("case format []string")
	}
	//if ok{
	//	fmt.Print(strings)
	//}else {
	//	fmt.Print("error")
	//}

}
func getElement(containerI interface{}) (elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T", containerI)
		return
	}
	return
}
