package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"homestay.com/nguyenduy/internal/app/repositories"
)

func PermissionCheck(permissionRepository repositories.PermissionRepository, requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No role found"})
			return
		}

		userPermissions, err := permissionRepository.GetPermissionsByRoleID(role.(string))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get permissions"})
			return
		}

		if len(userPermissions) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No permissions found"})
			return
		}

		hasPermission := false
		for _, p := range userPermissions {
			if strings.EqualFold(p.Code, requiredPermission) {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}

		c.Next()
	}
}
