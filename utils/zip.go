package utils

import (
	"github.com/mholt/archiver"
	"fmt"
)

func ZipDir(dir string, name string) (output string, err error) {
	err = archiver.Zip.Make(name, []string{dir})
	output = fmt.Sprintf("%s/%s", dir, name)
	return output, nil
}
