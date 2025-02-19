package singleton_demo

import (
	"sync"
	"testing"
)

const parCount = 100
func TestSingeltonDemo(t *testing.T){
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins2 != ins1{
		t.Fatal("instance is not equal")
	}
}
func TestParallelSingletonDemo(t *testing.T){
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances :=[parCount]*Singleton{}
	for i := 0; i< parCount;i++{
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i<parCount;i++{
		if instances[i] != instances[i-1]{
			t.Fatal("instance is not equal")
		}
	}
}
