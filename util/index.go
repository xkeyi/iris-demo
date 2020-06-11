package util

import (
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/iris-contrib/middleware/jwt"
)

// 生成hash密码
func GeneratePassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// 验证密码
func ValidatePassword(password string, hashed string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	}

	return true
}

// 生成token
func GetJWTString(id int64) (string, error) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// 根据需求，可以存一些必要的数据
		"userId":   id, // 用户ID

		// 签发人
		"iss": "iris",
		// 签发时间
		"iat": time.Now().Unix(),
		// 设定过期时间，设置120分钟过期
		"exp": time.Now().Add(120 * time.Minute * time.Duration(1)).Unix(),
	})

	// 使用设置的秘钥，签名生成jwt字符串
	tokenString, err := token.SignedString([]byte("My Secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 获取登录用户ID
func GetTokenUserId(ctx iris.Context) int64 {
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	userId := int64(jwtInfo["userId"].(float64))

	return userId
}