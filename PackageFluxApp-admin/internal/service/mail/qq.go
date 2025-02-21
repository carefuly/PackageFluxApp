/**
 * Description：
 * FileName：qq.go
 * Author：CJiaの用心
 * Create：2025/2/18 01:03:57
 * Remark：
 */

package mail

import (
	"context"
	"gopkg.in/gomail.v2"
)

// QQMailSender 是 QQ 邮箱的实现
type QQMailSender struct {
	From     string // 发件人邮箱
	Password string // 授权码（不是邮箱密码）
}

func NewQQMailSender(from, password string) EmailSender {
	return &QQMailSender{
		From:     from,
		Password: password,
	}
}

func (q *QQMailSender) Send(ctx context.Context, to, subject, body string) error {
	e := gomail.NewMessage()
	e.SetHeader("From", q.From)
	e.SetHeader("To", to)
	e.SetHeader("Subject", subject)
	e.SetBody("text/html", body)

	// QQ 邮箱的 SMTP 服务器
	smtpServer := "smtp.qq.com"
	auth := gomail.NewDialer(smtpServer, 465, q.From, q.Password)

	// 发送邮件
	err := auth.DialAndSend(e)
	return err
}
