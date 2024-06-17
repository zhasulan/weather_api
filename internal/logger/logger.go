package logger

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"time"
	"weather_api/internal/config"
	"weather_api/internal/utils"
	"weather_api/meta"
)

var WeatherApiLogger *logger.Entry

func Logger(inner http.Handler, name string, trace bool) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		ctx := request.Context()
		traceID := uuid.New().String()

		//response recorders keeps
		requestDump, _ := httputil.DumpRequest(request, true)
		rrw := newResponseRecorder(writer)
		ctx = context.WithValue(ctx, "TraceID", traceID)
		inner.ServeHTTP(rrw, request.WithContext(ctx))

		var response = ""
		if rrw.Header().Get("Content-Encoding") == "gzip" {
			response = rrw.GUnZipBody()
		} else {
			response = rrw.Body()
		}

		responseStatusCode := rrw.Status

		log := WeatherApiLogger.
			WithField(LF_requestStr, string(requestDump)).
			WithField(LF_responseStr, response).
			WithField(LF_methodNameStr, name).
			WithField(LF_serviceResponseSizeInt, len(response)).
			WithField(LF_serviceHttpStatusCode, responseStatusCode).
			WithField(LF_durationSecFloat, time.Since(start).Seconds()).
			WithField(LF_requestIdStr, traceID)

		if responseStatusCode >= http.StatusInternalServerError {
			log.Error("Server error")
			return
		}

		if responseStatusCode >= http.StatusBadRequest {
			log.Warn("Bad request")
			return
		}

		if trace {
			log.Info("trace")
		}
	})
}

func InitLogger() {
	//keyValueMap := embed.LoadKV()
	//
	//err := embed.LoadTomlConfig("app.toml", keyValueMap, &model.AppConfig)
	//if err != nil {
	//	_, _ = fmt.Fprintf(os.Stderr, "Missing or corrupted 'app.toml' file")
	//	panic(err)
	//}
	//
	logrusl := logger.New()
	//Comment this to see plain text logs
	logrusl.SetFormatter(&logger.JSONFormatter{
		FieldMap: logger.FieldMap{
			logger.FieldKeyMsg: "message",
		},
	})

	l, err := logger.ParseLevel(config.Config.App.LogLevel)
	if err != nil {
		l = logger.InfoLevel
	}
	logrusl.SetLevel(l)

	//logPath := model.AppConfig.Log.Path
	//if logPath == "" {
	//	logPath = "atlas.log"
	//} else {
	//	if logPath[len(logPath)-1] == '/' {
	//		logPath = logPath + "atlas.log"
	//	} else {
	//		logPath = logPath + "/" + "atlas.log"
	//	}
	//}
	//
	////Comment if statement to see logs in console or change app.toml
	//if model.AppConfig.Log.Out == "file" {
	//	logrusl.SetOutput(&lumberjack.Logger{
	//		Filename: logPath,
	//		MaxSize:  10,   // megabytes
	//		MaxAge:   1,    //days
	//		Compress: true, // disabled by default
	//	})
	//}

	WeatherApiLogger = logrusl.WithFields(logger.Fields{
		"service.name": config.Config.App.Name,
		"service.version": func() string {
			if meta.GitBranch != "" { // codebase, using go linker
				return meta.GitBranch + "_" + meta.GitHash
			}
			return config.Config.App.Version // from config
		}(),
	})
}

func WithStack(a any) string {
	if a == nil {
		return ""
	}

	err := utils.AnyError(a)
	return fmt.Sprintf("%+v", errors.WithStack(err))
}

const (
	LF_serviceResponseSizeInt = "responseSize"
	LF_serviceHttpStatusCode  = "httpStatusCode"
	LF_operationStr           = "operation"
	LF_errorErr               = "error"
	LF_stackStr               = "stack"
	LF_urlRequestStr          = "urlRequest"
	LF_methodNameStr          = "method"
	LF_durationSecFloat       = "durationSec"
	LF_requestStr             = "request"
	LF_responseStr            = "response"
	LF_requestIdStr           = "traceID"
)
