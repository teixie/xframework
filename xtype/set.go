package xtype

import (
	"fmt"
	"strings"
)

var (
	_ Collection = (*intSet)(nil)
	_ Collection = (*int64Set)(nil)
	_ Collection = (*stringSet)(nil)
)

type Collection interface {
	IsEmpty() bool
	Size() int
	Join(string) string
}

//------------------------------------------------------------------------------

type intSet struct {
	members []int
	exists  map[int]bool
}

func (s *intSet) IsEmpty() bool {
	return len(s.members) <= 0
}

func (s *intSet) Size() int {
	return len(s.members)
}

func (s *intSet) Join(sep string) string {
	str := ""
	for _, v := range s.members {
		str += fmt.Sprint(v) + sep
	}
	return strings.Trim(str, sep)
}

func (s *intSet) Add(members ...int) {
	for _, v := range members {
		if _, ok := s.exists[v]; !ok {
			s.members = append(s.members, v)
			s.exists[v] = true
		}
	}
}

func (s *intSet) Members() []int {
	return s.members
}

func NewIntSet(members ...int) *intSet {
	set := &intSet{
		exists: make(map[int]bool),
	}
	set.Add(members...)
	return set
}

//------------------------------------------------------------------------------

type int64Set struct {
	members []int64
	exists  map[int64]bool
}

func (s *int64Set) IsEmpty() bool {
	return len(s.members) <= 0
}

func (s *int64Set) Size() int {
	return len(s.members)
}

func (s *int64Set) Join(sep string) string {
	str := ""
	for _, v := range s.members {
		str += fmt.Sprint(v) + sep
	}
	return strings.Trim(str, sep)
}

func (s *int64Set) Add(members ...int64) {
	for _, v := range members {
		if _, ok := s.exists[v]; !ok {
			s.members = append(s.members, v)
			s.exists[v] = true
		}
	}
}

func (s *int64Set) Members() []int64 {
	return s.members
}

func NewInt64Set(members ...int64) *int64Set {
	set := &int64Set{
		exists: make(map[int64]bool),
	}
	set.Add(members...)
	return set
}

//------------------------------------------------------------------------------

type stringSet struct {
	members []string
	exists  map[string]bool
}

func (s *stringSet) IsEmpty() bool {
	return len(s.members) <= 0
}

func (s *stringSet) Size() int {
	return len(s.members)
}

func (s *stringSet) Join(sep string) string {
	return strings.Join(s.members, sep)
}

func (s *stringSet) Add(members ...string) {
	for _, v := range members {
		if _, ok := s.exists[v]; !ok {
			s.members = append(s.members, v)
			s.exists[v] = true
		}
	}
}

func (s *stringSet) Members() []string {
	return s.members
}

func NewStringSet(members ...string) *stringSet {
	set := &stringSet{
		exists: make(map[string]bool),
	}
	set.Add(members...)
	return set
}
