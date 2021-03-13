package organizebymodification

import (
	"io/ioutil"
	"os"
	"sync"
	"time"
	"tocV2/toc/application/organizebymodification"
	"tocV2/toc/infraestructure/common"
)

type OsOrganizer struct {
	common.CommonOrganizer
}

func (oss OsOrganizer) IterateOverDirectory(dirname string, fn organizebymodification.MoveByModification) error {

	var fileInfo []os.FileInfo
	var err error

	if fileInfo, err = ioutil.ReadDir(dirname); err != nil {
		return err
	}

	var wg sync.WaitGroup
	sem := make(chan int, oss.MaxGoroutines())

	for _, file := range fileInfo {

		if !file.IsDir() {
			wg.Add(1)
			sem <- 1

			go func(filename, dname string, mod time.Time) {

				defer func() {
					wg.Done()
					<-sem
				}()

				fn(filename, dname, mod)

			}(file.Name(), dirname, file.ModTime())

		}
	}

	wg.Wait()
	return nil
}
