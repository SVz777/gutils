package set

import "github.com/SVz777/gutils/collections"

type Set map[string]collections.Empty

func NewSet(keys ...string) Set {
	s := make(Set, len(keys))
	for _, key := range keys {
		s.Add(key)
	}
	return s
}
func (s Set) AllItems() []string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
func (s Set) IsContain(key string) bool {
	_, ok := s[key]
	return ok
}
func (s Set) Add(key string) {
	s[key] = collections.Empty{}
}
func (s Set) Delete(key string) {
	delete(s, key)
}

type IntSet map[int]collections.Empty

func NewIntSet(keys ...int) IntSet {
	s := make(IntSet, len(keys))
	for _, key := range keys {
		s.Add(key)
	}
	return s
}
func (s IntSet) AllItems() []int {
	keys := make([]int, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
func (s IntSet) IsContain(key int) bool {
	_, ok := s[key]
	return ok
}
func (s IntSet) Add(key int) {
	s[key] = collections.Empty{}

}
func (s IntSet) Delete(key int) {
	delete(s, key)
}

type Int8Set map[int8]collections.Empty

func NewInt8Set(keys ...int8) Int8Set {
	s := make(Int8Set, len(keys))
	for _, key := range keys {
		s.Add(key)
	}
	return s
}
func (s Int8Set) AllItems() []int8 {
	keys := make([]int8, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
func (s Int8Set) IsContain(key int8) bool {
	_, ok := s[key]
	return ok
}
func (s Int8Set) Add(key int8) {
	s[key] = collections.Empty{}

}
func (s Int8Set) Delete(key int8) {
	delete(s, key)
}

type Int64Set map[int64]collections.Empty

func NewInt64Set(keys ...int64) Int64Set {
	s := make(Int64Set, len(keys))
	for _, key := range keys {
		s.Add(key)
	}
	return s
}
func (s Int64Set) AllItems() []int64 {
	keys := make([]int64, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
func (s Int64Set) IsContain(key int64) bool {
	_, ok := s[key]
	return ok
}
func (s Int64Set) Add(key int64) {
	s[key] = collections.Empty{}

}
func (s Int64Set) Delete(key int64) {
	delete(s, key)
}

type Uint8Set map[uint8]collections.Empty

func NewUint8Set(keys ...uint8) Uint8Set {
	s := make(Uint8Set, len(keys))
	for _, key := range keys {
		s.Add(key)
	}
	return s
}
func (s Uint8Set) AllItems() []uint8 {
	keys := make([]uint8, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
func (s Uint8Set) IsContain(key uint8) bool {
	_, ok := s[key]
	return ok
}
func (s Uint8Set) Add(key uint8) {
	s[key] = collections.Empty{}

}
func (s Uint8Set) Delete(key uint8) {
	delete(s, key)
}

type Uint64Set map[uint64]collections.Empty

func NewUint64Set(keys ...uint64) Uint64Set {
	s := make(Uint64Set, len(keys))
	for _, key := range keys {
		s.Add(key)
	}
	return s
}
func (s Uint64Set) AllItems() []uint64 {
	keys := make([]uint64, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
func (s Uint64Set) IsContain(key uint64) bool {
	_, ok := s[key]
	return ok
}
func (s Uint64Set) Add(key uint64) {
	s[key] = collections.Empty{}

}
func (s Uint64Set) Delete(key uint64) {
	delete(s, key)
}

type StringSet map[string]collections.Empty

func NewStringSet(keys ...string) StringSet {
	s := make(StringSet, len(keys))
	for _, key := range keys {
		s.Add(key)
	}
	return s
}
func (s StringSet) AllItems() []string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
func (s StringSet) IsContain(key string) bool {
	_, ok := s[key]
	return ok
}
func (s StringSet) Add(key string) {
	s[key] = collections.Empty{}

}
func (s StringSet) Delete(key string) {
	delete(s, key)
}
