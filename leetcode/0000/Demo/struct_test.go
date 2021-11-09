package Demo

import (
	"fmt"
	"testing"
)

type person struct {
	name string
	city string
	age  int8
}

func Test_Problem2(t *testing.T) {
	var p1 person
	p1.name = "pprof.cn"
	p1.city = "北京"
	p1.age = 10
	fmt.Printf("p1=%v\n", p1)  //p1={pprof.cn 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"pprof.cn", city:"北京", age:18}
}
