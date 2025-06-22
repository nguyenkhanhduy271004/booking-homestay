package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func LoggerMiddleware() gin.HandlerFunc {
	logPath := "logs/http.log"

	logger := zerolog.New(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     5,
		Compress:   true,
		LocalTime:  true,
	}).With().Timestamp().Logger()

	return func(ctx *gin.Context) {
		startTime := time.Now()
		var requestBody string

		contentType := ctx.GetHeader("Content-Type")

		if ctx.Request.Body != nil && (ctx.Request.Method == http.MethodPost || ctx.Request.Method == http.MethodPut || ctx.Request.Method == http.MethodPatch) {
			bodyBytes, _ := io.ReadAll(ctx.Request.Body)
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			switch {
			case strings.HasPrefix(contentType, "application/json"):
				requestBody = string(bodyBytes)

			case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
				values, err := url.ParseQuery(string(bodyBytes))
				if err == nil {
					formData := make(map[string]string)
					for k, v := range values {
						formData[k] = strings.Join(v, ",")
					}
					formEncoded, _ := json.Marshal(formData)
					requestBody = string(formEncoded)
				} else {
					requestBody = string(bodyBytes)
				}

			case strings.HasPrefix(contentType, "multipart/form-data"):
				requestBody = "[multipart form-data omitted]"

			default:
				requestBody = string(bodyBytes)
			}
		}

		ctx.Next()

		latency := time.Since(startTime)
		statusCode := ctx.Writer.Status()

		logEvent := logger.Info()
		if statusCode >= 500 {
			logEvent = logger.Error()
		} else if statusCode >= 400 {
			logEvent = logger.Warn()
		}

		logEvent.
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("query", ctx.Request.URL.RawQuery).
			Str("ip", ctx.ClientIP()).
			Str("user_agent", ctx.Request.UserAgent()).
			Str("referer", ctx.Request.Referer()).
			Str("protocol", ctx.Request.Proto).
			Str("host", ctx.Request.Host).
			Str("remote_addr", ctx.Request.RemoteAddr).
			Str("request_uri", ctx.Request.RequestURI).
			Interface("headers", ctx.Request.Header).
			Str("body", requestBody).
			Int("status_code", statusCode).
			Int("response_size", ctx.Writer.Size()).
			Dur("latency", latency).
			Str("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()).
			Msg("HTTP Request")
	}
}
