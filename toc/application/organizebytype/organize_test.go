package organizebytype

import (
	"fmt"
	"testing"
	"tocV2/toc/domain"
)

func Test_MoveIfMatch(t *testing.T) {
	rman := NewRuleManager()
	rman.AddRule("txt$", "txt")

	var test string

	mockedOrganizeByType := &OrganizeByTypeMock{
		CreateDirectoryFunc: func(dirname string) error {
			panic("mock out the CreateDirectory method")
		},
		ExistsDirectoryFunc: func(dirname string) (bool, error) {
			panic("mock out the ExistsDirectory method")
		},
		IterateOverDirectoryFunc: func(dirname string, fn Move) error {
			panic("mock out the IterateOverDirectory method")
		},
		MoveFileFunc: func(file domain.FileOrganizer) error {
			test = file.Name
			return nil
		},
	}

	ser := NewService(mockedOrganizeByType, rman)
	ser.MoveIfMatch("sample.txt", "txt")

	if test != "sample.txt" {
		t.Fatalf("Expected sample.txt got %v ", test)
	}
}

func Test_MoveIfMatchNotMatch(t *testing.T) {
	rman := NewRuleManager()
	rman.AddRule("txt$", "txt")

	var test string

	mockedOrganizeByType := &OrganizeByTypeMock{
		MoveFileFunc: func(file domain.FileOrganizer) error {
			test = file.Name
			return nil
		},
	}

	ser := NewService(mockedOrganizeByType, rman)
	ser.MoveIfMatch("sample.pdf", "txt")

	if len(test) != 0 {
		t.Fatalf("Expected nothing got %v ", test)
	}
}

func Test_NotExistsDirectory(t *testing.T) {
	rman := NewRuleManager()

	mockedOrganizeByType := &OrganizeByTypeMock{
		CreateDirectoryFunc: func(dirname string) error {
			panic("mock out the CreateDirectory method")
		},
		ExistsDirectoryFunc: func(dirname string) (bool, error) {
			return false, fmt.Errorf("Dir %s does not exists ", dirname)
		},
		IterateOverDirectoryFunc: func(dirname string, fn Move) error {
			panic("mock out the IterateOverDirectory method")
		},
		MoveFileFunc: func(file domain.FileOrganizer) error {
			panic("mock out the MoveFileFunc method")
		},
	}

	ser := NewService(mockedOrganizeByType, rman)

	err := ser.Execute("sampleDir")

	if err == nil {
		t.Errorf("Error on ExistsDirectory function")
	}
}
