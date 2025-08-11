package librenms

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/javen-yan/librenms-go/types"
)

const (
	// BGP endpoints
	bgpEndpoint = "bgp"

	// OSPF endpoints
	ospfEndpoint        = "ospf"
	ospfPortsEndpoint   = "ospf_ports"
	ospfv3Endpoint      = "ospfv3"
	ospfv3PortsEndpoint = "ospfv3_ports"

	// VRF endpoints
	vrfEndpoint = "routing/vrf"

	// MPLS endpoints
	mplsServicesEndpoint = "routing/mpls/services"
	mplsSapsEndpoint     = "routing/mpls/saps"

	// IPSec endpoints
	ipsecEndpoint = "routing/ipsec/data"

	// IP resources endpoints
	ipAddressesEndpoint        = "resources/ip/addresses"
	ipNetworksEndpoint         = "resources/ip/networks"
	ipNetworkAddressesEndpoint = "resources/ip/networks"

	// BGP counters endpoint
	bgpCountersEndpoint = "routing/bgp/cbgp"
)

// ListBGP retrieves a list of BGP sessions from the LibreNMS API
func (r *RoutingAPI) ListBGP(query *types.BGPQuery) (*types.BGPResponse, error) {
	c := r.client
	params, err := parseParams(query)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodGet, bgpEndpoint, nil, params)
	if err != nil {
		return nil, err
	}

	bgpResp := new(types.BGPResponse)
	return bgpResp, c.do(req, bgpResp)
}

// GetBGP retrieves a BGP session by ID from the LibreNMS API
func (r *RoutingAPI) GetBGP(id string) (*types.BGPSessionResponse, error) {
	c := r.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s", bgpEndpoint, id), nil, nil)
	if err != nil {
		return nil, err
	}

	bgpResp := new(types.BGPSessionResponse)
	return bgpResp, c.do(req, bgpResp)
}

// UpdateBGPDescription updates the description of a BGP session
func (r *RoutingAPI) UpdateBGPDescription(id string, payload *types.BGPDescriptionUpdate) (*types.BaseResponse, error) {
	c := r.client
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/%s", bgpEndpoint, id), payload, nil)
	if err != nil {
		return nil, err
	}

	resp := new(types.BaseResponse)
	return resp, c.do(req, resp)
}

// ListBGPCounters retrieves a list of BGP counters from the LibreNMS API
func (r *RoutingAPI) ListBGPCounters(hostname string) (*types.BGPCountersResponse, error) {
	c := r.client
	var params *url.Values
	if hostname != "" {
		p := url.Values{}
		p.Set("hostname", hostname)
		params = &p
	}

	req, err := c.newRequest(http.MethodGet, bgpCountersEndpoint, nil, params)
	if err != nil {
		return nil, err
	}

	countersResp := new(types.BGPCountersResponse)
	return countersResp, c.do(req, countersResp)
}

