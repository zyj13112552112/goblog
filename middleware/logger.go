package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

// Logger logrus 日志
func Logger() gin.HandlerFunc {
	logPath := "log/log.log"
	src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}

	logger := logrus.New()
	logger.Out = src

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statuscode := c.Writer.Status()
		ip := c.ClientIP()
		agent := c.Request.UserAgent()
		datasize := c.Writer.Size()
		if datasize < 0 {
			datasize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statuscode,
			"SpendTime": spendTime,
			"IP":        ip,
			"Method":    method,
			"Path":      path,
			"DataSize":  datasize,
			"Agent":     agent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statuscode >= 500 {
			entry.Error()
		} else if statuscode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
