package librenms

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/javen-yan/librenms-go/types"
)

const (
	portsEndpoint = "ports"
)

// GetAllPorts retrieves all ports on all devices
// Strongly recommend using the columns parameter to avoid pulling too much data
//
// Documentation: https://docs.librenms.org/API/Ports/#get_all_ports
// Route: /api/v0/ports
func (p *PortAPI) GetAllPorts(params *types.PortsQueryParams) (*types.PortsResponse, error) {
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

	var resp types.PortsResponse
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
func (p *PortAPI) SearchPorts(search string, params *types.PortsQueryParams) (*types.PortsResponse, error) {
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

	var resp types.PortsResponse
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
func (p *PortAPI) SearchPortsInField(field, search string, params *types.PortsQueryParams) (*types.PortsResponse, error) {
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

	var resp types.PortsResponse
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
func (p *PortAPI) GetPortsWithMAC(mac string, params *types.PortsQueryParams) (*types.PortResponse, error) {
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

	var resp types.PortResponse
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
func (p *PortAPI) GetPortInfo(portID int, with ...string) (*types.PortResponse, error) {
	path := fmt.Sprintf("%s/%d", portsEndpoint, portID)

	// 处理with参数
	var queryParams *url.Values
	if len(with) > 0 {
		params := url.Values{}
		params.Set("with", with[0])
		queryParams = &params
	}

	var resp types.PortResponse
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
func (p *PortAPI) GetPortIPInfo(portID int) (*types.PortIPResponse, error) {
	path := fmt.Sprintf("%s/%d/ip", portsEndpoint, portID)
	var resp types.PortIPResponse
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
func (p *PortAPI) GetPortTransceiver(portID int) (*types.PortTransceiverResponse, error) {
	path := fmt.Sprintf("%s/%d/transceiver", portsEndpoint, portID)
	var resp types.PortTransceiverResponse
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
func (p *PortAPI) GetPortDescription(portID int) (*types.PortDescriptionResponse, error) {
	path := fmt.Sprintf("%s/%d/description", portsEndpoint, portID)
	var resp types.PortDescriptionResponse
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
func (p *PortAPI) UpdatePortDescription(portID int, description string) (*types.PortDescriptionResponse, error) {
	path := fmt.Sprintf("%s/%d/description", portsEndpoint, portID)
	req := &types.PortDescriptionUpdateRequest{Description: description}

	var resp types.PortDescriptionResponse
	httpReq, err := p.client.newRequest(http.MethodPatch, path, req, nil)
	if err != nil {
		return nil, err
	}
	err = p.client.do(httpReq, &resp)
	return &resp, err
}
