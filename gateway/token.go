package gateway

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v4"
	jardiniere "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func (m *Manager) loadToken() {
	// 前端发送token时带上前缀 Bearer + 空格
	m.tokener = jardiniere.New(jardiniere.Config{
		// 从请求头中的Authorization获取token进行验证
		Extractor: jardiniere.FromAuthHeader,
		// 密钥
		ValidationKeyGetter: func(token *jardiniere.Token) (any, error) {
			return []byte(C.TokenKey), nil
		},
		// 签名算法
		SigningMethod: jardiniere.SigningMethodHS256,
		// 错误处理函数
		ErrorHandler: func(ctx iris.Context, err error) {
			if err == nil {
				return
			}

			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(Result{
				Succeed: false,
				Msg:     err.Error(),
			})
		},
	})
}

func newToken(id int, username string) (string, error) {
	t := time.Now()

	token := jardiniere.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"iss":      "FANG",                                  // 签发人
		"iat":      t.Unix(),                                 // 签发时间
		"exp":      t.Add(time.Minute * 60 * 24 * 15).Unix(), // 过期时间
	})

	tokenString, err := token.SignedString([]byte(C.TokenKey))
	return tokenString, err
}
