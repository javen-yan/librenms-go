package librenms

import (
	"net/http"

	"github.com/javen-yan/librenms-go/types"
)

const (
	systemEndpoint = "system"
)

// Get retrieves system information from LibreNMS
func (s *SystemAPI) Get() (*types.SystemResponse, error) {
	c := s.client
	req, err := c.newRequest(http.MethodGet, systemEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	systemResp := new(types.SystemResponse)
	return systemResp, c.do(req, systemResp)
}
