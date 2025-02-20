/**
 * Description：
 * FileName：request_utils.go
 * Author：CJiaの用心
 * Create：2025/2/17 14:34:33
 * Remark：
 */

package requestUtils

import (
	"github.com/gin-gonic/gin"
)

// GetRequestUser 获取请求user
func GetRequestUser(c *gin.Context) string {
	email, ok := c.MustGet("email").(string)
	if !ok {
		return "AnonymousUser"
	}
	return email
}
