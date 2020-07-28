package main

import (
	"os"
	"path/filepath"
)

func CreateDir(path string) error {
	ex, err := DirExists(path)
	if err != nil {
		return err
	}
	if !ex {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func GetCurrentDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
