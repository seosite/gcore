package netx

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/seosite/gcore/pkg/core/jsonx"
)

var (
	errStatusCodeNot200 = errors.New("Status code != 200")
)

var (
	maxWaitTime = time.Second * 15
	retryMax    = 3
)

// RetryClient retriable client
type RetryClient struct {
	RetryMax     int
	RetryWaitMax time.Duration
	client       *http.Client
}

// NewRetryClient new retry client
func NewRetryClient() *RetryClient {
	client := retryablehttp.NewClient()
	client.RetryMax = retryMax
	client.RetryWaitMax = maxWaitTime
	return &RetryClient{
		RetryMax:     retryMax,
		RetryWaitMax: maxWaitTime,
		client:       client.StandardClient(),
	}
}

// Get .
func (c *RetryClient) Get(url string, cxt *gin.Context) (string, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errStatusCodeNot200
	}

	return string(body), nil
}

// Post .
func (c *RetryClient) Post(url, contentType string, bodyContent io.Reader) (string, error) {
	resp, err := c.client.Post(url, contentType, bodyContent)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errStatusCodeNot200
	}

	return string(body), nil
}

// PostJSON .
func (c *RetryClient) PostJSON(url string, params map[string]interface{}) (string, error) {
	jsonContent, err := jsonx.Marshal(params)
	if err != nil {
		return "", err
	}

	return c.Post(url, "application/Json", bytes.NewBuffer(jsonContent))
}

// PostForm .
func (c *RetryClient) PostForm(url string, bodyContent url.Values) (string, error) {
	resp, err := c.client.PostForm(url, bodyContent)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errStatusCodeNot200
	}

	return string(body), nil
}
