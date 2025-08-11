package librenms

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	portsEndpoint = "ports"
)

type (
	Port struct {
		PortID                  string  `json:"port_id"`
		DeviceID                string  `json:"device_id"`
		PortDescrType           *string `json:"port_descr_type"`
		PortDescrDescr          *string `json:"port_descr_descr"`
		PortDescrCircuit        *string `json:"port_descr_circuit"`
		PortDescrSpeed          *string `json:"port_descr_speed"`
		PortDescrNotes          *string `json:"port_descr_notes"`
		IfDescr                 string  `json:"ifDescr"`
		IfName                  string  `json:"ifName"`
		PortName                *string `json:"portName"`
		IfIndex                 string  `json:"ifIndex"`
		IfSpeed                 string  `json:"ifSpeed"`
		IfConnectorPresent      string  `json:"ifConnectorPresent"`
		IfPromiscuousMode       string  `json:"ifPromiscuousMode"`
		IfHighSpeed             string  `json:"ifHighSpeed"`
		IfOperStatus            string  `json:"ifOperStatus"`
		IfOperStatusPrev        *string `json:"ifOperStatus_prev"`
		IfAdminStatus           string  `json:"ifAdminStatus"`
		IfAdminStatusPrev       *string `json:"ifAdminStatus_prev"`
		IfDuplex                string  `json:"ifDuplex"`
		IfMtu                   string  `json:"ifMtu"`
		IfType                  string  `json:"ifType"`
		IfAlias                 string  `json:"ifAlias"`
		IfPhysAddress           string  `json:"ifPhysAddress"`
		IfHardType              *string `json:"ifHardType"`
		IfLastChange            string  `json:"ifLastChange"`
		IfVlan                  string  `json:"ifVlan"`
		IfTrunk                 string  `json:"ifTrunk"`
		IfVrf                   string  `json:"ifVrf"`
		CounterIn               *string `json:"counter_in"`
		CounterOut              *string `json:"counter_out"`
		Ignore                  string  `json:"ignore"`
		Disabled                string  `json:"disabled"`
		Detailed                string  `json:"detailed"`
		Deleted                 string  `json:"deleted"`
		PagpOperationMode       *string `json:"pagpOperationMode"`
		PagpPortState           *string `json:"pagpPortState"`
		PagpPartnerDeviceId     *string `json:"pagpPartnerDeviceId"`
		PagpPartnerLearnMethod  *string `json:"pagpPartnerLearnMethod"`
		PagpPartnerIfIndex      *string `json:"pagpPartnerIfIndex"`
		PagpPartnerGroupIfIndex *string `json:"pagpPartnerGroupIfIndex"`
		PagpPartnerDeviceName   *string `json:"pagpPartnerDeviceName"`
		PagpEthcOperationMode   *string `json:"pagpEthcOperationMode"`
		PagpDeviceId            *string `json:"pagpDeviceId"`
		PagpGroupIfIndex        *string `json:"pagpGroupIfIndex"`
		IfInUcastPkts           string  `json:"ifInUcastPkts"`
		IfInUcastPktsPrev       string  `json:"ifInUcastPkts_prev"`
		IfInUcastPktsDelta      string  `json:"ifInUcastPkts_delta"`
		IfInUcastPktsRate       string  `json:"ifInUcastPkts_rate"`
		IfOutUcastPkts          string  `json:"ifOutUcastPkts"`
		IfOutUcastPktsPrev      string  `json:"ifOutUcastPkts_prev"`
		IfOutUcastPktsDelta     string  `json:"ifOutUcastPkts_delta"`
		IfOutUcastPktsRate      string  `json:"ifOutUcastPkts_rate"`
		IfInErrors              string  `json:"ifInErrors"`
		IfInErrorsPrev          string  `json:"ifInErrors_prev"`
		IfInErrorsDelta         string  `json:"ifInErrors_delta"`
		IfInErrorsRate          string  `json:"ifInErrors_rate"`
		IfOutErrors             string  `json:"ifOutErrors"`
		IfOutErrorsPrev         string  `json:"ifOutErrors_prev"`
		IfOutErrorsDelta        string  `json:"ifOutErrors_delta"`
		IfOutErrorsRate         string  `json:"ifOutErrors_rate"`
		IfInOctets              string  `json:"ifInOctets"`
		IfInOctetsPrev          string  `json:"ifInOctets_prev"`
		IfInOctetsDelta         string  `json:"ifInOctets_delta"`
		IfInOctetsRate          string  `json:"ifInOctets_rate"`
		IfOutOctets             string  `json:"ifOutOctets"`
		IfOutOctetsPrev         string  `json:"ifOutOctets_prev"`
		IfOutOctetsDelta        string  `json:"ifOutOctets_delta"`
		IfOutOctetsRate         string  `json:"ifOutOctets_rate"`
		PollTime                string  `json:"poll_time"`
		PollPrev                string  `json:"poll_prev"`
		PollPeriod              string  `json:"poll_period"`
	}

	PortIPAddress struct {
		IPv4AddressID string `json:"ipv4_address_id"`
		IPv4Address   string `json:"ipv4_address"`
		IPv4PrefixLen string `json:"ipv4_prefixlen"`
		IPv4NetworkID string `json:"ipv4_network_id"`
		PortID        string `json:"port_id"`
		ContextName   string `json:"context_name"`
	}

	PortTransceiver struct {
		ID         int     `json:"id"`
		CreatedAt  string  `json:"created_at"`
		UpdatedAt  string  `json:"updated_at"`
		DeviceID   int     `json:"device_id"`
		PortID     int     `json:"port_id"`
		Index      string  `json:"index"`
		Type       string  `json:"type"`
		Vendor     string  `json:"vendor"`
		OUI        string  `json:"oui"`
		Model      *string `json:"model"`
		Revision   string  `json:"revision"`
		Serial     string  `json:"serial"`
		Date       *string `json:"date"`
		DDM        bool    `json:"ddm"`
		Encoding   *string `json:"encoding"`
		Cable      string  `json:"cable"`
		Distance   int     `json:"distance"`
		Wavelength int     `json:"wavelength"`
		Connector  string  `json:"connector"`
		Channels   int     `json:"channels"`
	}

	PortDescriptionUpdateRequest struct {
		Description string `json:"description"`
	}

	// API响应结构 - 根据文档更新
	PortsResponse struct {
		BaseResponse
		Ports []Port `json:"ports"`
	}

	PortResponse struct {
		BaseResponse
		Port []Port `json:"port"`
	}

	PortIPResponse struct {
		BaseResponse
		Addresses []PortIPAddress `json:"addresses"`
	}

	PortTransceiverResponse struct {
		BaseResponse
		Transceivers []PortTransceiver `json:"transceivers"`
	}

	PortDescriptionResponse struct {
		BaseResponse
		PortDescription string `json:"port_description"`
	}

	// 查询参数结构
	PortsQueryParams struct {
		Columns *string `url:"columns,omitempty"`
		Filter  *string `url:"filter,omitempty"`
	}
)

