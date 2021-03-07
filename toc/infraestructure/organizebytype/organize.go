package organizebytype

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"tocV2/toc/application/organizebytype"
	"tocV2/toc/domain"
)

type OsOrganizer struct{}

func (oss OsOrganizer) ExistsDirectory(dirname string) (bool, error) {

	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		return false, fmt.Errorf("Directorio %s no existe ", dirname)
	}

	return true, nil
}
func (oss OsOrganizer) CreateDirectory(dirname string) error {

	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		fmt.Printf("Creando directorio %s\n", dirname)
		return os.MkdirAll(dirname, 0700)
	}
	return nil
}

func (oss OsOrganizer) MoveFile(file domain.FileOrganizer) error {
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

func (oss OsOrganizer) IterateOverDirectory(dirname string, fn organizebytype.Move) error {

	var fileInfo []os.FileInfo
	var err error

	if fileInfo, err = ioutil.ReadDir(dirname); err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, file := range fileInfo {

		if !file.IsDir() {
			wg.Add(1)

			go func(filename, dname string) {

				defer wg.Done()
				fn(filename, dname)

			}(file.Name(), dirname)

		}
	}

	wg.Wait()
	return nil
}
