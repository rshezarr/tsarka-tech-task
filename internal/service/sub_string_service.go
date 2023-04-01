package service

import (
	"strings"
	"unicode/utf8"
)

type Finder interface {
	FindMaxSubstring(s string) string
}

type FindService struct{}

func NewFindService() *FindService {
	return &FindService{}
}

func (f *FindService) FindMaxSubstring(str string) string {
	var result string
	for i := 0; i < len(str); i++ {
		sub := ""
		for j := i; j < len(str); j++ {
			if index := strings.IndexByte(sub, str[j]); index == -1 {
				sub += string(str[j])
			} else {
				break
			}
		}
		if utf8.RuneCountInString(sub) > utf8.RuneCountInString(result) {
			result = sub
		}
	}
	return result
}
