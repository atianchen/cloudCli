package encrypt

import (
	"cloudCli/utils/log"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

/**
 *
 * @author jensen.chen
 * @date 2022/7/7
 */
func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func AESEncrypt(src, k, n string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()
	key := []byte(k)
	plaintext := []byte(src)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce, _ := hex.DecodeString(n)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return fmt.Sprintf("%x", ciphertext), nil
}

func AESDecrypt(src, k, n string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()
	key := []byte(k)
	ciphertext, _ := hex.DecodeString(src)

	nonce, _ := hex.DecodeString(n)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
