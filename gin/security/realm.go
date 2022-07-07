package security

import "github.com/golang-jwt/jwt"

/**
 *
 * @author jensen.chen
 * @date 2022/7/7
 */
type Realm struct {
	Id   string `json:"id"`
	Code string `json:"code"`
	jwt.StandardClaims
}
