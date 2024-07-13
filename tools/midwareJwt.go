package tools

import (
	"fmt"
	"frz/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		Response(ctx, http.StatusOK, 401, gin.H{}, "权限不足")
		ctx.Abort()
		return
	}
	tokenString = tokenString[7:]
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		Response(ctx, http.StatusOK, 401, gin.H{
			"error": err,
		}, "token解析失败")
		ctx.Abort()
		return
	}
	key := claims.Key
	if key != models.ACCOUNT {
		Response(ctx, http.StatusOK, 401, gin.H{}, "用户不存在")
		ctx.Abort()
		return
	}
	ctx.Set("account", key)
	ctx.Next()
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return token, claims, err
}
