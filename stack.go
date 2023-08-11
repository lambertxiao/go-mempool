package mempool

import "errors"

var (
	ErrStackEmpty = errors.New("stack is empty")
)

type Stack struct {
	head *Node
	hasp int32
	lk   Locker
}

func newStack() *Stack {
	return &Stack{
		head: nil,
		lk:   newCasLocker(),
	}
}

type Node struct {
	next  *Node
	Value interface{}
}

func newNode(value interface{}) *Node {
	return &Node{
		Value: value,
	}
}

func (s *Stack) Push(newHead *Node) {
	s.lk.Lock()
	defer s.lk.Unlock()

	newHead.next = s.head
	s.head = newHead
}

func (s *Stack) Pop() (*Node, error) {
	s.lk.Lock()
	defer s.lk.Unlock()

	if s.head == nil {
		s.hasp = 0
		return nil, ErrStackEmpty
	}

	tmpHead := s.head
	s.head = tmpHead.next
	tmpHead.next = nil

	return tmpHead, nil
}

func (s *Stack) IsEmpty() bool {
	s.lk.Lock()
	defer s.lk.Unlock()

	return s.head == nil
}

func (s *Stack) Cap() (cnt int) {
	s.lk.Lock()
	defer s.lk.Unlock()

	tmpHead := s.head
	for tmpHead != nil {
		cnt++
		tmpHead = tmpHead.next
	}

	return cnt
}
