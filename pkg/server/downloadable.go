package server

import (
	"fmt"
	"os"
	"path/filepath"
)

type Downloadable struct {
	Name string
	Path string
}

const downloadFolder = "download"

func getDownloadables() []Downloadable {
	var dls []Downloadable
	filepath.Walk(downloadFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("foo: %s", err.Error())
			}

			if !info.IsDir() {
				dl := Downloadable{
					Name: info.Name(),
					Path: path,
				}
				dls = append(dls, dl)
			}
			return nil
		})

	return dls
}
