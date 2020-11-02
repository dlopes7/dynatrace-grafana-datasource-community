package dynatrace

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context/ctxhttp"
	"io"
	"net/http"
)

type HttpClient struct {
	Client http.Client
	Token  string
}

func (c *HttpClient) makeRequest(ctx context.Context, method string, url string, model interface{}, body interface{}) error {

	req, err := c.newRequest(method, url, body)
	if err != nil {
		return err
	}

	if model == nil {
		return c.do(ctx, *req, nil)
	} else {
		return c.do(ctx, *req, &model)
	}
}

func (c *HttpClient) do(ctx context.Context, req http.Request, model interface{}) error {

	resp, err := ctxhttp.Do(ctx, &c.Client, &req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error making request for %s, status code: %d", req.URL, resp.StatusCode)
	}

	if model != nil {
		err = json.NewDecoder(resp.Body).Decode(model)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *HttpClient) newRequest(method string, url string, body interface{}) (*http.Request, error) {

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Api-Token %s", c.Token))

	return req, nil
}
