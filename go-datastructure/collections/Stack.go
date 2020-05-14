package collections

type Stack struct {
	top *node
	len int
}

type node struct {
	prev *node
	val  interface{}
}

func New() *Stack {
	return &Stack{top: nil, len: 0}
}

func (s *Stack) Push(val interface{}) {
	n := &node{prev: s.top, val: val}
	s.top = n
	s.len++
}

func (s *Stack) Pop() interface{} {
	if s.len == 0 {
		return nil
	} else {
		n := s.top
		s.top = s.top.prev
		s.len--
		return n.val
	}
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func (s *Stack) Peek() interface{} {
	if s.len == 0 {
		return nil
	} else {
		return s.top.val
	}
}

func (s *Stack) Length() int {
	return s.len
}
