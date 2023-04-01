package service

import "strings"

type Finder interface {
	FindMaxSubstring(s string) string
}

type FindService struct{}

func NewFindService() *FindService {
	return &FindService{}
}

func (f *FindService) FindMaxSubstring(str string) string {
	s := strings.Builder{}
	s.Grow(len(str))
	m := make(map[rune]bool, len(str))

	for _, v := range str {
		if !m[v] {
			s.WriteRune(v)
			m[v] = true
		}
	}

	return s.String()
}
