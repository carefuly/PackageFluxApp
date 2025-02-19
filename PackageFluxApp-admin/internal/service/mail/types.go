/**
 * Description：
 * FileName：types.go
 * Author：CJiaの用心
 * Create：2025/2/18 01:03:40
 * Remark：
 */

package mail

import "context"

// EmailSender 是一个通用的邮件发送接口
type EmailSender interface {
	Send(ctx context.Context, to, subject, body string) error
}
