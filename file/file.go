package file

import (
	"os"
)

func GetOrCreate(name string) (f *os.File, err error) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err = CreateParentDirAll(name)
		if err != nil {
			return nil, err
		}
		f, err = os.Create(name)
		if err != nil {
			return nil, err
		}
	} else {
		f, err = os.OpenFile(name, os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			return nil, err
		}
	}
	return f, nil
}

func CreateParentDirAll(name string) error {
	i := len(name)
	for i > 0 && !os.IsPathSeparator(name[i-1]) {
		i--
	}
	if i > 1 {
		err := os.MkdirAll(name[:i-1], 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
