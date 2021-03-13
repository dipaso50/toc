package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"tocV2/toc/domain"
)

type CommonOrganizer struct{}

const MAX_GOROUTINES = 20
const GOROUTINES_LIMIT = "TOC_GOROUTIMES_LIMIT"

func (oss CommonOrganizer) MaxGoroutines() int {
	value, exists := os.LookupEnv(GOROUTINES_LIMIT)
	if !exists {
		return MAX_GOROUTINES
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		return MAX_GOROUTINES
	}
	return v
}

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
