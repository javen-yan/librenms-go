package librenms

import (
	"net/http"
)

const (
	systemEndpoint = "system"
)

// SystemInfo represents system information from LibreNMS
type SystemInfo struct {
	LocalVer    string `json:"local_ver"`
	LocalSha    string `json:"local_sha"`
	LocalDate   string `json:"local_date"`
	LocalBranch string `json:"local_branch"`
	DBSchema    string `json:"db_schema"`
	PHPVer      string `json:"php_ver"`
	PythonVer   string `json:"python_ver"`
	DatabaseVer string `json:"database_ver"`
	RRDToolVer  string `json:"rrdtool_ver"`
	NetSNMPVer  string `json:"netsnmp_ver"`
}

// SystemResponse represents the response from the system API endpoint
type SystemResponse struct {
	BaseResponse
	System []SystemInfo `json:"system"`
	Count  int          `json:"count"`
}

// Get retrieves system information from LibreNMS
func (s *SystemAPI) Get() (*SystemResponse, error) {
	c := s.client
	req, err := c.newRequest(http.MethodGet, systemEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	systemResp := new(SystemResponse)
	return systemResp, c.do(req, systemResp)
}
