/**
 * @Description https://pkg.go.dev/github.com/golang-jwt/jwt#section-readme
 **/
package test

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

// 设置声明
type UserClaims struct {
	*jwt.StandardClaims
	Name string
	Uid  uint
}

//func (u UserClaims) Valid() error {
//	panic("implement me")
//}

// 创建Jwt
func TestCreateJWTByNewWithClaims(t *testing.T) {
	nowTime := time.Now()
	// 使用NewWithClaims创建声明
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		StandardClaims: &jwt.StandardClaims{
			// 设置有效期为5分钟
			ExpiresAt: nowTime.Add(time.Hour * 15).Unix(),
			Issuer:    "Go学习",         // 签发人
			IssuedAt:  nowTime.Unix(), // 签发时间
		},
		Name: "张三",
		Uid:  100,
	})
	// 使用指定的secret签名,获取字符串token
	signedString, err := claims.SignedString([]byte("这是我定义的secret"))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("生成JWT:", signedString)
}

// 解析JWT
func TestValidJWT(t *testing.T) {
	jwtStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDc0MDgxMjMsImlhdCI6MTY0NzM1NDEyMywiaXNzIjoiR2_lrabkuaAiLCJOYW1lIjoi5byg5LiJIiwiVWlkIjoxMDB9.Ftu206lTqHMgMXOnaCwk8yX_t1g84eUfKvs3yvUOzlU"
	// 解析jwt,第三个参数是个函数，返回生成jwt设置的secret
	token, err := jwt.ParseWithClaims(jwtStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("这是我定义的secret"), nil
	})
	if err != nil {
		t.Error("解析失败:", err.Error())
		return
	}
	// 断言类型
	claim, ok := token.Claims.(*UserClaims)
	// 验证
	if !ok || !token.Valid {
		t.Error("解析失败,Token不合法!")
		return
	}
	fmt.Printf("解析结果: %+v 内嵌: %+v \n", claim, *claim.StandardClaims)
}
