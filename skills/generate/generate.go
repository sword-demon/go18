package generate

type StringSet struct {
	List []string `json:"list"`
}

type IntSet struct {
	List []int `json:"list"`
}

type Set[T any] struct {
	List []T `json:"list"`
}

func NewSet[T any]() *Set[T] {
	return &Set[T]{}
}

func (s *Set[T]) Add(item T) {
	s.List = append(s.List, item)
}
