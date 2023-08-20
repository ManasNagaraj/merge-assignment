package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merge/shopping-card/internal/model"
	"github.com/merge/shopping-card/internal/store"
)

func AuthMiddleware(ac store.AccessTokenStore, userType string) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		tokenString = tokenString[7:]

		at, err := ac.FindByAccessToken(c, tokenString)

		if err != nil || at == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		}

		if userType == at.Role || at.Role == string(model.RoleAdmin) {
			c.Set("userId", at.UserID)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		}
	}
}
