package service

type Service struct {
	Finder       Finder
	EmailChecker EmailChecker
}

func NewService() *Service {
	return &Service{
		Finder:       NewFindService(),
		EmailChecker: NewEmailCheckService(),
	}
}
