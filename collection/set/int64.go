package set

// Int64Set is a set of int64
type Int64Set struct {
	Data map[int64]struct{}
}

// Has return true if set contains the key
func (s *Int64Set) Has(key int64) bool {
	_, ok := s.Data[key]
	return ok
}

// Add a key into set
func (s *Int64Set) Add(key int64) {
	s.Data[key] = struct{}{}
}

// Append append keys into set
func (s *Int64Set) Append(keys ...int64) {
	for _, key := range keys {
		s.Data[key] = struct{}{}
	}
}

// Size return the size of set
func (s *Int64Set) Size() int {
	return len(s.Data)
}

// ToSlice return key slice
func (s *Int64Set) ToSlice() []int64 {
	l := make([]int64, 0, len(s.Data))
	for k := range s.Data {
		l = append(l, k)
	}
	return l
}

// NewInt64Set make a int64 set
func NewInt64Set() *Int64Set {
	return &Int64Set{
		Data: map[int64]struct{}{},
	}
}

// NewInt64SetWithValues make a int64 set with values
func NewInt64SetWithValues(keys []int64) *Int64Set {
	set := &Int64Set{
		Data: map[int64]struct{}{},
	}
	for _, key := range keys {
		set.Add(key)
	}
	return set
}

// NewFixedLengthInt64Set make a int64 set with fixed length
func NewFixedLengthInt64Set(length int) *Int64Set {
	return &Int64Set{
		Data: make(map[int64]struct{}, length),
	}
}
