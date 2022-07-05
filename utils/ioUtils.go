package utils

import (
	"os"
	"strings"
)

/**
 * IO工具箱
 * @author jensen.chen
 * @date 2022/7/5
 */

/**
获取内部文件的实际路径
*/
func GetFilePath(file string) (string, error) {
	pwd, err1 := os.Getwd()
	_, err := os.Stat(file)
	if err != nil {
		if err1 != nil {
			return "", err1
		} else {
			_, err3 := os.Stat(pwd + "/" + file)
			if err3 != nil {
				return "", err3
			} else {
				return pwd + "/" + file, nil
			}
		}
	} else {
		return file, nil
	}
}

func CreateFileDirectory(filePath string) bool {
	sep := ""
	if strings.Index(filePath, "/") >= 0 {
		sep = "/"
	}
	if strings.Index(filePath, "\\") >= 0 {
		sep = "\\"
	}
	if len(sep) > 0 {
		dir := filePath[0:strings.LastIndex(filePath, sep)]
		_, err := os.Stat(dir)
		if err != nil {
			if os.IsNotExist(err) {
				mkErr := os.MkdirAll(dir, os.ModePerm)
				if mkErr == nil {
					return true
				} else {
					return false
				}
			}
		} else {
			return true
		}
	}
	return false
}

func CreateDirectory(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			mkErr := os.MkdirAll(dir, os.ModePerm)
			if mkErr == nil {
				return true
			} else {
				return false
			}
		}
	} else {
		return true
	}
	return false
}
