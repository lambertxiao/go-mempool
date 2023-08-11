package mempool

import (
	"reflect"
	"testing"
	"time"
)

func TestGoMemPoolBasic(t *testing.T) {
	memPool := NewGoMemPool(0, func() interface{} {
		return nil
	})

	buf := make([]byte, 1, 4096)
	memPool.Put(buf)
	buf2 := memPool.Get().([]byte)

	if !reflect.DeepEqual(buf, buf2) {
		t.Errorf("Unexpect get item from mem pool")
	}
}

func TestGoMemPoolGet(t *testing.T) {
	type Foo struct{}
	memPool := NewGoMemPool(5, func() interface{} {
		return &Foo{}
	})

	capacity := memPool.Cap()
	if capacity != 5 {
		t.Errorf("The capacity of the memory pool is incorrect")
	}

	for i := 0; i < 5; i++ {
		node := memPool.Get()
		if node == nil {
			t.Fatal("Failed to retrieve a node from the pool")
		}
		_, ok := node.(*Foo)
		if !ok {
			t.Fatal("Failed to retrieve a node from the pool")
		}
	}

	memPool.Destory()

	if !memPool.stack.IsEmpty() {
		t.Errorf("The memory pool has not been emptied")
	}
}

func BenchmarkGet(b *testing.B) {
	memPool := NewGoMemPool(uint(10000), func() interface{} {
		buf := make([]byte, 1024*1024)
		return buf
	})

	for i := 0; i < 10000; i++ {
		memPool.Get()
	}

	b.ReportAllocs()
}

func BenchmarkGetByTime(b *testing.B) {
	memPool := NewGoMemPool(uint(10000), func() interface{} {
		buf := make([]byte, 1024*1024)
		return buf
	})

	for i := 0; i < 10001; i++ {
		memPool.GetByTime(2 * time.Second)
	}

	b.ReportAllocs()
}

func BenchmarkPut(b *testing.B) {
	memPool := NewGoMemPool(uint(0), nil)

	for i := 0; i < 10000; i++ {
		memPool.Put(make([]byte, 1024*1024))
	}

	for i := 0; i < 10000; i++ {
		memPool.Get()
	}

	b.ReportAllocs()
}
