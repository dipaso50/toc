package organizebymodification

import (
	"fmt"
	"time"
	"tocV2/toc/domain"
)

type MoveByModification func(filename, filePath string, fileModification time.Time) error

//go generate moq -out OrganizeByModification_mock.go . OrganizeByModification

type OrganizeByModification interface {
	domain.Organize
	IterateOverDirectory(dirname string, fn MoveByModification) error
}

type Service struct {
	operations OrganizeByModification
}

func NewService(op OrganizeByModification) Service {
	return Service{operations: op}
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

func (s Service) MoveIfMatch(filename, filePath string, fileModification time.Time) error {

	destination := createLocation(fileModification)

	ff, err := domain.NewFileOrganizer(filePath, destination, filename)

	if err != nil {
		return err
	}

	return s.operations.MoveFile(ff)
}

func createLocation(fTime time.Time) string {
	return fmt.Sprintf("%04d/%02d/%02d", fTime.Year(), fTime.Month(), fTime.Day())
}
