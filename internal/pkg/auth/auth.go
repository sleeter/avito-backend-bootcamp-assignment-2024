package auth

import (
	"backend-bootcamp-assignment-2024/internal/model/dto/request"
	"backend-bootcamp-assignment-2024/internal/model/entity"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"slices"
	"strings"
	"time"
)

func AuthModerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := GetToken(ctx)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		userRole := claims["role"].(string)
		if ok && token.Valid && userRole == entity.USERTYPE_MODERATOR {
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := GetToken(ctx)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		userRole := claims["role"].(string)
		if ok && token.Valid && slices.Contains(entity.UserTypes, userRole) {
			ctx.Set("User-Type", userRole)
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func GetToken(ctx *gin.Context) (*jwt.Token, error) {
	bearerToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) != 2 {
		return nil, errors.New("invalid token")
	}
	tokenString := splitToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		//TODO: change secret
		return []byte("secret"), nil
	})
	return token, err
}

func CreateDummyJWT(role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	//TODO: os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateUserJWT(req request.Login, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": req.Id,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"role": role,
	})
	//TODO: os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
