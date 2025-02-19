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
	return "AnonymousUser"
	// token := c.GetHeader("Authorization")
	// if token == "" {
	// 	return "AnonymousUser"
	// }
	// seg := strings.SplitN(token, " ", 2)
}
