/**
 * Description：
 * FileName：types.go
 * Author：CJiaの用心
 * Create：2025/2/19 16:07:50
 * Remark：
 */

package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/carefuly/PackageFluxApp/pkg/response"
	"github.com/gin-gonic/gin"
	"io"
)

// LoggingReader 自定义读取器，用于记录读取的内容
type LoggingReader struct {
	reader io.Reader
	buffer *bytes.Buffer
}

func (lr *LoggingReader) Read(p []byte) (n int, err error) {
	n, err = lr.reader.Read(p)
	if n > 0 {
		lr.buffer.Write(p[:n])
	}
	return
}

func (lr *LoggingReader) format() any {
	var result map[string]any
	err := json.Unmarshal([]byte(lr.buffer.String()), &result)
	if err != nil {
		return nil
	}
	return result
}

// customGinResponseWriter 自定义 Gin 响应写入器
type customGinResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (crw *customGinResponseWriter) Write(b []byte) (int, error) {
	// 记录响应数据
	crw.body.Write(b)
	return crw.ResponseWriter.Write(b)
}

func (crw *customGinResponseWriter) format(src string) response.Response {
	var result response.Response
	err := json.Unmarshal([]byte(src), &result)
	if err != nil {
		return response.Response{}
	}
	return result
}