// GetAllPorts retrieves all ports on all devices
// Strongly recommend using the columns parameter to avoid pulling too much data
//
// Documentation: https://docs.librenms.org/API/Ports/#get_all_ports
// Route: /api/v0/ports
func (p *PortAPI) GetAllPorts(params *PortsQueryParams) (*PortsResponse, error) {
	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Columns != nil {
			query.Set("columns", *params.Columns)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp PortsResponse
	httpReq, err := p.client.newRequest(http.MethodGet, portsEndpoint, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = p.client.do(httpReq, &resp)
	return &resp, err
}

// SearchPorts searches for ports matching the query
// Search in fields: ifAlias, ifDescr, and ifName
//
// Documentation: https://docs.librenms.org/API/Ports/#search_ports
// Route: /api/v0/ports/search/:search
func (p *PortAPI) SearchPorts(search string, params *PortsQueryParams) (*PortsResponse, error) {
	path := fmt.Sprintf("%s/search/%s", portsEndpoint, search)

	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Columns != nil {
			query.Set("columns", *params.Columns)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp PortsResponse
	httpReq, err := p.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = p.client.do(httpReq, &resp)
	return &resp, err
}

// SearchPortsInField searches for ports in specific field(s)
// Specific search for ports matching the query in specified fields
//
// Documentation: https://docs.librenms.org/API/Ports/#search_ports_in_specific_fields
// Route: /api/v0/ports/search/:field/:search
func (p *PortAPI) SearchPortsInField(field, search string, params *PortsQueryParams) (*PortsResponse, error) {
	path := fmt.Sprintf("%s/search/%s/%s", portsEndpoint, field, search)

	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Columns != nil {
			query.Set("columns", *params.Columns)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp PortsResponse
	httpReq, err := p.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = p.client.do(httpReq, &resp)
	return &resp, err
}

// GetPortsWithMAC searches for ports matching the search MAC address
// Search a MAC address in FDB and print the ports ordered by the MAC count of the associated port
//
// Documentation: https://docs.librenms.org/API/Ports/#ports_with_associated_mac
// Route: /api/v0/ports/mac/:search
func (p *PortAPI) GetPortsWithMAC(mac string, params *PortsQueryParams) (*PortResponse, error) {
	path := fmt.Sprintf("%s/mac/%s", portsEndpoint, mac)

	var queryParams *url.Values
	if params != nil {
		query := url.Values{}
		if params.Filter != nil {
			query.Set("filter", *params.Filter)
		}
		if len(query) > 0 {
			queryParams = &query
		}
	}

	var resp PortResponse
	httpReq, err := p.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = p.client.do(httpReq, &resp)
	return &resp, err
}

// GetPortInfo retrieves detailed information for a specific port
// It's possible to add allowed associated relations using the with option (vlans, device)
//
// Documentation: https://docs.librenms.org/API/Ports/#get_port_info
// Route: /api/v0/ports/:portid
func (p *PortAPI) GetPortInfo(portID int, with ...string) (*PortResponse, error) {
	path := fmt.Sprintf("%s/%d", portsEndpoint, portID)

	// 处理with参数
	var queryParams *url.Values
	if len(with) > 0 {
		params := url.Values{}
		params.Set("with", with[0])
		queryParams = &params
	}

	var resp PortResponse
	httpReq, err := p.client.newRequest(http.MethodGet, path, nil, queryParams)
	if err != nil {
		return nil, err
	}
	err = p.client.do(httpReq, &resp)
	return &resp, err
}

// GetPortIPInfo retrieves IP information for a specific port
// Get all IP info (v4 and v6) for a given port id
//
// Documentation: https://docs.librenms.org/API/Ports/#get_port_ip_info
// Route: /api/v0/ports/:portid/ip
func (p *PortAPI) GetPortIPInfo(portID int) (*PortIPResponse, error) {
	path := fmt.Sprintf("%s/%d/ip", portsEndpoint, portID)
	var resp PortIPResponse
	httpReq, err := p.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}
	err = p.client.do(httpReq, &resp)
	return &resp, err
}

// GetPortTransceiver retrieves transceiver information for a specific port
// Get transceiver info with metrics
//
// Documentation: https://docs.librenms.org/API/Ports/#get_port_transceiver
// Route: /api/v0/ports/:portid/transceiver
func (p *PortAPI) GetPortTransceiver(portID int) (*PortTransceiverResponse, error) {
	path := fmt.Sprintf("%s/%d/transceiver", portsEndpoint, portID)
	var resp PortTransceiverResponse
	httpReq, err := p.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}
	err = p.client.do(httpReq, &resp)
	return &resp, err
}

// GetPortDescription retrieves the description (ifAlias) for a specific port
//
// Documentation: https://docs.librenms.org/API/Ports/#get_port_description
// Route: /api/v0/ports/:portid/description
func (p *PortAPI) GetPortDescription(portID int) (*PortDescriptionResponse, error) {
	path := fmt.Sprintf("%s/%d/description", portsEndpoint, portID)
	var resp PortDescriptionResponse
	httpReq, err := p.client.newRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}
	err = p.client.do(httpReq, &resp)
	return &resp, err
}

// UpdatePortDescription updates the description (ifAlias) for a specific port
// Sending an empty string will reset the description to default
//
// Documentation: https://docs.librenms.org/API/Ports/#update_port_description
// Route: /api/v0/ports/:portid/description
func (p *PortAPI) UpdatePortDescription(portID int, description string) (*PortDescriptionResponse, error) {
	path := fmt.Sprintf("%s/%d/description", portsEndpoint, portID)
	req := &PortDescriptionUpdateRequest{Description: description}

	var resp PortDescriptionResponse
	httpReq, err := p.client.newRequest(http.MethodPatch, path, req, nil)
	if err != nil {
		return nil, err
	}
	err = p.client.do(httpReq, &resp)
	return &resp, err
}
