package middleware

import (
	"dispatcher/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authMid struct {
	db *gorm.DB
}

type AuthenticationMiddleware interface {
	AuthMiddleware() gin.HandlerFunc
}

func NewAuthMiddleware(db *gorm.DB) AuthenticationMiddleware {
	return &authMid{db: db}
}

func (m *authMid) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			ctx.Abort()
			return
		}

		// find token in db!!! it is not good practice
		// TODO use exist instead
		err = m.db.First(&model.Passenger{}, "uuid = ?", token).Error
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
