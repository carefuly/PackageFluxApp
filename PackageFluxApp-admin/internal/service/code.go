/**
 * Description：
 * FileName：code.go
 * Author：CJiaの用心
 * Create：2025/2/19 16:32:48
 * Remark：
 */

package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/carefuly/PackageFluxApp/internal/repository"
	"github.com/carefuly/PackageFluxApp/internal/service/mail"
	"math/big"
)

var (
	ErrCodeSendTooMany   = repository.ErrCodeSendTooMany
	ErrUserBlocked       = repository.ErrUserBlocked
	ErrCodeNotFound      = repository.ErrCodeNotFound
	ErrCodeVerifyTooMany = repository.ErrCodeVerifyTooMany
	ErrCodeIncorrect     = repository.ErrCodeIncorrect
)

type CodeService interface {
	Send(ctx context.Context, to, subject, biz string) (string, error)
	Verify(ctx context.Context, biz, to, inputCode string) (bool, error)
}

type codeService struct {
	repo repository.CodeRepository
	mail mail.EmailSender
}

func NewCodeService(repo repository.CodeRepository, mail mail.EmailSender) CodeService {
	return &codeService{
		repo: repo,
		mail: mail,
	}
}

func (svc *codeService) Send(ctx context.Context, to, subject, biz string) (string, error) {
	code := svc.generate()
	body := svc.body(code)
	err := svc.repo.Set(ctx, biz, to, code)
	if err != nil {
		return code, err
	}
	return code, svc.mail.Send(ctx, to, subject, body)
}

func (svc *codeService) Verify(ctx context.Context, biz, to, inputCode string) (bool, error) {
	ok, err := svc.repo.Verify(ctx, biz, to, inputCode)

	switch {
	case errors.Is(err, repository.ErrUserBlocked):
		return false, repository.ErrUserBlocked
	case errors.Is(err, repository.ErrCodeNotFound):
		return false, repository.ErrCodeNotFound
	case errors.Is(err, repository.ErrCodeVerifyTooMany):
		return false, repository.ErrCodeVerifyTooMany
	case errors.Is(err, repository.ErrCodeIncorrect):
		return false, repository.ErrCodeIncorrect
	}

	return ok, err
}

func (svc *codeService) generate() string {
	// 生成6位随机数字
	code, err := svc.generateRandomCode(6)
	if err != nil {
		fmt.Println("生成失败:", err)
		return code
	}
	return code
}

func (svc *codeService) generateRandomCode(length int) (string, error) {
	const digits = "0123456789"
	code := make([]byte, length)
	for i := range code {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		code[i] = digits[num.Int64()]
	}
	return string(code), nil
}

func (svc *codeService) body(code string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 600px;
            margin: 20px auto;
            background-color: #ffffff;
            border-radius: 8px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            overflow: hidden;
        }
        .header {
            background-color: #007BFF;
            color: #ffffff;
            text-align: center;
            padding: 20px;
        }
        .header h1 {
            margin: 0;
        }
        .content {
            padding: 20px;
        }
        .content p {
            line-height: 1.6;
            margin-bottom: 20px;
        }
        .code-box {
            background-color: #f0f0f0;
            padding: 15px;
            text-align: center;
            border-radius: 4px;
            font-size: 24px;
            font-weight: bold;
            color: #007BFF;
        }
        .footer {
            background-color: #f4f4f4;
            text-align: center;
            padding: 10px;
            color: #777;
            font-size: 12px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>验证码邮件</h1>
        </div>
        <div class="content">
            <p>尊敬的用户，您好！</p>
            <p>您正在进行相关操作，本次的验证码是：</p>
            <div class="code-box">
                <!-- 这里替换为实际的验证码 -->
                %s
            </div>
            <p>验证码的有效期为 10 分钟，请尽快完成验证。如果这不是您本人的操作，请忽略此邮件。</p>
        </div>
        <div class="footer">
            <p>此邮件由系统自动发送，请勿直接回复。</p>
        </div>
    </div>
</body>
</html>
    `, code)
}
