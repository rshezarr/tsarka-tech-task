package service

import "regexp"

type EmailChecker interface {
	CheckEmail(s []byte) map[string][]string
}

type EmailCheckService struct{}

func NewEmailCheckService() *EmailCheckService {
	return &EmailCheckService{}
}

func (e *EmailCheckService) CheckEmail(s []byte) map[string][]string {
	emailRegex := regexp.MustCompile(`(?i)\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b`)

	emails := emailRegex.FindAllString(string(s), -1)

	results := make(map[string][]string)
	results["emails"] = emails

	return results
}
