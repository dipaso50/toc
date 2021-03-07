package organizebytype

import (
	"fmt"
	"tocV2/toc/domain"
)

type Move func(filename, filePath string) error

//go:generate moq -out OrganizeByType_test.go . OrganizeByType

type OrganizeByType interface {
	ExistsDirectory(dirname string) (bool, error)
	CreateDirectory(dirname string) error
	MoveFile(file domain.FileOrganizer) error
	IterateOverDirectory(dirname string, fn Move) error
}

type Service struct {
	operations  OrganizeByType
	ruleManager *RuleManager
}

func NewService(op OrganizeByType, ruleMan *RuleManager) Service {
	return Service{operations: op, ruleManager: ruleMan}
}

func (s Service) Execute(dirname string) error {
	exists, err := s.operations.ExistsDirectory(dirname)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("Directorio %s no existe", dirname)
	}

	s.operations.IterateOverDirectory(dirname, s.MoveIfMatch)

	return nil
}

func (s Service) MoveIfMatch(filename, filePath string) error {

	fileMatch, destinationFolder := s.ruleManager.FileMatch(filename)

	if !fileMatch {
		return nil
	}

	ff, err := domain.NewFileOrganizer(filePath, destinationFolder, filename)

	if err != nil {
		return err
	}

	return s.operations.MoveFile(ff)
}
