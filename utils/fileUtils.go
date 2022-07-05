package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"strings"
)

/**
 *文件工具箱
 * @author jensen.chen
 * @date 2022/7/5
 */
/**
从文件的全路径或则文件名
*/
func GetFileName(filePath string) string {
	sep := ""
	if strings.Index(filePath, "/") >= 0 {
		sep = "/"
	}
	if strings.Index(filePath, "\\") >= 0 {
		sep = "\\"
	}
	if len(sep) > 0 {
		return filePath[strings.LastIndex(filePath, sep)+1 : len(filePath)]
	}
	return filePath
}

/**
获取文件HASH
*/
func GetFileHash(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
