package common

import (
        "fmt"
        "path/filepath"
)

func CheckMainFilePath(path string) error {
        var extension = filepath.Ext(path)
        if extension != ".sl" {
                return fmt.Errorf("Wrong file extension in file '%s'", path)
        }
        return nil

}

func GetFileExtension(path string) string {
        return filepath.Ext(path)
}

func GetFileNoExt(path string) string {
        return path[0:len(path) -  len(GetFileExtension(path))]
}

