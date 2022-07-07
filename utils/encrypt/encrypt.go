package encrypt

import (
	"crypto/md5"
	"encoding/hex"
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
