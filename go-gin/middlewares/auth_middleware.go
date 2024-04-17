package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-gin/services"
	"net/http"
	"strings"
)

func AuthMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized) // 現在のハンドラーは停止されない
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized) // 現在のハンドラーは停止されない
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		user, err := authService.GetUserFromToken(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized) // 現在のハンドラーは停止されない
			return
		}

		ctx.Set("user", user) // requestのライフサイクル中に生存する
		ctx.Next()            // 次のミドルウェアまたは目的の処理に移す
	}
}
