package mempool

import "time"

type GoMemPool struct {
	fn    func() interface{}
	stack *Stack
}

func NewGoMemPool(size uint, fn func() interface{}) *GoMemPool {
	stack := newStack()

	for i := 0; i < int(size); i++ {
		n := newNode(fn())
		stack.Push(n)
	}

	return &GoMemPool{
		stack: stack,
		fn:    fn,
	}
}

func (g *GoMemPool) Get() interface{} {
	for {
		node, err := g.stack.Pop()
		if err != nil {
			continue
		}
		return node.Value
	}
}

func (g *GoMemPool) GetByTime(timeout time.Duration) interface{} {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			return nil // 超时返回空值
		default:
			node, err := g.stack.Pop()
			if err != nil {
				continue
			}
			return node.Value
		}
	}
}

func (g *GoMemPool) Put(item interface{}) {
	n := newNode(item)
	g.stack.Push(n)
}

func (g *GoMemPool) Cap() int {
	return g.stack.Cap()
}

func (g *GoMemPool) Destory() {
	for !g.stack.IsEmpty() {
		g.stack.Pop()
	}
}
