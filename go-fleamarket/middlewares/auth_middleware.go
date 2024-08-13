package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-fleamarket/services"
	"net/http"
	"strings"
)

func AuthMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			// 後続のハンドラー関数の実行をストップできるが現在のハンドラー関数の実行はストップしない
			c.AbortWithStatus(http.StatusUnauthorized)
			// returnで現在のハンドラー関数の実行はストップ
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			// 後続のハンドラー関数の実行をストップできるが現在のハンドラー関数の実行はストップしない
			c.AbortWithStatus(http.StatusUnauthorized)
			// returnで現在のハンドラー関数の実行はストップ
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		user, err := authService.GetUserFromToken(tokenString)
		if err != nil {
			// 後続のハンドラー関数の実行をストップできるが現在のハンドラー関数の実行はストップしない
			c.AbortWithStatus(http.StatusUnauthorized)
			// returnで現在のハンドラー関数の実行はストップ
			return
		}

		c.Set("user", user)

		// 必ず呼ぶ
		c.Next()
	}
}
