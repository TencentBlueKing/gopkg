package set

import "strings"

// StringSet is a set of string
type StringSet struct {
	Data map[string]struct{}
}

// Has return true if set contains the key
func (s *StringSet) Has(key string) bool {
	_, ok := s.Data[key]
	return ok
}

// Add a key into set
func (s *StringSet) Add(key string) {
	s.Data[key] = struct{}{}
}

// Append append keys into set
func (s *StringSet) Append(keys ...string) {
	for _, key := range keys {
		s.Data[key] = struct{}{}
	}
}

// Size return the size of set
func (s *StringSet) Size() int {
	return len(s.Data)
}

// ToSlice return key slice
func (s *StringSet) ToSlice() []string {
	l := make([]string, 0, len(s.Data))
	for k := range s.Data {
		l = append(l, k)
	}
	return l
}

// ToString join the string with sep
func (s *StringSet) ToString(sep string) string {
	l := s.ToSlice()
	return strings.Join(l, sep)
}

// Diff will return the difference of two set
func (s *StringSet) Diff(b *StringSet) *StringSet {
	diffSet := NewStringSet()

	for k := range s.Data {
		if !b.Has(k) {
			diffSet.Add(k)
		}
	}
	return diffSet
}


// NewStringSet make a string set
func NewStringSet() *StringSet {
	return &StringSet{
		Data: map[string]struct{}{},
	}
}

// NewStringSetWithValues make a string set with values
func NewStringSetWithValues(keys []string) *StringSet {
	set := &StringSet{
		Data: map[string]struct{}{},
	}
	for _, key := range keys {
		set.Add(key)
	}
	return set
}

// NewFixedLengthStringSet make a string set with fixed length
func NewFixedLengthStringSet(length int) *StringSet {
	return &StringSet{
		Data: make(map[string]struct{}, length),
	}
}

// SplitStringToSet make a string set by split a string into parts
func SplitStringToSet(s string, sep string) *StringSet {
	if s == "" {
		return &StringSet{Data: map[string]struct{}{}}
	}

	data := map[string]struct{}{}
	keys := strings.Split(s, sep)
	for _, key := range keys {
		data[key] = struct{}{}
	}
	return &StringSet{Data: data}
}
