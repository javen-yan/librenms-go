package librenms

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/javen-yan/librenms-go/types"
)

const (
	vlansEndpoint = "resources/vlans"
	linksEndpoint = "resources/links"
	fdbEndpoint   = "resources/fdb"
	nacEndpoint   = "resources/nac"
)

// GetAllVLANs retrieves a list of all VLANs
//
// Documentation: https://docs.librenms.org/API/Switching/#list_vlans
// Route: /api/v0/resources/vlans
func (s *SwitchingAPI) GetAllVLANs(params *types.SwitchingQueryParams) (*types.VLANsResponse, error) {
	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Columns != nil {
			query.Set("columns", *params.Columns)
		}
		if params.Filter != nil {
			query.Set("filter", *params.Filter)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp types.VLANsResponse
	httpReq, err := s.client.newRequest(http.MethodGet, vlansEndpoint, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = s.client.do(httpReq, &resp)
	return &resp, err
}

// GetDeviceVLANs retrieves a list of all VLANs for a given device
// hostname can be either the device hostname or id
//
// Documentation: https://docs.librenms.org/API/Switching/#get_vlans
// Route: /api/v0/devices/:hostname/vlans
func (s *SwitchingAPI) GetDeviceVLANs(hostname string, params *types.SwitchingQueryParams) (*types.VLANsResponse, error) {
	path := fmt.Sprintf("devices/%s/vlans", hostname)

	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Columns != nil {
			query.Set("columns", *params.Columns)
		}
		if params.Filter != nil {
			query.Set("filter", *params.Filter)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp types.VLANsResponse
	httpReq, err := s.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = s.client.do(httpReq, &resp)
	return &resp, err
}

// GetAllLinks retrieves a list of all Links
//
// Documentation: https://docs.librenms.org/API/Switching/#list_links
// Route: /api/v0/resources/links
func (s *SwitchingAPI) GetAllLinks(params *types.SwitchingQueryParams) (*types.LinksResponse, error) {
	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Columns != nil {
			query.Set("columns", *params.Columns)
		}
		if params.Filter != nil {
			query.Set("filter", *params.Filter)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp types.LinksResponse
	httpReq, err := s.client.newRequest(http.MethodGet, linksEndpoint, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = s.client.do(httpReq, &resp)
	return &resp, err
}

// GetDeviceLinks retrieves a list of Links per given device
// hostname can be either the device hostname or id
//
// Documentation: https://docs.librenms.org/API/Switching/#get_links
// Route: /api/v0/devices/:hostname/links
func (s *SwitchingAPI) GetDeviceLinks(hostname string, params *types.SwitchingQueryParams) (*types.LinksResponse, error) {
	path := fmt.Sprintf("devices/%s/links", hostname)

	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Columns != nil {
			query.Set("columns", *params.Columns)
		}
		if params.Filter != nil {
			query.Set("filter", *params.Filter)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp types.LinksResponse
	httpReq, err := s.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = s.client.do(httpReq, &resp)
	return &resp, err
}

// GetLink retrieves Link by ID
//
// Documentation: https://docs.librenms.org/API/Switching/#get_link
// Route: /api/v0/resources/links/:id
func (s *SwitchingAPI) GetLink(linkID int, params *types.SwitchingQueryParams) (*types.LinksResponse, error) {
	path := fmt.Sprintf("%s/%d", linksEndpoint, linkID)

	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Columns != nil {
			query.Set("columns", *params.Columns)
		}
		if params.Filter != nil {
			query.Set("filter", *params.Filter)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp types.LinksResponse
	httpReq, err := s.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = s.client.do(httpReq, &resp)
	return &resp, err
}

// GetPortFDB retrieves a list of all ports FDB
// mac is the specific MAC address you would like to query
//
// Documentation: https://docs.librenms.org/API/Switching/#list_fdb
// Route: /api/v0/resources/fdb/:mac
func (s *SwitchingAPI) GetPortFDB(mac string, params *types.SwitchingQueryParams) (*types.PortFDBResponse, error) {
	path := fdbEndpoint
	if mac != "" {
		path = fmt.Sprintf("%s/%s", fdbEndpoint, mac)
	}

	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Columns != nil {
			query.Set("columns", *params.Columns)
		}
		if params.Filter != nil {
			query.Set("filter", *params.Filter)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp types.PortFDBResponse
	httpReq, err := s.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = s.client.do(httpReq, &resp)
	return &resp, err
}

// GetPortFDBDetail retrieves a list of all ports FDB with human readable device and interface names
// mac is the specific MAC address you would like to query
//
// Documentation: https://docs.librenms.org/API/Switching/#list_fdb_detail
// Route: /api/v0/resources/fdb/:mac/detail
func (s *SwitchingAPI) GetPortFDBDetail(mac string, params *types.SwitchingQueryParams) (*types.PortFDBDetailResponse, error) {
	path := fmt.Sprintf("%s/%s/detail", fdbEndpoint, mac)

	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Columns != nil {
			query.Set("columns", *params.Columns)
		}
		if params.Filter != nil {
			query.Set("filter", *params.Filter)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp types.PortFDBDetailResponse
	httpReq, err := s.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = s.client.do(httpReq, &resp)
	return &resp, err
}

// GetPortNAC retrieves a list of all ports NAC
// mac is the specific MAC address you would like to query
//
// Documentation: https://docs.librenms.org/API/Switching/#list_nac
// Route: /api/v0/resources/nac/:mac
func (s *SwitchingAPI) GetPortNAC(mac string, params *types.SwitchingQueryParams) (*types.PortNACResponse, error) {
	path := nacEndpoint
	if mac != "" {
		path = fmt.Sprintf("%s/%s", nacEndpoint, mac)
	}

	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Columns != nil {
			query.Set("columns", *params.Columns)
		}
		if params.Filter != nil {
			query.Set("filter", *params.Filter)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp types.PortNACResponse
	httpReq, err := s.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = s.client.do(httpReq, &resp)
	return &resp, err
}
