package searcher

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

type Searcher struct {
	Directory string
	FileName  string
}

func (s *Searcher) FileSearch() error {
	err := filepath.Walk(s.Directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Access for the directory is denied", path)
		}

		if info.IsDir() {
			return nil
		}

		if info.Name() == s.FileName {
			fmt.Println("found the file here ", path)
			return nil
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
	return nil
}
