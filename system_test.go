package librenms_test

import (
	"net/http"
	"testing"

	"github.com/javen-yan/librenms-go"
	"github.com/stretchr/testify/require"
)

const (
	testEndpointSystem = "/api/v0/system"
)

// This init function will register handlers for system-related API endpoints.
func init() {
	handleEndpoint(testEndpointSystem, mockResponses{
		http.MethodGet: loadMockResponse("get_system_200.json"),
	})
}

func TestClient_GetSystem(t *testing.T) {
	r := require.New(t)

	r.NotNil(testAPIClient, "Global testAPIClient should be initialized")

	systemResp, err := testAPIClient.System.Get()

	r.NoError(err, "GetSystem returned an error")
	r.NotNil(systemResp, "GetSystem response is nil")

	r.Equal("ok", systemResp.Status, "Expected status 'ok'")
	r.Equal(1, systemResp.Count, "Expected count 1")
	r.Len(systemResp.System, 1, "Expected 1 system info")

	system := systemResp.System[0]
	r.Equal("23.11.0", system.LocalVer, "Expected LocalVer '23.11.0'")
	r.Equal("abc123def456", system.LocalSha, "Expected LocalSha 'abc123def456'")
	r.Equal("2023-11-15", system.LocalDate, "Expected LocalDate '2023-11-15'")
	r.Equal("master", system.LocalBranch, "Expected LocalBranch 'master'")
	r.Equal("2023_11_01_000000", system.DBSchema, "Expected DBSchema '2023_11_01_000000'")
	r.Equal("8.1.0", system.PHPVer, "Expected PHPVer '8.1.0'")
	r.Equal("3.9.0", system.PythonVer, "Expected PythonVer '3.9.0'")
	r.Equal("10.5.0", system.DatabaseVer, "Expected DatabaseVer '10.5.0'")
	r.Equal("1.7.2", system.RRDToolVer, "Expected RRDToolVer '1.7.2'")
	r.Equal("5.9.1", system.NetSNMPVer, "Expected NetSNMPVer '5.9.1'")
}

func TestSystemInfo_StructFields(t *testing.T) {
	r := require.New(t)

	// Test that SystemInfo struct can be properly marshaled/unmarshaled
	systemInfo := librenms.SystemInfo{
		LocalVer:    "1.0.0",
		LocalSha:    "test123",
		LocalDate:   "2023-01-01",
		LocalBranch: "develop",
		DBSchema:    "2023_01_01_000000",
		PHPVer:      "8.0.0",
		PythonVer:   "3.8.0",
		DatabaseVer: "10.0.0",
		RRDToolVer:  "1.6.0",
		NetSNMPVer:  "5.8.0",
	}

	r.Equal("1.0.0", systemInfo.LocalVer)
	r.Equal("test123", systemInfo.LocalSha)
	r.Equal("2023-01-01", systemInfo.LocalDate)
	r.Equal("develop", systemInfo.LocalBranch)
	r.Equal("2023_01_01_000000", systemInfo.DBSchema)
	r.Equal("8.0.0", systemInfo.PHPVer)
	r.Equal("3.8.0", systemInfo.PythonVer)
	r.Equal("10.0.0", systemInfo.DatabaseVer)
	r.Equal("1.6.0", systemInfo.RRDToolVer)
	r.Equal("5.8.0", systemInfo.NetSNMPVer)
}

func TestSystemResponse_StructFields(t *testing.T) {
	r := require.New(t)

	// Test that SystemResponse struct can be properly initialized
	systemResp := &librenms.SystemResponse{
		BaseResponse: librenms.BaseResponse{
			Status:  "ok",
			Message: "Success",
			Count:   1,
		},
		System: []librenms.SystemInfo{
			{
				LocalVer: "1.0.0",
				LocalSha: "test123",
			},
		},
		Count: 1,
	}

	r.Equal("ok", systemResp.Status)
	r.Equal("Success", systemResp.Message)
	r.Equal(1, systemResp.Count)
	r.Len(systemResp.System, 1)
	r.Equal("1.0.0", systemResp.System[0].LocalVer)
	r.Equal("test123", systemResp.System[0].LocalSha)
}
