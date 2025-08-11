package librenms

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"log/slog"

	"github.com/hashicorp/go-cleanhttp"
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
	c.baseURL, err = c.baseURL.Parse("api/v0/")
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
