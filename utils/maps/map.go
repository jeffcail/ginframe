package maps

import "sync"

// MergeMap
// desc: 合并源 map 到目标 map
func MergeMap(dest, src map[interface{}]interface{}) map[interface{}]interface{} {
	out := make(map[interface{}]interface{}, len(dest))
	for k, v := range dest {
		out[k] = v
	}
	for k, v := range src {
		value := v
		if av, ok := out[k]; ok {
			if v, ok := v.(map[interface{}]interface{}); ok {
				if av, ok := av.(map[interface{}]interface{}); ok {
					out[k] = MergeMap(av, v)
				} else {
					out[k] = v
				}
			} else {
				out[k] = value
			}
		} else {
			out[k] = v
		}
	}
	return out
}

type ConcurrencyRwMap struct {
	Map map[interface{}]interface{}
	sync.RWMutex
}

// Go语言原生的map类型并不支持并发读写。
// concurrent-map提供了一种高性能的解决方案:通过对内部map进行分片，降低锁粒度，从而达到最少的锁等待时间(锁冲突)
// 在Go 1.9之前，go语言标准库中并没有实现并发map。
// 在Go 1.9中，引入了sync.Map。新的sync.Map与此concurrent-map有几个关键区别。
// 标准库中的sync.Map是专为append-only场景设计的。
// 因此，如果您想将Map用于一个类似内存数据库，那么使用我们的版本可能会受益。
// 你可以在golang repo上读到更多，这里 and 这里 译注:sync.Map在读多写少性能比较好，否则并发性能很差
// https://github.com/orcaman/concurrent-map
// 三中方案 sync.Map、 concurrent-map 、以及 NewConcurrencyRwMap 自行选择

// NewConcurrencyRwMap 加锁实现map并发安全，缺点锁的粒度大
func NewConcurrencyRwMap(capacity int) *ConcurrencyRwMap {
	if capacity < 0 {
		capacity = 0
	}
	return &ConcurrencyRwMap{
		Map: make(map[interface{}]interface{}, capacity),
	}
}

// Set 写入 或者 更新
func (m *ConcurrencyRwMap) Set(key interface{}, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.Map[key] = value
}

// Get 读
func (m *ConcurrencyRwMap) Get(key interface{}) interface{} {
	m.RLock()
	defer m.RUnlock()
	return m.Map[key]
}

// Delete 删
func (m *ConcurrencyRwMap) Delete(key interface{}) {
	m.Lock()
	defer m.Unlock()
	delete(m.Map, key)
}
