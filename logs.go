package librenms

import (
	"fmt"
	"net/http"

	"github.com/javen-yan/librenms-go/types"
)

const (
	// logsEndpoint is the base API endpoint for logs.
	logsEndpoint = "logs"
)

// LogsAPI provides log-related operations
type LogsAPI struct {
	client *Client
}

// ListEventLogs retrieves event logs for a specific device.
// The identifier can be either a device ID or hostname.
func (l *LogsAPI) ListEventLogs(identifier string, query *types.LogsQuery) (*types.LogsResponse, error) {
	uri := fmt.Sprintf("%s/eventlog/%s", logsEndpoint, identifier)
	return l.listLogs(uri, query)
}

// ListSysLogs retrieves system logs for a specific device.
// The identifier can be either a device ID or hostname.
func (l *LogsAPI) ListSysLogs(identifier string, query *types.LogsQuery) (*types.LogsResponse, error) {
	uri := fmt.Sprintf("%s/syslog/%s", logsEndpoint, identifier)
	return l.listLogs(uri, query)
}

// ListAlertLogs retrieves alert logs for a specific device.
// The identifier can be either a device ID or hostname.
func (l *LogsAPI) ListAlertLogs(identifier string, query *types.LogsQuery) (*types.LogsResponse, error) {
	uri := fmt.Sprintf("%s/alertlog/%s", logsEndpoint, identifier)
	return l.listLogs(uri, query)
}

// ListAuthLogs retrieves authentication logs for a specific device.
// The identifier can be either a device ID or hostname.
func (l *LogsAPI) ListAuthLogs(identifier string, query *types.LogsQuery) (*types.LogsResponse, error) {
	uri := fmt.Sprintf("%s/authlog/%s", logsEndpoint, identifier)
	return l.listLogs(uri, query)
}

// ListLogs is an alias for ListEventLogs for backward compatibility.
// All list_*logs calls are aliased to list_logs in the LibreNMS API.
func (l *LogsAPI) ListLogs(identifier string, query *types.LogsQuery) (*types.LogsResponse, error) {
	return l.ListEventLogs(identifier, query)
}

// Syslogsink sends syslog messages to the LibreNMS syslog sink endpoint.
// This endpoint accepts any JSON messages and passes them to further syslog processing.
// It can handle single messages or an array of multiple messages.
func (l *LogsAPI) Syslogsink(messages types.SyslogsinkRequest) (*types.BaseResponse, error) {
	uri := fmt.Sprintf("%s/syslogsink", logsEndpoint)

	req, err := l.client.newRequest(http.MethodPost, uri, messages, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create syslogsink request: %w", err)
	}

	var response types.BaseResponse
	if err := l.client.do(req, &response); err != nil {
		return nil, fmt.Errorf("failed to send syslogsink request: %w", err)
	}

	return &response, nil
}

// listLogs is a helper method that handles the common logic for listing logs.
func (l *LogsAPI) listLogs(uri string, query *types.LogsQuery) (*types.LogsResponse, error) {
	params, err := parseParams(query)
	if err != nil {
		return nil, fmt.Errorf("failed to parse query parameters: %w", err)
	}

	req, err := l.client.newRequest(http.MethodGet, uri, nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create logs request: %w", err)
	}

	var response types.LogsResponse
	if err := l.client.do(req, &response); err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}

	return &response, nil
}
