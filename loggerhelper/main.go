package loggerhelper

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetLevel(5)
	setupOutput(l)
	return l
}

type WriterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func setupOutput(l *logrus.Logger) {
	l.SetOutput(ioutil.Discard) // Send all logs to nowhere by default

	infoFile, err := os.OpenFile("log_info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.Fatal("Failed to log to file, using default stderr")
	}
	errorFile, err := os.OpenFile("log_error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.Fatal("Failed to log to file, using default stderr")
	}

	l.AddHook(&WriterHook{ // Send info and debug logs to stdout
		Writer:    infoFile,
		LogLevels: []logrus.Level{logrus.InfoLevel, logrus.DebugLevel, logrus.WarnLevel, logrus.TraceLevel, logrus.DebugLevel},
	})

	l.AddHook(&WriterHook{ // Send logs with level higher than warning to stderr
		Writer:    errorFile,
		LogLevels: []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel},
	})
}

var timeFormat = "02-01-2006:15:04:05"

func LoggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Starting time
		startTime := time.Now()

		ctx.Next()

		// End Time
		endTime := time.Now()

		// execution time
		latencyTime := endTime.Sub(startTime)

		// Request method
		reqMethod := ctx.Request.Method

		// Request route
		path := ctx.Request.URL.Path

		// status code
		statusCode := ctx.Writer.Status()

		// Request IP
		clientIP := ctx.ClientIP()

		dataLength := ctx.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		clientUserAgent := ctx.Request.UserAgent()
		referer := ctx.Request.Referer()

		entry := logger.WithFields(logrus.Fields{
			"METHOD":            reqMethod,
			"PATH":              path,
			"STATUS":            statusCode,
			"LATENCY":           latencyTime,
			"CLIENT_IP":         clientIP,
			"CLIENT_USER_AGENT": clientUserAgent,
		})

		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%s - %s [%s] \"%s \" %d %d \"%s\" \"%s\" (%dms) %s", clientIP, startTime.Format(timeFormat), reqMethod, path, statusCode, dataLength, referer, clientUserAgent, latencyTime, ctx.Errors.String())
			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}

		ctx.Next()
	}
}
