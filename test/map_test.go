package test

import (
	"fmt"
	"github.com/jeffcail/ginframe/server-common/utils/maps"
	"testing"
	"time"
)

func TestMergeMap(t *testing.T) {
	dest := make(map[interface{}]interface{})
	dest["a"] = "fsdsdfasd"
	dest["b"] = 124

	src := make(map[interface{}]interface{})
	src["a"] = "fsdsdfasd"
	src["b"] = 12
	src["g"] = []int{1, 2, 3}
	src["f"] = []string{"aaa", "bb", "ccc"}

	mergeMap := maps.MergeMap(dest, src)
	fmt.Println(mergeMap)
}

// 测试并发 写 - 读 - 删
func TestConcurrencyRwMap(t *testing.T) {
	// 写
	rmap := maps.NewConcurrencyRwMap(1000)
	for i := 0; i < 300; i++ {
		go func(i int) {
			rmap.Set(fmt.Sprintf("%v", i), i*2)
		}(i)
	}
	time.Sleep(time.Second * 4)
	fmt.Println(len(rmap.Map))
	fmt.Println(rmap.Map)

	// 读
	for i := 0; i < 300; i++ {
		go func(i int) {
			res := rmap.Get(fmt.Sprintf("%v", i))
			fmt.Printf("res:  key:%d => value: %d\n", i, res)
		}(i)
	}
	time.Sleep(time.Second * 4)

	// 删
	for i := 0; i < 30; i++ {
		go func(i int) {
			rmap.Delete(fmt.Sprintf("%v", i))
		}(i)
	}
}
