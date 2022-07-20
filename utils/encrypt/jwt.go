package encrypt

import "github.com/golang-jwt/jwt"

/**
 * JWT的实现
 * @author jensen.chen
 * @date 2022/7/7
 */
var mySigningKey = []byte("yonyou-cloud-cli-2022-07-07")

func GenerateToken[T jwt.Claims](data T) (string, error) {
	// 使用HS256加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	signToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return signToken, nil

}
func ParseToken(signToken string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(signToken, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if token != nil && token.Valid {
		return nil
	} else {
		return err
	}

}
