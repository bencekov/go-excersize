package api

import (
	"testing"

	"github.com/golang/mock/gomock"
)

//go:generate mockgen -build_flags=--mod=mod -package api -destination ./mock_logger.go -source=../../internal/logging/interface.go
//go:generate mockgen -build_flags=--mod=mod -package api -destination ./mock_api.go -source=./interface.go

func TestRemoveVowelsSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := NewMockLoggerInterface(ctrl)
	mockLogger.EXPECT().Infof(gomock.Any(), gomock.Any()).Times(1).Return()

	testString := "test string"
	resultString := "tst strng"

	r, err := NewService(mockLogger).RemoveVowels(testString)

	if r != resultString {
		t.Fatalf("expected response to be %s not  %s", resultString, r)
	}

	if err != nil {
		t.Fatalf("expected error to be nil not  %v", err)
	}
}

func TestRemoveVowelsFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := NewMockLoggerInterface(ctrl)
	mockLogger.EXPECT().Errorf(gomock.Any(), gomock.Any()).Times(1).Return()

	testString := ""
	resultString := ""

	r, err := NewService(mockLogger).RemoveVowels(testString)

	if r != resultString {
		t.Fatalf("expected response to be %s not  %s", resultString, r)
	}

	if err.Error() != "Input is empty string" {
		t.Fatalf("expected error message to be `Input is empty string` not  %v", err)
	}
}

func TestGetCounters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := NewMockLoggerInterface(ctrl)
	mockLogger.EXPECT().Infof(gomock.Any(), gomock.Any()).Times(4).Return()

	testString := "test"
	resultCount := 3

	s := NewService(mockLogger)
	s.RemoveVowels(testString)
	s.RemoveVowels(testString)
	s.RemoveVowels(testString)
	r := s.GetCounter()

	if r != resultCount {
		t.Fatalf("expected response to be %v not  %v", resultCount, r)
	}

}
