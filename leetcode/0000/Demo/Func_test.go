package Demo

import (
	"fmt"
	"testing"
)

//func test(x,y int,s string)(int,string){
//	n := x+y
//	return  n,fmt.Sprintln(s,n)
//}
func test(fn func() int) int{
	return fn()
}
type FormatFunc func(s string,x,y int) string

func format (fn FormatFunc,s string, x,y int) string{
	return fn(s,x,y)
}
func Test_Problem3(t *testing.T) {
	//test(1,2,"a")
	s1 := test(func() int {return 100})
	s2 := format(func(s string, x,y int) string {
		return fmt.Sprintf(s,x,y)
	},"%d, %d",10, 20)
	println(s1,s2)
}

