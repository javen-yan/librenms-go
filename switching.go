package librenms

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	vlansEndpoint = "resources/vlans"
	linksEndpoint = "resources/links"
	fdbEndpoint   = "resources/fdb"
	nacEndpoint   = "resources/nac"
)

type (
	VLAN struct {
		VLANID     string `json:"vlan_id"`
		DeviceID   string `json:"device_id"`
		VLANVLAN   string `json:"vlan_vlan"`
		VLANDomain string `json:"vlan_domain"`
		VLANName   string `json:"vlan_name"`
		VLANType   string `json:"vlan_type"`
		VLANState  int    `json:"vlan_state"`
	}

	Link struct {
		ID                int     `json:"id"`
		LocalPortID       int     `json:"local_port_id"`
		LocalDeviceID     int     `json:"local_device_id"`
		RemotePortID      int     `json:"remote_port_id"`
		Active            int     `json:"active"`
		Protocol          string  `json:"protocol"`
		RemoteHostname    string  `json:"remote_hostname"`
		RemoteDeviceID    int     `json:"remote_device_id"`
		RemotePort        string  `json:"remote_port"`
		RemotePlatform    *string `json:"remote_platform"`
		RemoteVersion     string  `json:"remote_version"`
	}

	PortFDB struct {
		PortsFDBID int    `json:"ports_fdb_id"`
		PortID     int    `json:"port_id"`
		MACAddress string `json:"mac_address"`
		VLANID     int    `json:"vlan_id"`
		DeviceID   int    `json:"device_id"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	}

	PortFDBDetail struct {
		Hostname  string `json:"hostname"`
		SysName   string `json:"sysName"`
		IfName    string `json:"ifName"`
		IfAlias   string `json:"ifAlias"`
		IfDescr   string `json:"ifDescr"`
		LastSeen  string `json:"last_seen"`
		UpdatedAt string `json:"updated_at"`
	}

	PortNAC struct {
		PortsNACID      string  `json:"ports_nac_id"`
		AuthID          string  `json:"auth_id"`
		DeviceID        int     `json:"device_id"`
		PortID          int     `json:"port_id"`
		Domain          string  `json:"domain"`
		Username        string  `json:"username"`
		MACAddress      string  `json:"mac_address"`
		IPAddress       string  `json:"ip_address"`
		HostMode        string  `json:"host_mode"`
		AuthzStatus     string  `json:"authz_status"`
		AuthzBy         string  `json:"authz_by"`
		AuthcStatus     string  `json:"authc_status"`
		Method          string  `json:"method"`
		Timeout         string  `json:"timeout"`
		TimeLeft        string  `json:"time_left"`
		VLAN            int     `json:"vlan"`
		TimeElapsed     *string `json:"time_elapsed"`
		CreatedAt       string  `json:"created_at"`
		UpdatedAt       string  `json:"updated_at"`
		Historical      int     `json:"historical"`
	}

	// API响应结构
	VLANsResponse struct {
		BaseResponse
		VLANs []VLAN `json:"vlans"`
	}

	LinksResponse struct {
		BaseResponse
		Links []Link `json:"links"`
	}

	PortFDBResponse struct {
		BaseResponse
		PortsFDB []PortFDB `json:"ports_fdb"`
	}

	PortFDBDetailResponse struct {
		BaseResponse
		MAC     string           `json:"mac"`
		MACOUI  string           `json:"mac_oui"`
		PortsFDB []PortFDBDetail `json:"ports_fdb"`
	}

	PortNACResponse struct {
		BaseResponse
		PortsNAC []PortNAC `json:"ports_nac"`
	}

	// 查询参数结构
	SwitchingQueryParams struct {
		Columns *string `url:"columns,omitempty"`
		Filter  *string `url:"filter,omitempty"`
	}
)

// GetAllVLANs retrieves a list of all VLANs
//
// Documentation: https://docs.librenms.org/API/Switching/#list_vlans
// Route: /api/v0/resources/vlans
func (s *SwitchingAPI) GetAllVLANs(params *SwitchingQueryParams) (*VLANsResponse, error) {
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

	var resp VLANsResponse
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
func (s *SwitchingAPI) GetDeviceVLANs(hostname string, params *SwitchingQueryParams) (*VLANsResponse, error) {
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

	var resp VLANsResponse
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
func (s *SwitchingAPI) GetAllLinks(params *SwitchingQueryParams) (*LinksResponse, error) {
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

	var resp LinksResponse
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
func (s *SwitchingAPI) GetDeviceLinks(hostname string, params *SwitchingQueryParams) (*LinksResponse, error) {
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

	var resp LinksResponse
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
func (s *SwitchingAPI) GetLink(linkID int, params *SwitchingQueryParams) (*LinksResponse, error) {
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

	var resp LinksResponse
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
func (s *SwitchingAPI) GetPortFDB(mac string, params *SwitchingQueryParams) (*PortFDBResponse, error) {
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

	var resp PortFDBResponse
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
func (s *SwitchingAPI) GetPortFDBDetail(mac string, params *SwitchingQueryParams) (*PortFDBDetailResponse, error) {
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

	var resp PortFDBDetailResponse
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
func (s *SwitchingAPI) GetPortNAC(mac string, params *SwitchingQueryParams) (*PortNACResponse, error) {
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

	var resp PortNACResponse
	httpReq, err := s.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = s.client.do(httpReq, &resp)
	return &resp, err
}
