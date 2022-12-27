package set

import (
	"sync"
)

//集合
type Set struct {
	m map[int64]bool
	sync.RWMutex
}

// 创建新的set
func NewSet() *Set {
	return &Set{
		m: map[int64]bool{},
	}
}

// 往set中添加元素
func (s *Set) Add(item int64) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

// 移除set中的元素
func (s *Set) Remove(item int64) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

// 判断set中是否有某个元素
func (s *Set) Has(item int64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// 获取set的长度
func (s *Set) Len() int {
	return len(s.List())
}

// 清空set
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int64]bool{}
}

//判断set是否为空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

//获取set的所有元素
func (s *Set) List() []int64 {
	s.RLock()
	defer s.RUnlock()
	list := []int64{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
