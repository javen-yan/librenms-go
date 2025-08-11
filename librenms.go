// LibreNMS API Documentation: https://docs.librenms.org/API/
package librenms

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	apiVersion = "v0"
	authHeader = "X-Auth-Token"
)

// newRequest creates a new HTTP request with the given method and path.
// A relative URI should be provided and should not have a leading slash.
func (c *Client) newRequest(method, uri string, body any, query *url.Values) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(body); err != nil {
			return nil, err
		}
	}
	ctx := context.Background()

	// Parse the URI and construct the full URL
	fullURL, err := c.baseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, method, fullURL.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set necessary headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set(authHeader, c.token)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Add query parameters if provided
	if query != nil && len(*query) > 0 {
		req.URL.RawQuery = query.Encode()
	}

	c.log.LogAttrs(ctx, slog.LevelDebug, "http request", logRequestAttr(req))
	return req, nil
}

// rawDo sends an HTTP request and returns the raw response body. We should normally
// use do() which JSON-decodes and closes the response body, but if there is a non-JSON
// endpoint or other reason to not decode, this can be used.
func (c *Client) rawDo(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	c.log.LogAttrs(context.Background(), slog.LevelDebug, "http response", logResponseAttr(resp))
	return resp, checkResponse(resp)
}

// do sends an HTTP request and decodes the JSON response into the provided response object.
func (c *Client) do(req *http.Request, respObj any) error {
	if respObj == nil {
		return errors.New("response object cannot be nil")
	}

	resp, err := c.rawDo(req)
	if err != nil {
		return err
	}
	defer closeBody(resp.Body)

	switch v := respObj.(type) {
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if errors.Is(decErr, io.EOF) {
			decErr = nil // No content to decode, treat as success
		}
		if decErr != nil {
			err = fmt.Errorf("failure decoding response: %w", decErr)
		}
	}
	return err
}

// checkResponse checks the HTTP response for errors.
func checkResponse(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errorResponse := &ErrorResponse{
			Response: resp,
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorResponse.Message = fmt.Sprintf("failed to read response body: %v", err)
			defer closeBody(resp.Body)
			return errorResponse
		}

		if len(body) > 0 {
			if err = json.Unmarshal(body, errorResponse); err != nil {
				errorResponse.Message = string(body)
			}
		}

		return errorResponse
	}
	return nil
}

func closeBody(body io.ReadCloser) {
	_ = body.Close()
}

// parseParams is a helper function that parses the provided value into URL query parameters.
func parseParams(v any) (*url.Values, error) {
	if v == nil {
		return new(url.Values), nil
	}

	p, err := query.Values(v)
	if err != nil {
		return nil, fmt.Errorf("failed to parse query parameters: %w", err)
	}
	return &p, nil
}
