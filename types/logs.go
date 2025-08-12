package types

type (
	// Log represents a log entry in LibreNMS.
	Log struct {
		Hostname  string `json:"hostname,omitempty"`
		SysName   string `json:"sysName,omitempty"`
		EventID   string `json:"event_id,omitempty"`
		Host      string `json:"host,omitempty"`
		DeviceID  string `json:"device_id,omitempty"`
		DateTime  string `json:"datetime,omitempty"`
		Message   string `json:"message,omitempty"`
		Type      string `json:"type,omitempty"`
		Reference string `json:"reference,omitempty"`
		Username  string `json:"username,omitempty"`
		Severity  string `json:"severity,omitempty"`
	}

	// LogsResponse represents a response containing logs from the LibreNMS API.
	LogsResponse struct {
		BaseResponse
		Total string `json:"total,omitempty"`
		Logs  []Log  `json:"logs"`
	}

	// LogsQuery represents the query parameters for filtering logs.
	LogsQuery struct {
		Start     int    `url:"start,omitempty"`     // The page number to request
		Limit     int    `url:"limit,omitempty"`     // The limit of results to be returned
		From      string `url:"from,omitempty"`      // The date and time or the event id to search from
		To        string `url:"to,omitempty"`        // The date and time or the event id to search to
		SortOrder string `url:"sortorder,omitempty"` // Sort order (ASC/DESC)
	}

	// SyslogMessage represents a single syslog message for the syslogsink endpoint.
	SyslogMessage struct {
		Msg       string `json:"msg"`                  // The log message
		Host      string `json:"host"`                 // The hostname or IP address
		Facility  int    `json:"facility,omitempty"`   // Syslog facility (optional)
		Priority  string `json:"priority,omitempty"`   // Syslog priority (optional)
		Program   string `json:"program,omitempty"`    // Program name (optional)
		Timestamp string `json:"@timestamp,omitempty"` // ISO timestamp (optional)
		Severity  int    `json:"severity,omitempty"`   // Severity level (optional)
		Level     string `json:"level,omitempty"`      // Log level (optional)
	}

	// SyslogsinkRequest represents the request body for the syslogsink endpoint.
	// It can be a single message or an array of messages.
	SyslogsinkRequest []SyslogMessage
)
