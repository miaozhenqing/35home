package middleware

import (
	"net/http"
	"strings"

	"35home/common"
	"35home/dto"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT声明结构
type Claims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// JWTMiddleware JWT认证中间件
func JWTMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Authorization header获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, dto.NewErrorResponse(common.ErrInvalidRequest))
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, dto.NewErrorResponse(common.ErrInvalidRequest))
			c.Abort()
			return
		}

		// 解析token
		tokenString := parts[1]
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, dto.NewErrorResponse(common.ErrInvalidRequest))
			c.Abort()
			return
		}

		// 提取claims
		if claims, ok := token.Claims.(*Claims); ok {
			// 将用户信息存储到上下文中
			c.Set("userID", claims.UserID)
			c.Set("username", claims.Username)
			c.Set("email", claims.Email)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, dto.NewErrorResponse(common.ErrInvalidRequest))
			c.Abort()
			return
		}
	}
}
