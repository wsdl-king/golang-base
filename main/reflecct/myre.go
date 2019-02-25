package reflecct

//反射的练习，取自头条go的简单demo
import (
	"fmt"
	"reflect"
)

type Tee struct {
	A int
	B string
}

func DD() {
	t := Tee{23., "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	fmt.Print(s)
	mm := s.Type()
	fmt.Println(mm)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d :%s %s=%v\n", i, mm.Field(i).Name, f.Type(), f.Interface())
	}
}
