package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
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

/**
讲byte[]转换为HASH值字符串
*/
func ConvertReaderToHash(reader io.Reader) (string, error) {
	hasher := sha256.New()
	if _, err := io.Copy(hasher, reader); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

/**
讲byte[]转换为HASH值字符串
*/
func ConvertByteToHash(content []byte) (string, error) {
	hasher := sha256.New()
	hasher.Write(content)
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func GetFileStringContent(filePath string) (string, error) {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func GetStringContent(reader io.Reader) (string, error) {
	var sb strings.Builder
	buf := make([]byte, 256)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}
		sb.Write(buf[:n])
	}
	return sb.String(), nil
}
