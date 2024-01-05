package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/safciplak/capila/src/logger"
)

// LogMiddleware Logs request and responses
func LogMiddleware(logProvider logger.InterfaceLogger) gin.HandlerFunc {
	return func(context *gin.Context) {
		var buffer bytes.Buffer

		// Creates a reader which keeps the original buffer intact
		body, err := io.ReadAll(io.TeeReader(context.Request.Body, &buffer))
		if err != nil {
			panic("unable to create a buffer to logProvider the request")
		}

		// Resets the buffer to the original state
		context.Request.Body = io.NopCloser(&buffer)

		logWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: context.Writer}
		context.Writer = logWriter

		ctx := context.Request.Context()

		log := logProvider.Log(ctx)
		// Logs the request
		log.Info("request",
			zap.String("url", context.Request.URL.String()),
			zap.String("method", context.Request.Method),
			zap.Strings("headers", headerToStringArray(context.Request.Header)),
			zap.String("body", string(body)),
		)

		context.Next()

		// Logs the response
		log.Info("response",
			zap.Int("status", logWriter.Status()),
			zap.Strings("headers", headerToStringArray(logWriter.Header())),
			zap.String("body", logWriter.body.String()),
		)
	}
}

// bodyLogWriter is used to write the response body to a separate buffer
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write overrides the Write function of the ResponseWriter and also writes to the internal buffer
func (w bodyLogWriter) Write(b []byte) (n int, err error) {
	n, err = w.body.Write(b)
	if err != nil {
		return
	}

	return w.ResponseWriter.Write(b)
}

// WriteString overrides the WriteString function of the ResponseWriter and also writes to the internal buffer
func (w bodyLogWriter) WriteString(s string) (n int, err error) {
	n, err = w.body.WriteString(s)
	if err != nil {
		return
	}

	return w.ResponseWriter.WriteString(s)
}

// headerToStringArray Converts the go Header string map of string arrays into a flattened string array
func headerToStringArray(headerMap http.Header) []string {
	fields := make([]string, 0, len(headerMap))

	for headerIndex, headerValue := range headerMap {
		field := headerIndex + ":"
		for _, value := range headerValue {
			field = field + " " + value
		}

		fields = append(fields, field)
	}

	// Order of headers is randomized, thus we sort them alphabetically
	sort.Strings(fields)

	return fields
}
