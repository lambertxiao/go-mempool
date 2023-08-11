package mempool

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := newStack()
	node1 := newNode(1)
	node2 := newNode(2)
	node3 := newNode(3)

	stack.Push(node1)
	stack.Push(node2)
	stack.Push(node3)

	if stack.head != node3 {
		t.Errorf("Head element is not the last pushed element")
	}
}

func TestStackPop(t *testing.T) {
	stack := newStack()
	node1 := newNode(1)
	node2 := newNode(2)
	node3 := newNode(3)

	stack.Push(node1)
	stack.Push(node2)
	stack.Push(node3)

	if n, err := stack.Pop(); err != nil || n.Value.(int) != 3 {
		t.Fatal("Unexpected error")
	}

	if n, err := stack.Pop(); err != nil || n.Value.(int) != 2 {
		t.Fatal("Unexpected error")
	}

	if n, err := stack.Pop(); err != nil || n.Value.(int) != 1 {
		t.Fatal("Unexpected error")
	}

	_, err := stack.Pop()
	if err != ErrStackEmpty {
		t.Fatal("Unexpected error")
	}
}
