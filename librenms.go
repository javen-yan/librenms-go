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
	"os"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-cleanhttp"
)

const (
	apiVersion = "v0"
	authHeader = "X-Auth-Token"
)

// Client represents the main LibreNMS client
type Client struct {
	baseURL *url.URL
	client  *http.Client
	log     *slog.Logger
	token   string

	// API interfaces
	Device      *DeviceAPI
	Alert       *AlertAPI
	AlertRule   *AlertRuleAPI
	DeviceGroup *DeviceGroupAPI
	Location    *LocationAPI
	Service     *ServiceAPI
	System      *SystemAPI
	Port        *PortAPI
	Inventory   *InventoryAPI
	Routing     *RoutingAPI
	Switching   *SwitchingAPI
	Logs        *LogsAPI
}

// DeviceAPI provides device-related operations
type DeviceAPI struct {
	client *Client
}

// PortAPI provides port-related operations
type PortAPI struct {
	client *Client
}

// InventoryAPI provides inventory-related operations
type InventoryAPI struct {
	client *Client
}

// AlertAPI provides alert-related operations
type AlertAPI struct {
	client *Client
}

// AlertRuleAPI provides alert rule-related operations
type AlertRuleAPI struct {
	client *Client
}

// DeviceGroupAPI provides device group-related operations
type DeviceGroupAPI struct {
	client *Client
}

// LocationAPI provides location-related operations
type LocationAPI struct {
	client *Client
}

// ServiceAPI provides service-related operations
type ServiceAPI struct {
	client *Client
}

// SystemAPI provides system-related operations
type SystemAPI struct {
	client *Client
}

// RoutingAPI provides routing-related operations
type RoutingAPI struct {
	client *Client
}

// SwitchingAPI provides switching-related operations
type SwitchingAPI struct {
	client *Client
}

// Option is a function that configures the Client
type Option func(*Client)

// WithHTTPClient sets the HTTP client for the LibreNMS client
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

// WithLogger sets a custom logger for the LibreNMS client
func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) {
		if logger == nil {
			panic("logger cannot be nil")
		}
		c.log = logger
	}
}

// WithLogLevel sets the logging level for the default client logger
func WithLogLevel(level slog.Level) Option {
	return func(c *Client) {
		c.log = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: level,
		}))
	}
}

// NewClient creates a new LibreNMS client with the given base URL and options
// The base URL should be in the format 'http[s]://<host>[:port]/'
func NewClient(baseURL, token string, opts ...Option) (*Client, error) {
	c := &Client{
		token:  token,
		client: cleanhttp.DefaultPooledClient(),
		log: slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})),
	}

	// Append a trailing slash to the base URL if it doesn't have one
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	// Parse the base URL
	var err error
	c.baseURL, err = url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	if c.baseURL.Path != "/" {
		return nil, fmt.Errorf("invalid base URL format, expected: 'http[s]://<host>[:port]/'")
	}

	// Append the API version to the base URL path
	c.baseURL, err = c.baseURL.Parse(fmt.Sprintf("api/%s/", apiVersion))
	if err != nil {
		return nil, fmt.Errorf("failed to parse API version in base URL: %w", err)
	}

	// Process options
	for _, opt := range opts {
		opt(c)
	}

	// Initialize API interfaces
	c.Device = &DeviceAPI{client: c}
	c.Alert = &AlertAPI{client: c}
	c.AlertRule = &AlertRuleAPI{client: c}
	c.DeviceGroup = &DeviceGroupAPI{client: c}
	c.Location = &LocationAPI{client: c}
	c.Service = &ServiceAPI{client: c}
	c.System = &SystemAPI{client: c}
	c.Port = &PortAPI{client: c}
	c.Inventory = &InventoryAPI{client: c}
	c.Routing = &RoutingAPI{client: c}
	c.Switching = &SwitchingAPI{client: c}
	c.Logs = &LogsAPI{client: c}

	return c, nil
}

// New is an alias for NewClient for backward compatibility
func New(baseURL, token string, opts ...Option) (*Client, error) {
	return NewClient(baseURL, token, opts...)
}

// GetBaseURL returns the base URL of the client
func (c *Client) GetBaseURL() string {
	return c.baseURL.String()
}

// GetToken returns the authentication token
func (c *Client) GetToken() string {
	return c.token
}

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
	// 去掉空值
	for k, v := range p {
		if len(v) == 0 || v[0] == "" {
			delete(p, k)
		}
	}
	return &p, nil
}
