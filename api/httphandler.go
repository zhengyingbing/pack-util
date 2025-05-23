package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	c      *Client
	header map[string]string
)

func init() {
	c = NewClient()
	header = make(map[string]string)
	header["Host"] = "center.hoolai.com"
	header["Access-Control-Allow-Origin"] = "*"
	header["X-TOKEN"] = "eyJob3BzTmFtZSI6InpoZW5neWluZ2JpbmciLCJob3BzUGFzc3dvcmQiOiIiLCJ0b2tlbiI6ImV5SmhiR2NpT2lKSVV6STFOaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpUYVdkdVNXNW1ieUk2ZXlKMWMyVnlYMmxrSWpveE1Dd2lkWE5sY2w5dVlXMWxJam9pZW1obGJtZDVhVzVuWW1sdVp5SXNJbTVoYldVaU9pTHBnNUhvdjQ3bGhiVWlMQ0p5YjJ4bFgybGtJam94T1gwc0ltVjRjQ0k2TVRjME56TTROelUwT1N3aWFXRjBJam94TnpRM016QXhNVFE1TENKcGMzTWlPaUpoYkNKOS44Z2xKZVJOVWRjUnFMc1RId1NSOFlrZ29hZHNCTlpnRFE5dkdHRmszM1IwIiwicmVtZW1iZXIiOmZhbHNlLCJqYXZhSG9tZSI6IkQ6XFxBbmRyb2lkXFxyZXNvdXJjZXNcXGphdmEiLCJhbmRyb2lkSG9tZSI6IkQ6XFxBbmRyb2lkXFxyZXNvdXJjZXNcXGFuZHJvaWQiLCJidWlsZERpciI6IiIsIm91dEFwa0RpciI6IiIsImFwa05hbWVGb3JtYXQiOiIiLCJwcmVQcm9kdWN0RW52IjoiIiwicHJveHlQb3J0IjoiODg4OCIsInByb3h5RGF0YVR5cGUiOlsic2RrIiwiYmkiLCJyZXl1biJdLCJwYWNrYWdlVW5pcXVlSWQiOiIxNzQ1MjIxODczMTcxIiwicGFja2FnZVJlZ2lvbiI6MX0="
}

func AddHeader(k, v string) {
	header[k] = v
}

type Request struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    interface{}
	Timeout time.Duration
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (resp Response) IsSuccess() bool {
	return resp.Code == 200
}
func Get(apiType ApiType) *Response {
	return c.Send(&Request{
		URL:     Host + string(apiType),
		Method:  http.MethodGet,
		Headers: header,
	})
}

func Post(apiType ApiType, body map[string]interface{}) *Response {
	return c.Send(&Request{
		URL:     Host + string(apiType),
		Method:  http.MethodPost,
		Headers: header,
		Body:    body,
	})
}

func (c *Client) Send(req *Request) *Response {
	if req.Timeout > 0 {
		c.httpClient.Timeout = req.Timeout
	}
	log.Println("1")
	var body io.Reader
	if req.Body != nil {
		data, err := json.Marshal(req.Body)
		if err != nil {
			return nil
		}
		body = bytes.NewBuffer(data)
	}
	log.Println("2")
	httpReq, err := http.NewRequest(req.Method, req.URL, body)
	if err != nil {
		return nil
	}
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}
	if req.Body != nil && httpReq.Header.Get("Content-Type") == "" {
		httpReq.Header.Set("Content-Type", "application/json")
	}
	log.Println("3")
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil
	}
	resp := &Response{
		Code: httpResp.StatusCode,
		Data: respBody,
	}
	log.Println("4")
	return resp
}
