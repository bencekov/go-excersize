package api

import (
	"errors"

	"github.com/bencekov/go-exercise/internal/logging"
)

type Service struct {
	logger  logging.LoggerInterface
	counter int
}

func NewService(logger logging.LoggerInterface) *Service {
	s := new(Service)
	s.logger = logger
	return s
}

func (s *Service) RemoveVowels(input string) (string, error) {
	if len(input) == 0 {
		s.logger.Errorf("Issue with input %s", input)
		return "", errors.New("Input is empty string")
	}
	s.CounterAdd()
	ret := ""

	for _, i := range input {
		ok := s.CheckVowel(string(i))
		if !ok {
			ret += string(i)
		}
	}
	s.logger.Infof("New message: %s", ret)
	return ret, nil
}

func (s *Service) CheckVowel(input string) bool {
	vowels := "aeiouAEIOU"

	for _, i := range vowels {
		if string(i) == input {
			return true
		}
	}
	return false
}

func (s *Service) CounterAdd() {
	s.counter += 1
}

func (s *Service) GetCounter() int {
	s.logger.Infof("Counter status: %v", s.counter)
	return s.counter
}
