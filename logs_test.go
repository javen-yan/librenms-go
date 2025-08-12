package librenms

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/javen-yan/librenms-go/types"
)

func TestLogsAPI_ListEventLogs(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for event logs
		if r.URL.Path != "/api/v0/logs/eventlog/testdevice" {
			t.Errorf("Expected path /api/v0/logs/eventlog/testdevice, got %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("Expected method GET, got %s", r.Method)
		}

		// Return a mock response
		response := `{
			"status": "ok",
			"message": "",
			"count": 2,
			"total": "15",
			"logs": [
				{
					"hostname": "testdevice",
					"sysName": "testdevice.local",
					"event_id": "10050349",
					"host": "279",
					"device_id": "279",
					"datetime": "2017-07-22 19:57:47",
					"message": "ifAlias:  ->  <pptp-something-something-tunnel-something>",
					"type": "interface",
					"reference": "NULL",
					"username": "",
					"severity": "3"
				},
				{
					"hostname": "testdevice",
					"sysName": "testdevice.local",
					"event_id": "10050353",
					"host": "279",
					"device_id": "279",
					"datetime": "2017-07-22 19:57:47",
					"message": "ifHighSpeed:  ->  0",
					"type": "interface",
					"reference": "NULL",
					"username": "",
					"severity": "3"
				}
			]
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create a client
	client, err := NewClient(server.URL, "testtoken")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test query parameters
	query := &types.LogsQuery{
		Start:     1,
		Limit:     20,
		From:      "2017-07-22 23:00:00",
		SortOrder: "DESC",
	}

	// Call the API
	response, err := client.Logs.ListEventLogs("testdevice", query)
	if err != nil {
		t.Fatalf("Failed to list event logs: %v", err)
	}

	// Verify the response
	if response.Status != "ok" {
		t.Errorf("Expected status 'ok', got '%s'", response.Status)
	}
	if response.Count != 2 {
		t.Errorf("Expected count 2, got %d", response.Count)
	}
	if response.Total != 15 {
		t.Errorf("Expected total '15', got '%d'", response.Total)
	}
	if len(response.Logs) != 2 {
		t.Errorf("Expected 2 logs, got %d", len(response.Logs))
	}

	// Verify the first log entry
	log := response.Logs[0]
	if log.Hostname != "testdevice" {
		t.Errorf("Expected hostname 'testdevice', got '%s'", log.Hostname)
	}
	if log.EventID != 10050349 {
		t.Errorf("Expected event_id '10050349', got '%d'", log.EventID)
	}
	if log.Message != "ifAlias:  ->  <pptp-something-something-tunnel-something>" {
		t.Errorf("Expected message 'ifAlias:  ->  <pptp-something-something-tunnel-something>', got '%s'", log.Message)
	}
}

func TestLogsAPI_ListSysLogs(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for sys logs
		if r.URL.Path != "/api/v0/logs/syslog/testdevice" {
			t.Errorf("Expected path /api/v0/logs/syslog/testdevice, got %s", r.URL.Path)
		}

		// Return a mock response
		response := `{
			"status": "ok",
			"message": "",
			"count": 1,
			"total": "5",
			"logs": [
				{
					"hostname": "testdevice",
					"sysName": "testdevice.local",
					"event_id": "10050350",
					"host": "279",
					"device_id": "279",
					"datetime": "2017-07-22 19:58:00",
					"message": "System log message",
					"type": "system",
					"reference": "NULL",
					"username": "",
					"severity": "5"
				}
			]
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create a client
	client, err := NewClient(server.URL, "testtoken")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Call the API
	response, err := client.Logs.ListSysLogs("testdevice", nil)
	if err != nil {
		t.Fatalf("Failed to list sys logs: %v", err)
	}

	// Verify the response
	if response.Status != "ok" {
		t.Errorf("Expected status 'ok', got '%s'", response.Status)
	}
	if response.Count != 1 {
		t.Errorf("Expected count 1, got %d", response.Count)
	}
}

func TestLogsAPI_ListAlertLogs(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for alert logs
		if r.URL.Path != "/api/v0/logs/alertlog/testdevice" {
			t.Errorf("Expected path /api/v0/logs/alertlog/testdevice, got %s", r.URL.Path)
		}

		// Return a mock response
		response := `{
			"status": "ok",
			"message": "",
			"count": 1,
			"total": "3",
			"logs": [
				{
					"hostname": "testdevice",
					"sysName": "testdevice.local",
					"event_id": "10050351",
					"host": "279",
					"device_id": "279",
					"datetime": "2017-07-22 19:59:00",
					"message": "Alert log message",
					"type": "alert",
					"reference": "NULL",
					"username": "",
					"severity": "1"
				}
			]
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create a client
	client, err := NewClient(server.URL, "testtoken")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Call the API
	response, err := client.Logs.ListAlertLogs("testdevice", nil)
	if err != nil {
		t.Fatalf("Failed to list alert logs: %v", err)
	}

	// Verify the response
	if response.Status != "ok" {
		t.Errorf("Expected status 'ok', got '%s'", response.Status)
	}
	if response.Count != 1 {
		t.Errorf("Expected count 1, got %d", response.Count)
	}
}

func TestLogsAPI_ListAuthLogs(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for auth logs
		if r.URL.Path != "/api/v0/logs/authlog/testdevice" {
			t.Errorf("Expected path /api/v0/logs/authlog/testdevice, got %s", r.URL.Path)
		}

		// Return a mock response
		response := `{
			"status": "ok",
			"message": "",
			"count": 1,
			"total": "2",
			"logs": [
				{
					"hostname": "testdevice",
					"sysName": "testdevice.local",
					"event_id": "10050352",
					"host": "279",
					"device_id": "279",
					"datetime": "2017-07-22 20:00:00",
					"message": "Authentication log message",
					"type": "auth",
					"reference": "NULL",
					"username": "admin",
					"severity": "4"
				}
			]
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create a client
	client, err := NewClient(server.URL, "testtoken")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Call the API
	response, err := client.Logs.ListAuthLogs("testdevice", nil)
	if err != nil {
		t.Fatalf("Failed to list auth logs: %v", err)
	}

	// Verify the response
	if response.Status != "ok" {
		t.Errorf("Expected status 'ok', got '%s'", response.Status)
	}
	if response.Count != 1 {
		t.Errorf("Expected count 1, got %d", response.Count)
	}
}

func TestLogsAPI_ListLogs(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for event logs (alias)
		if r.URL.Path != "/api/v0/logs/eventlog/testdevice" {
			t.Errorf("Expected path /api/v0/logs/eventlog/testdevice, got %s", r.URL.Path)
		}

		// Return a mock response
		response := `{
			"status": "ok",
			"message": "",
			"count": 1,
			"total": "10",
			"logs": [
				{
					"hostname": "testdevice",
					"sysName": "testdevice.local",
					"event_id": "10050354",
					"host": "279",
					"device_id": "279",
					"datetime": "2017-07-22 20:01:00",
					"message": "Generic log message",
					"type": "generic",
					"reference": "NULL",
					"username": "",
					"severity": "6"
				}
			]
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create a client
	client, err := NewClient(server.URL, "testtoken")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Call the API (alias method)
	response, err := client.Logs.ListLogs("testdevice", nil)
	if err != nil {
		t.Fatalf("Failed to list logs: %v", err)
	}

	// Verify the response
	if response.Status != "ok" {
		t.Errorf("Expected status 'ok', got '%s'", response.Status)
	}
	if response.Count != 1 {
		t.Errorf("Expected count 1, got %d", response.Count)
	}
}

func TestLogsAPI_Syslogsink(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for syslogsink
		if r.URL.Path != "/api/v0/logs/syslogsink" {
			t.Errorf("Expected path /api/v0/logs/syslogsink, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, got %s", r.Method)
		}

		// Return a mock response
		response := `{
			"status": "ok",
			"message": "Messages processed successfully",
			"count": 0
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create a client
	client, err := NewClient(server.URL, "testtoken")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Create test messages
	messages := types.SyslogsinkRequest{
		{
			Msg:  "kernel: minimum Message",
			Host: "mydevice.fqdn.com",
		},
		{
			Msg:       "Line protocol on Interface GigabitEthernet1/0/41, changed state to up",
			Facility:  23,
			Priority:  "189",
			Program:   "LINEPROTO-5-UPDOWN",
			Host:      "172.29.10.24",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Severity:  5,
			Level:     "ERROR",
		},
	}

	// Call the API
	response, err := client.Logs.Syslogsink(messages)
	if err != nil {
		t.Fatalf("Failed to send syslogsink: %v", err)
	}

	// Verify the response
	if response.Status != "ok" {
		t.Errorf("Expected status 'ok', got '%s'", response.Status)
	}
	if response.Message != "Messages processed successfully" {
		t.Errorf("Expected message 'Messages processed successfully', got '%s'", response.Message)
	}
}

// Helper function to create int pointers
func intPtr(i int) *int {
	return &i
}
