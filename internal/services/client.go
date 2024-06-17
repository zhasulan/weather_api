package services

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"
	"weather_api/internal/logger"
	"weather_api/internal/utils"
)

type Client struct {
	client  *http.Client
	hostUrl string
	headers map[string]string
}

func NewClient(host string, headers map[string]string, timeout int) *Client {
	return &Client{
		client: &http.Client{Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			MaxIdleConns:       64,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: false,
			DisableKeepAlives:  false,
		},
			Timeout: time.Duration(timeout) * time.Second,
		},
		hostUrl: host,
		headers: headers,
	}
}

func (sc Client) ExecuteRequest(ctx context.Context, methodType int, urlParams string, requestData, responseBody interface{}, trace bool) (responseBodyBytes []byte, err error) {
	var start = time.Now()
	var httpStatusCode int = http.StatusOK
	var baseURL string = fmt.Sprintf("%s%s", sc.hostUrl, methodMap[methodType].endpoint)
	var requestURL string = fmt.Sprintf("%s?%s", baseURL, urlParams)

	var requestDataBytes []byte

	traceID := ctx.Value("TraceID")

	defer func() {
		if e := recover(); e != nil {
			err = utils.AnyError(e)
		}

		lightLog := logger.WeatherApiLogger.
			WithField(logger.LF_requestIdStr, traceID).
			WithField(logger.LF_operationStr, methodMap[methodType].name).
			WithField(logger.LF_serviceHttpStatusCode, httpStatusCode).
			WithField(logger.LF_durationSecFloat, time.Since(start).Seconds()).
			WithField(logger.LF_serviceResponseSizeInt, len(responseBodyBytes)).
			WithField(logger.LF_urlRequestStr, baseURL)

		// if request payload is absent, request URL is used as request body
		if requestDataBytes == nil {
			requestDataBytes = []byte(requestURL)
		}

		//maskedResponse, _ := logger.Masker.MaskBytes(responseBodyBytes)
		traceLog := lightLog.
			WithField(logger.LF_requestStr, requestDataBytes).
			WithField(logger.LF_responseStr, string(responseBodyBytes))

		// error handling
		if err != nil {
			traceLog.
				WithError(err).
				WithField(logger.LF_stackStr, logger.WithStack(err)).
				Error("error on service request")
			return
		}

		if httpStatusCode >= http.StatusInternalServerError {
			traceLog.Error("service request internal error")

			return
		}

		if httpStatusCode >= http.StatusBadRequest {
			traceLog.Warn("service request bad request")

			return
		}

		if trace {
			traceLog.Info("service trace")

			return
		}

		lightLog.Info("service log")
	}()

	if requestData != nil {
		requestDataBytes, err = json.Marshal(requestData)
		if err != nil {
			return
		}
	}

	request, err := http.NewRequestWithContext(ctx, methodMap[methodType].method, requestURL, bytes.NewReader(requestDataBytes))
	if err != nil {
		return
	}

	for key, value := range sc.headers {
		request.Header.Set(key, value)
	}

	response, err := sc.client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.Body != nil {
		responseBodyBytes, err = io.ReadAll(response.Body)
		if err != nil {
			return
		}
	}

	httpStatusCode = response.StatusCode

	if response.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("incorrect status code: %d", response.StatusCode))

		return
	}

	err = json.Unmarshal(responseBodyBytes, &responseBody)

	return
}
