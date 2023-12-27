package path

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func MoveToEnviados(dirName, fileName string) {
	pathEnviado := dirName + "/ENVIADOS/"
	oldDir := filepath.FromSlash(dirName + "/" + fileName)
	newDir := filepath.FromSlash(createDir(pathEnviado) + fileName)
	if err := os.Rename(oldDir, newDir); err != nil {
		newName := strings.ReplaceAll(strings.ToUpper(oldDir), ".TXT.GZ", "_ENVIADO.TXT.GZ")
		if err := os.Rename(oldDir, newName); err != nil {
			print(err)
		}
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func createDir(path string) string {
	isExist, _ := exists(path)
	if !isExist {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Println("Nao foi possivel criar o diretorio.")
		}
	}
	return path
}
