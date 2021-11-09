package Demo

import (
	"fmt"
	"testing"
)

func Test_Problem1(t *testing.T) {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)
	v1,ok := scoreMap["张三s"]
	fmt.Println(ok,v1)
	//遍历map
	for k,v :=range scoreMap{
		fmt.Println(k,v)
	}
	delete(scoreMap,"张三")
	l := len(scoreMap)
	fmt.Println(l)
}