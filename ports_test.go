package librenms_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testEndpointPorts = "/api/v0/ports"
)

// This init function will register handlers for port-related API endpoints.
// This is used when running the full test suite.
func init() {
	handleEndpoint(testEndpointPorts, mockResponses{
		http.MethodGet: loadMockResponse("get_ports_200.json"),
	})
}

// TestClient_GetPorts tests the GetPorts functionality
func TestClient_GetPorts(t *testing.T) {
	r := require.New(t)

	r.NotNil(testAPIClient, "Global testAPIClient should be initialized")

	portsResp, err := testAPIClient.Port.GetAllPorts(nil)

	r.NoError(err, "GetPorts returned an error")
	r.NotNil(portsResp, "GetPorts response is nil")

	r.Equal("ok", portsResp.Status, "Expected status 'ok'")
	r.Equal(1, portsResp.Count, "Expected count 1")
	r.Len(portsResp.Ports, 1, "Expected 1 port")

	port := portsResp.Ports[0]
	r.Equal("1", port.PortID, "Expected Port ID 1")
	r.Equal("1", port.DeviceID, "Expected DeviceID 1")
	r.Equal("1", port.IfIndex, "Expected IfIndex 1")
	r.Equal("GigabitEthernet1/0/1", port.IfName, "Expected IfName 'GigabitEthernet1/0/1'")
	r.Equal("Uplink to Core Switch", port.IfDescr, "Expected IfDescr 'Uplink to Core Switch'")
	r.Equal("Core Uplink", port.IfAlias, "Expected IfAlias 'Core Uplink'")
	r.Equal("1000000000", port.IfSpeed, "Expected IfSpeed 1000000000")
	r.Equal("up", port.IfOperStatus, "Expected IfOperStatus 'up'")
	r.Equal("up", port.IfAdminStatus, "Expected IfAdminStatus 'up'")
	r.Equal("ethernetCsmacd", port.IfType, "Expected IfType 'ethernetCsmacd'")
	r.Equal("00:11:22:33:44:55", port.IfPhysAddress, "Expected IfPhysAddress '00:11:22:33:44:55'")
	r.Equal("1000000", port.IfInOctets, "Expected IfInOctets 1000000")
	r.Equal("2000000", port.IfOutOctets, "Expected IfOutOctets 2000000")
	r.Equal("0", port.IfInErrors, "Expected IfInErrors 0")
	r.Equal("0", port.IfOutErrors, "Expected IfOutErrors 0")
	r.Equal("300", port.PollPeriod, "Expected PollPeriod 300")
	r.Equal("0", port.Ignore, "Expected Ignore false")
	r.Equal("0", port.Disabled, "Expected Disabled false")
	r.Equal("0", port.Deleted, "Expected Deleted false")
}
