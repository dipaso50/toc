package organizebyrule

import (
	"fmt"
	"tocV2/toc/domain"
)

type Move func(filename, filePath string) error

//go:generate moq -out organizebyrule_mock.go . OrganizeByRule

type OrganizeByRule interface {
	domain.Organize
	IterateOverDirectory(dirname string, fn Move) error
}

type Service struct {
	operations  OrganizeByRule
	ruleManager *RuleManager
}

func NewService(op OrganizeByRule, ruleMan *RuleManager) Service {
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
