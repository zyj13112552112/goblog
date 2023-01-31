package middleware

//jwt，token认证
import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"goblog/utils"
	"goblog/utils/errmsg"
	"net/http"
	"strings"
	"time"
)

var JwKey = []byte(utils.JwtKey) //jwt密钥

type MyClaims struct {
	Username string `json:"username,omitempty"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	var expireTime *jwt.Time = &jwt.Time{time.Now().Add(10 * time.Hour)} //10个小时
	SetClaim := MyClaims{
		Username: username,
		//Password:       password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "ginblog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaim)
	ToKen, err := token.SignedString(JwKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return ToKen, errmsg.SUCCSE
}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	Token, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwKey, nil
	})
	if key, _ := Token.Claims.(*MyClaims); Token.Valid {
		return key, errmsg.SUCCSE
	} else {
		return nil, errmsg.ERROR
	}
}

var code int

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		} //token不存在

		checktoken := strings.SplitN(tokenHeader, " ", 2)
		if len(checktoken) != 2 && checktoken[0] != "Bearer" { //Bearer author
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		} //token格式错误

		token, i := CheckToken(checktoken[1])
		if i == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		} //token错误

		if time.Now().Unix() > token.ExpiresAt.Unix() {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		} //token过期

		c.Set("username", token.Username)
		c.Next()
	}
}