// ListIPAddresses retrieves a list of IP addresses from the LibreNMS API
func (r *RoutingAPI) ListIPAddresses(addressFamily string) (*types.IPAddressesResponse, error) {
	c := r.client
	endpoint := ipAddressesEndpoint
	if addressFamily != "" {
		endpoint = fmt.Sprintf("%s/%s", ipAddressesEndpoint, addressFamily)
	}

	req, err := c.newRequest(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	addressesResp := new(types.IPAddressesResponse)
	return addressesResp, c.do(req, addressesResp)
}

// GetNetworkIPAddresses retrieves IP addresses for a specific network
func (r *RoutingAPI) GetNetworkIPAddresses(networkID string) (*types.IPAddressesResponse, error) {
	c := r.client
	endpoint := fmt.Sprintf("%s/%s/ip", ipNetworkAddressesEndpoint, networkID)

	req, err := c.newRequest(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	addressesResp := new(types.IPAddressesResponse)
	return addressesResp, c.do(req, addressesResp)
}

// ListIPNetworks retrieves a list of IP networks from the LibreNMS API
func (r *RoutingAPI) ListIPNetworks(addressFamily string) (*types.IPNetworksResponse, error) {
	c := r.client
	endpoint := ipNetworksEndpoint
	if addressFamily != "" {
		endpoint = fmt.Sprintf("%s/%s", ipNetworksEndpoint, addressFamily)
	}

	req, err := c.newRequest(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	networksResp := new(types.IPNetworksResponse)
	return networksResp, c.do(req, networksResp)
}

// ListIPSec retrieves a list of IPSec tunnels from the LibreNMS API
func (r *RoutingAPI) ListIPSec(hostname string) (*types.IPSecResponse, error) {
	c := r.client
	endpoint := fmt.Sprintf("%s/%s", ipsecEndpoint, hostname)

	req, err := c.newRequest(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	ipsecResp := new(types.IPSecResponse)
	return ipsecResp, c.do(req, ipsecResp)
}

// ListOSPF retrieves a list of OSPF neighbors from the LibreNMS API
func (r *RoutingAPI) ListOSPF(hostname string) (*types.OSPFResponse, error) {
	c := r.client
	var params *url.Values
	if hostname != "" {
		p := url.Values{}
		p.Set("hostname", hostname)
		params = &p
	}

	req, err := c.newRequest(http.MethodGet, ospfEndpoint, nil, params)
	if err != nil {
		return nil, err
	}

	ospfResp := new(types.OSPFResponse)
	return ospfResp, c.do(req, ospfResp)
}

// ListOSPFPorts retrieves a list of OSPF ports from the LibreNMS API
func (r *RoutingAPI) ListOSPFPorts() (*types.OSPFPortsResponse, error) {
	c := r.client
	req, err := c.newRequest(http.MethodGet, ospfPortsEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	portsResp := new(types.OSPFPortsResponse)
	return portsResp, c.do(req, portsResp)
}

// ListOSPFv3 retrieves a list of OSPFv3 neighbors from the LibreNMS API
func (r *RoutingAPI) ListOSPFv3(hostname string) (*types.OSPFv3Response, error) {
	c := r.client
	var params *url.Values
	if hostname != "" {
		p := url.Values{}
		p.Set("hostname", hostname)
		params = &p
	}

	req, err := c.newRequest(http.MethodGet, ospfv3Endpoint, nil, params)
	if err != nil {
		return nil, err
	}

	ospfv3Resp := new(types.OSPFv3Response)
	return ospfv3Resp, c.do(req, ospfv3Resp)
}

// ListOSPFv3Ports retrieves a list of OSPFv3 ports from the LibreNMS API
func (r *RoutingAPI) ListOSPFv3Ports() (*types.OSPFv3PortsResponse, error) {
	c := r.client
	req, err := c.newRequest(http.MethodGet, ospfv3PortsEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	portsResp := new(types.OSPFv3PortsResponse)
	return portsResp, c.do(req, portsResp)
}

// ListVRF retrieves a list of VRFs from the LibreNMS API
func (r *RoutingAPI) ListVRF(query *types.VRFQuery) (*types.VRFResponse, error) {
	c := r.client
	params, err := parseParams(query)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodGet, vrfEndpoint, nil, params)
	if err != nil {
		return nil, err
	}

	vrfResp := new(types.VRFResponse)
	return vrfResp, c.do(req, vrfResp)
}

// GetVRF retrieves a VRF by ID from the LibreNMS API
func (r *RoutingAPI) GetVRF(id string) (*types.VRFResponse, error) {
	c := r.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s", vrfEndpoint, id), nil, nil)
	if err != nil {
		return nil, err
	}

	vrfResp := new(types.VRFResponse)
	return vrfResp, c.do(req, vrfResp)
}

// ListMPLSServices retrieves a list of MPLS services from the LibreNMS API
func (r *RoutingAPI) ListMPLSServices(hostname string) (*types.MPLSServicesResponse, error) {
	c := r.client
	var params *url.Values
	if hostname != "" {
		p := url.Values{}
		p.Set("hostname", hostname)
		params = &p
	}

	req, err := c.newRequest(http.MethodGet, mplsServicesEndpoint, nil, params)
	if err != nil {
		return nil, err
	}

	servicesResp := new(types.MPLSServicesResponse)
	return servicesResp, c.do(req, servicesResp)
}

// ListMPLSSAPs retrieves a list of MPLS SAPs from the LibreNMS API
func (r *RoutingAPI) ListMPLSSAPs(hostname string) (*types.MPLSSAPsResponse, error) {
	c := r.client
	var params *url.Values
	if hostname != "" {
		p := url.Values{}
		p.Set("hostname", hostname)
		params = &p
	}

	req, err := c.newRequest(http.MethodGet, mplsSapsEndpoint, nil, params)
	if err != nil {
		return nil, err
	}

	sapsResp := new(types.MPLSSAPsResponse)
	return sapsResp, c.do(req, sapsResp)
}
