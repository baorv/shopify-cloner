package utils

import "os"

func CreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		return err
	}
	return nil
}