package set

import (
	"slices"
)

type Set[K comparable, V any] struct {
	cache map[K]V
	keys  []K
}

func NewSet[K comparable, V any]() *Set[K, V] {
	return &Set[K, V]{
		cache: make(map[K]V),
		keys:  []K{},
	}
}

func (s *Set[K, V]) Get(k K) (V, bool) {
	v, exists := s.cache[k]
	return v, exists
}

func (s *Set[K, V]) Add(k K, v V) {
	if _, exists := s.cache[k]; !exists {
		s.keys = append(s.keys, k)
	}
	s.cache[k] = v
}

func (s *Set[K, V]) Delete(k K) {
	delete(s.cache, k)

	idx := slices.Index(s.keys, k)

	if idx != -1 {
		s.keys = slices.Delete(s.keys, idx, idx+1)
	}
}

func (s *Set[K, V]) Iterator() func() (*int, *K, V) {
	keys := s.keys

	j := 0

	return func() (_ *int, _ *K, _ V) {
		if j > len(keys)-1 {
			return
		}

		row := keys[j]
		j++

		return &[]int{j - 1}[0], &row, s.cache[row]
	}
}

func (s *Set[K, V]) Size() int {
	return len(s.cache)
}
