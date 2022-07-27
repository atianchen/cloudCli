package server

import (
	"cloudCli/utils/encrypt"
	"time"
)

const (
	AES_KEY   = "yycloudcli"
	AES_NONCE = "9845a9dc986732dcab7cdd5a"
	FORMAT    = "01#01&"
)

/**
 *
 * @author jensen.chen
 * @date 2022/7/27
 */
func Encrypt(content string) (string, error) {
	now := time.Now()
	return encrypt.AESEncrypt(content, AES_KEY+now.Format(FORMAT), AES_NONCE)
}

func Decrypt(content string) (string, error) {
	now := time.Now()
	return encrypt.AESDecrypt(content, AES_KEY+now.Format(FORMAT), AES_NONCE)
}
