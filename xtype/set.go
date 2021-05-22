package xtype

import (
	"encoding/json"
	"strconv"
	"strings"
)

var (
	_ IntCollection    = (*intCollection)(nil)
	_ Int64Collection  = (*int64Collection)(nil)
	_ StringCollection = (*stringCollection)(nil)
	_ IntCollection    = (*intSet)(nil)
	_ Int64Collection  = (*int64Set)(nil)
	_ StringCollection = (*stringSet)(nil)
)

type Collection interface {
	IsEmpty() bool
	Size() int
	Join(string) string
	ToJSONString() (string, error)
}

type IntCollection interface {
	Collection

	Add(...int)
	Members() []int
	Contains(int) bool
}

type Int64Collection interface {
	Collection

	Add(...int64)
	Members() []int64
	Contains(int64) bool
}

type StringCollection interface {
	Collection

	Add(...string)
	Members() []string
	Contains(string) bool
}

//------------------------------------------------------------------------------

type intCollection struct {
	members []int
}

func (s *intCollection) IsEmpty() bool {
	return len(s.members) <= 0
}

func (s *intCollection) Size() int {
	return len(s.members)
}

func (s *intCollection) Join(sep string) string {
	switch len(s.members) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(s.members[0])
	}
	str := strconv.Itoa(s.members[0])
	for _, v := range s.members[1:] {
		str += sep + strconv.Itoa(v)
	}
	return str
}

func (s *intCollection) ToJSONString() (string, error) {
	b, err := json.Marshal(s.members)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (s *intCollection) Add(members ...int) {
	for _, v := range members {
		s.members = append(s.members, v)
	}
}

func (s *intCollection) Members() []int {
	return s.members
}

func (s *intCollection) Contains(val int) bool {
	for _, v := range s.members {
		if v == val {
			return true
		}
	}
	return false
}

func NewIntCollection(members ...int) *intCollection {
	if len(members) > 0 {
		return &intCollection{members: members}
	}
	return &intCollection{members: make([]int, 0)}
}

//------------------------------------------------------------------------------

type int64Collection struct {
	members []int64
}

func (s *int64Collection) IsEmpty() bool {
	return len(s.members) <= 0
}

func (s *int64Collection) Size() int {
	return len(s.members)
}

func (s *int64Collection) Join(sep string) string {
	switch len(s.members) {
	case 0:
		return ""
	case 1:
		return strconv.FormatInt(s.members[0], 10)
	}
	str := strconv.FormatInt(s.members[0], 10)
	for _, v := range s.members[1:] {
		str += sep + strconv.FormatInt(v, 10)
	}
	return str
}

func (s *int64Collection) ToJSONString() (string, error) {
	b, err := json.Marshal(s.members)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (s *int64Collection) Add(members ...int64) {
	for _, v := range members {
		s.members = append(s.members, v)
	}
}

func (s *int64Collection) Members() []int64 {
	return s.members
}

func (s *int64Collection) Contains(val int64) bool {
	for _, v := range s.members {
		if v == val {
			return true
		}
	}
	return false
}

func NewInt64Collection(members ...int64) *int64Collection {
	if len(members) > 0 {
		return &int64Collection{members: members}
	}
	return &int64Collection{members: make([]int64, 0)}
}

//------------------------------------------------------------------------------

type stringCollection struct {
	members []string
}

func (s *stringCollection) IsEmpty() bool {
	return len(s.members) <= 0
}

func (s *stringCollection) Size() int {
	return len(s.members)
}

func (s *stringCollection) Join(sep string) string {
	return strings.Join(s.members, sep)
}

func (s *stringCollection) ToJSONString() (string, error) {
	b, err := json.Marshal(s.members)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (s *stringCollection) Add(members ...string) {
	for _, v := range members {
		s.members = append(s.members, v)
	}
}

func (s *stringCollection) Members() []string {
	return s.members
}

func (s *stringCollection) Contains(val string) bool {
	for _, v := range s.members {
		if v == val {
			return true
		}
	}
	return false
}

func NewStringCollection(members ...string) *stringCollection {
	if len(members) > 0 {
		return &stringCollection{members: members}
	}
	return &stringCollection{members: make([]string, 0)}
}

//------------------------------------------------------------------------------

type intSet struct {
	intCollection

	exists map[int]bool
}

func (s *intSet) Add(members ...int) {
	for _, v := range members {
		if _, ok := s.exists[v]; !ok {
			s.members = append(s.members, v)
			s.exists[v] = true
		}
	}
}

func (s *intSet) Contains(val int) bool {
	return s.exists[val]
}

func NewIntSet(members ...int) *intSet {
	set := &intSet{
		intCollection: intCollection{
			members: make([]int, 0),
		},
		exists: make(map[int]bool),
	}
	set.Add(members...)
	return set
}

//------------------------------------------------------------------------------

type int64Set struct {
	int64Collection

	exists map[int64]bool
}

func (s *int64Set) Add(members ...int64) {
	for _, v := range members {
		if _, ok := s.exists[v]; !ok {
			s.members = append(s.members, v)
			s.exists[v] = true
		}
	}
}

func (s *int64Set) Contains(val int64) bool {
	return s.exists[val]
}

func NewInt64Set(members ...int64) *int64Set {
	set := &int64Set{
		int64Collection: int64Collection{
			members: make([]int64, 0),
		},
		exists: make(map[int64]bool),
	}
	set.Add(members...)
	return set
}

//------------------------------------------------------------------------------

type stringSet struct {
	stringCollection

	exists map[string]bool
}

func (s *stringSet) Add(members ...string) {
	for _, v := range members {
		if _, ok := s.exists[v]; !ok {
			s.members = append(s.members, v)
			s.exists[v] = true
		}
	}
}

func (s *stringSet) Contains(val string) bool {
	return s.exists[val]
}

func NewStringSet(members ...string) *stringSet {
	set := &stringSet{
		stringCollection: stringCollection{
			members: make([]string, 0),
		},
		exists: make(map[string]bool),
	}
	set.Add(members...)
	return set
}
