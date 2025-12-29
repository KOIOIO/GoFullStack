// middleware/cors.go
package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

// CorsMiddleware 全局跨域中间件
func CorsMiddleware() rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 设置跨域头（在调用前确保存在）
			setCorsHeaders := func(h http.Header) {
				// 使用通配符 Origin，若需要支持凭证请改为具体 Origin 并移除 '*' 或设置为请求 Origin
				h.Set("Access-Control-Allow-Origin", "*")
				h.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				h.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				h.Set("Access-Control-Allow-Credentials", "true")
			}

			setCorsHeaders(w.Header())

			// 如果是OPTIONS请求，直接返回204 No Content
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			// 继续处理请求，并在处理后再次确保 CORS 头存在（以防被覆盖）
			next(w, r)
			setCorsHeaders(w.Header())
		}
	}
}
