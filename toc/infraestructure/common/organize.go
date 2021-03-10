package common

import (
	"fmt"
	"os"
	"path/filepath"
	"tocV2/toc/domain"
)

type CommonOrganizer struct{}

const MAX_GOROUTINES = 20

func (oss CommonOrganizer) ExistsDirectory(dirname string) (bool, error) {

	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		return false, fmt.Errorf("Directorio %s no existe ", dirname)
	}

	return true, nil
}
func (oss CommonOrganizer) CreateDirectory(dirname string) error {

	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		fmt.Printf("Creando directorio %s\n", dirname)
		return os.MkdirAll(dirname, 0700)
	}
	return nil
}

func (oss CommonOrganizer) MoveFile(file domain.FileOrganizer) error {
	ori := filepath.Join(file.Origin, file.Name)
	destinationFolder := filepath.Join(file.Origin, file.Destination)
	dest := filepath.Join(destinationFolder, file.Name)

	oss.CreateDirectory(destinationFolder)

	fmt.Printf("Moviendo %s -> %s \n", ori, dest)

	err := os.Rename(ori, dest)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
