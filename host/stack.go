package host

import (
	"encoding/json"
	"errors"
	"sync"
)

var ErrEmptyStack = errors.New("empty stack")

type Stack struct {
	lock sync.RWMutex
	s []int
}

func NewStack() *Stack {
	return &Stack{sync.RWMutex{},make([]int, 0)}
}

func (s *Stack) Json() ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return json.Marshal(s.s)
}

func (s *Stack) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.s)
}

func (s *Stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *Stack) PushMany(v []int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v...)
}

func (s *Stack) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = s.s[:0]
}

func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, ErrEmptyStack
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}
