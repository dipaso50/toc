package organizebyrule

import (
	"io/ioutil"
	"os"
	"sync"
	"tocV2/toc/application/organizebyrule"
	"tocV2/toc/infraestructure/common"
)

type OsOrganizer struct {
	common.CommonOrganizer
}

func (oss OsOrganizer) IterateOverDirectory(dirname string, fn organizebyrule.Move) error {

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
