package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"
)

var _ Client = client{}
var _ Request = request(func() *Response { return nil })

type (
	Client interface {
		Request(t *testing.T, method, url string, input ...RequestInput) Request
	}

	Request interface {
		To(interface{}) Request
		Run() *Response
	}

	RequestInput struct {
		body  interface{}
		query map[string]string
	}

	Response struct {
		test *testing.T
		data *http.Response
		dump string
	}

	client struct {
		client *http.Client
		origin *url.URL
	}

	request func() *Response
)

func (r request) To(placeholder interface{}) Request {
	return request(func() *Response {
		res := r()

		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.data.Body)

		err := json.NewDecoder(res.data.Body).Decode(placeholder)
		if err != nil {
			fmt.Println(res.dump)
			res.test.Fatalf("Could not parse JSON response to the struct provided: %s", err)
		}

		return res
	})
}

func (r request) Run() *Response {
	return r()
}

func (c client) Request(t *testing.T, method, path string, inputs ...RequestInput) Request {
	var in RequestInput
	if len(inputs) > 0 {
		in = inputs[0]
	}

	return request(func() *Response {
		var body []byte
		if in.body != nil {
			b, err := json.Marshal(in.body)
			if err != nil {
				t.Fatalf("Could not parse body of struct %v: %s", in.body, err)
			}

			body = b
		}

		req, err := http.NewRequest(method, c.origin.String()+path, bytes.NewReader(body))
		if err != nil {
			t.Fatalf("Could not mount of request to %v: %s", path, err)
		}

		var query url.Values
		if in.query != nil {
			query = url.Values{}
			for key, val := range in.query {
				query.Set(key, val)
			}
			req.URL.RawQuery = query.Encode()
		}

		return c.DoRequest(t, req)
	})
}

func (c client) DoRequest(t *testing.T, req *http.Request) *Response {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not receive response: %s", err)
	}

	var dump string
	body, err := httputil.DumpResponse(res, true)
	if err == nil {
		dump = string(body)
	}

	return &Response{
		test: t,
		data: res,
		dump: dump,
	}
}
