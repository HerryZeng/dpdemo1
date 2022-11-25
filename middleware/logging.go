package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/HerryZeng/dpdemo1/MyLog"
	"github.com/HerryZeng/dpdemo1/model/log"
	"github.com/HerryZeng/dpdemo1/myerr"
	"github.com/HerryZeng/dpdemo1/res"
	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
)

func Logging() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now().UTC()
		path := context.Request.URL.Path

		reqUrl := context.Request.URL.Path
		if strings.Contains(reqUrl, "image") {
			//context.Next()
			return
		}

		var bodies []byte
		if context.Request.Body != nil {
			bodies, _ = ioutil.ReadAll(context.Request.Body)
		}
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodies))
		method := context.Request.Method
		ip := context.ClientIP()

		blw := &log.BodyLoggerWriter{
			ResponseWriter: context.Writer,
			Body:           bytes.NewBufferString(""),
		}
		context.Writer = blw

		context.Next()

		end := time.Now().UTC()
		sub := end.Sub(start)
		code := -100
		message := ""
		var response res.Response
		if err := json.Unmarshal(blw.Body.Bytes(), &response); err != nil {
			code = myerr.InternalServerError.Code
			message = err.Error()
			MyLog.Log.Error(err)
		} else {
			code = response.Code
			message = response.Message
		}
		MyLog.Log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", sub, ip, pad.Right(method, 5, ""), path, code, message)
	}
}
