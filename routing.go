package librenms

import (
	"fmt"
	"net/http"
	"net/url"
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

type (
	// BGP Session represents a BGP session in LibreNMS
	BGPSession struct {
		BGPPeerID                  int    `json:"bgpPeer_id"`
		DeviceID                   int    `json:"device_id"`
		VRFID                      *int   `json:"vrf_id"`
		ASText                     string `json:"astext"`
		BGPPeerIdentifier          string `json:"bgpPeerIdentifier"`
		BGPPeerRemoteAS            int    `json:"bgpPeerRemoteAs"`
		BGPPeerState               string `json:"bgpPeerState"`
		BGPPeerAdminStatus         string `json:"bgpPeerAdminStatus"`
		BGPPeerLastErrorCode       int    `json:"bgpPeerLastErrorCode"`
		BGPPeerLastErrorSubCode    int    `json:"bgpPeerLastErrorSubCode"`
		BGPPeerLastErrorText       string `json:"bgpPeerLastErrorText"`
		BGPPeerIface               int    `json:"bgpPeerIface"`
		BGPLocalAddr               string `json:"bgpLocalAddr"`
		BGPPeerRemoteAddr          string `json:"bgpPeerRemoteAddr"`
		BGPPeerDescr               string `json:"bgpPeerDescr"`
		BGPPeerInUpdates           int64  `json:"bgpPeerInUpdates"`
		BGPPeerOutUpdates          int64  `json:"bgpPeerOutUpdates"`
		BGPPeerInTotalMessages     int64  `json:"bgpPeerInTotalMessages"`
		BGPPeerOutTotalMessages    int64  `json:"bgpPeerOutTotalMessages"`
		BGPPeerFsmEstablishedTime  int64  `json:"bgpPeerFsmEstablishedTime"`
		BGPPeerInUpdateElapsedTime int64  `json:"bgpPeerInUpdateElapsedTime"`
		ContextName                string `json:"context_name"`
	}

	// BGPQuery represents the query parameters for filtering BGP sessions
	BGPQuery struct {
		Hostname      string `url:"hostname,omitempty"`
		ASN           int    `url:"asn,omitempty"`
		RemoteASN     int    `url:"remote_asn,omitempty"`
		RemoteAddress string `url:"remote_address,omitempty"`
		LocalAddress  string `url:"local_address,omitempty"`
		BGPDescr      string `url:"bgp_descr,omitempty"`
		BGPState      string `url:"bgp_state,omitempty"`
		BGPAdminState string `url:"bgp_adminstate,omitempty"`
		BGPFamily     int    `url:"bgp_family,omitempty"`
	}

	// BGPResponse represents a response containing BGP sessions
	BGPResponse struct {
		BaseResponse
		BGPSessions []BGPSession `json:"bgp_sessions"`
	}

	// BGPSessionResponse represents a response containing a single BGP session
	BGPSessionResponse struct {
		BaseResponse
		BGPSession []BGPSession `json:"bgp_session"`
	}

	// BGPDescriptionUpdate represents the request body for updating BGP description
	BGPDescriptionUpdate struct {
		BGPDescr string `json:"bgp_descr"`
	}

	// BGPCounters represents BGP counters in LibreNMS
	BGPCounters struct {
		DeviceID                string `json:"device_id"`
		BGPPeerIdentifier       string `json:"bgpPeerIdentifier"`
		AFI                     string `json:"afi"`
		SAFI                    string `json:"safi"`
		AcceptedPrefixes        string `json:"AcceptedPrefixes"`
		DeniedPrefixes          string `json:"DeniedPrefixes"`
		PrefixAdminLimit        string `json:"PrefixAdminLimit"`
		PrefixThreshold         string `json:"PrefixThreshold"`
		PrefixClearThreshold    string `json:"PrefixClearThreshold"`
		AdvertisedPrefixes      string `json:"AdvertisedPrefixes"`
		SuppressedPrefixes      string `json:"SuppressedPrefixes"`
		WithdrawnPrefixes       string `json:"WithdrawnPrefixes"`
		AcceptedPrefixesDelta   string `json:"AcceptedPrefixes_delta"`
		AcceptedPrefixesPrev    string `json:"AcceptedPrefixes_prev"`
		DeniedPrefixesDelta     string `json:"DeniedPrefixes_delta"`
		DeniedPrefixesPrev      string `json:"DeniedPrefixes_prev"`
		AdvertisedPrefixesDelta string `json:"AdvertisedPrefixes_delta"`
		AdvertisedPrefixesPrev  string `json:"AdvertisedPrefixes_prev"`
		SuppressedPrefixesDelta string `json:"SuppressedPrefixes_delta"`
		SuppressedPrefixesPrev  string `json:"SuppressedPrefixes_prev"`
		WithdrawnPrefixesDelta  string `json:"WithdrawnPrefixes_delta"`
		WithdrawnPrefixesPrev   string `json:"WithdrawnPrefixes_prev"`
		ContextName             string `json:"context_name"`
	}

	// BGPCountersResponse represents a response containing BGP counters
	BGPCountersResponse struct {
		BaseResponse
		BGPCounters []BGPCounters `json:"bgp_counters"`
	}

	// OSPFNeighbor represents an OSPF neighbor in LibreNMS
	OSPFNeighbor struct {
		DeviceID                string `json:"device_id"`
		PortID                  string `json:"port_id"`
		OSPFNbrID               string `json:"ospf_nbr_id"`
		OSPFNbrIPAddr           string `json:"ospfNbrIpAddr"`
		OSPFNbrAddressLessIndex string `json:"ospfNbrAddressLessIndex"`
		OSPFNbrRtrID            string `json:"ospfNbrRtrId"`
		OSPFNbrOptions          string `json:"ospfNbrOptions"`
		OSPFNbrPriority         string `json:"ospfNbrPriority"`
		OSPFNbrState            string `json:"ospfNbrState"`
		OSPFNbrEvents           string `json:"ospfNbrEvents"`
		OSPFNbrLsRetransQLen    string `json:"ospfNbrLsRetransQLen"`
		OSPFNbmaNbrStatus       string `json:"ospfNbmaNbrStatus"`
		OSPFNbmaNbrPermanence   string `json:"ospfNbmaNbrPermanence"`
		OSPFNbrHelloSuppressed  string `json:"ospfNbrHelloSuppressed"`
		ContextName             string `json:"context_name"`
	}

	// OSPFResponse represents a response containing OSPF neighbors
	OSPFResponse struct {
		BaseResponse
		OSPFNeighbors []OSPFNeighbor `json:"ospf_neighbours"`
	}

	// OSPFPort represents an OSPF port in LibreNMS
	OSPFPort struct {
		ID                           int     `json:"id"`
		DeviceID                     int     `json:"device_id"`
		PortID                       int     `json:"port_id"`
		OSPFPortID                   string  `json:"ospf_port_id"`
		OSPFIfIPAddress              string  `json:"ospfIfIpAddress"`
		OSPFAddressLessIf            int     `json:"ospfAddressLessIf"`
		OSPFIfAreaID                 string  `json:"ospfIfAreaId"`
		OSPFIfType                   string  `json:"ospfIfType"`
		OSPFIfAdminStat              string  `json:"ospfIfAdminStat"`
		OSPFIfRtrPriority            int     `json:"ospfIfRtrPriority"`
		OSPFIfTransitDelay           int     `json:"ospfIfTransitDelay"`
		OSPFIfRetransInterval        int     `json:"ospfIfRetransInterval"`
		OSPFIfHelloInterval          int     `json:"ospfIfHelloInterval"`
		OSPFIfRtrDeadInterval        int     `json:"ospfIfRtrDeadInterval"`
		OSPFIfPollInterval           int     `json:"ospfIfPollInterval"`
		OSPFIfState                  string  `json:"ospfIfState"`
		OSPFIfDesignatedRouter       string  `json:"ospfIfDesignatedRouter"`
		OSPFIfBackupDesignatedRouter string  `json:"ospfIfBackupDesignatedRouter"`
		OSPFIfEvents                 int     `json:"ospfIfEvents"`
		OSPFIfAuthKey                string  `json:"ospfIfAuthKey"`
		OSPFIfStatus                 string  `json:"ospfIfStatus"`
		OSPFIfMulticastForwarding    string  `json:"ospfIfMulticastForwarding"`
		OSPFIfDemand                 string  `json:"ospfIfDemand"`
		OSPFIfAuthType               int     `json:"ospfIfAuthType"`
		OSPFIfMetricIPAddress        string  `json:"ospfIfMetricIpAddress"`
		OSPFIfMetricAddressLessIf    int     `json:"ospfIfMetricAddressLessIf"`
		OSPFIfMetricTOS              int     `json:"ospfIfMetricTOS"`
		OSPFIfMetricValue            int     `json:"ospfIfMetricValue"`
		OSPFIfMetricStatus           string  `json:"ospfIfMetricStatus"`
		ContextName                  *string `json:"context_name"`
	}

	// OSPFPortsResponse represents a response containing OSPF ports
	OSPFPortsResponse struct {
		BaseResponse
		OSPFPorts []OSPFPort `json:"ospf_ports"`
	}

	// OSPFv3Neighbor represents an OSPFv3 neighbor in LibreNMS
	OSPFv3Neighbor struct {
		ID                               int    `json:"id"`
		DeviceID                         int    `json:"device_id"`
		OSPFv3InstanceID                 int    `json:"ospfv3_instance_id"`
		PortID                           int    `json:"port_id"`
		RouterID                         string `json:"router_id"`
		OSPFv3NbrIfIndex                 int    `json:"ospfv3NbrIfIndex"`
		OSPFv3NbrIfInstID                int    `json:"ospfv3NbrIfInstId"`
		OSPFv3NbrRtrID                   int    `json:"ospfv3NbrRtrId"`
		OSPFv3NbrAddressType             string `json:"ospfv3NbrAddressType"`
		OSPFv3NbrAddress                 string `json:"ospfv3NbrAddress"`
		OSPFv3NbrOptions                 int    `json:"ospfv3NbrOptions"`
		OSPFv3NbrPriority                int    `json:"ospfv3NbrPriority"`
		OSPFv3NbrState                   string `json:"ospfv3NbrState"`
		OSPFv3NbrEvents                  int    `json:"ospfv3NbrEvents"`
		OSPFv3NbrLsRetransQLen           int    `json:"ospfv3NbrLsRetransQLen"`
		OSPFv3NbrHelloSuppressed         string `json:"ospfv3NbrHelloSuppressed"`
		OSPFv3NbrIfID                    int    `json:"ospfv3NbrIfId"`
		OSPFv3NbrRestartHelperStatus     string `json:"ospfv3NbrRestartHelperStatus"`
		OSPFv3NbrRestartHelperAge        int    `json:"ospfv3NbrRestartHelperAge"`
		OSPFv3NbrRestartHelperExitReason string `json:"ospfv3NbrRestartHelperExitReason"`
		ContextName                      string `json:"context_name"`
	}

	// OSPFv3Response represents a response containing OSPFv3 neighbors
	OSPFv3Response struct {
		BaseResponse
		OSPFv3Neighbors []OSPFv3Neighbor `json:"ospfv3_neighbours"`
	}

	// OSPFv3Port represents an OSPFv3 port in LibreNMS
	OSPFv3Port struct {
		ID                                 int    `json:"id"`
		DeviceID                           int    `json:"device_id"`
		OSPFv3InstanceID                   int    `json:"ospfv3_instance_id"`
		OSPFv3AreaID                       int    `json:"ospfv3_area_id"`
		PortID                             int    `json:"port_id"`
		OSPFv3IfIndex                      int    `json:"ospfv3IfIndex"`
		OSPFv3IfInstID                     int    `json:"ospfv3IfInstId"`
		OSPFv3IfAreaID                     int    `json:"ospfv3IfAreaId"`
		OSPFv3IfType                       string `json:"ospfv3IfType"`
		OSPFv3IfAdminStatus                string `json:"ospfv3IfAdminStatus"`
		OSPFv3IfRtrPriority                int    `json:"ospfv3IfRtrPriority"`
		OSPFv3IfTransitDelay               int    `json:"ospfv3IfTransitDelay"`
		OSPFv3IfRetransInterval            int    `json:"ospfv3IfRetransInterval"`
		OSPFv3IfHelloInterval              int    `json:"ospfv3IfHelloInterval"`
		OSPFv3IfRtrDeadInterval            int    `json:"ospfv3IfRtrDeadInterval"`
		OSPFv3IfPollInterval               int    `json:"ospfv3IfPollInterval"`
		OSPFv3IfState                      string `json:"ospfv3IfState"`
		OSPFv3IfDesignatedRouter           string `json:"ospfv3IfDesignatedRouter"`
		OSPFv3IfBackupDesignatedRouter     string `json:"ospfv3IfBackupDesignatedRouter"`
		OSPFv3IfEvents                     int    `json:"ospfv3IfEvents"`
		OSPFv3IfDemand                     string `json:"ospfv3IfDemand"`
		OSPFv3IfMetricValue                int    `json:"ospfv3IfMetricValue"`
		OSPFv3IfLinkScopeLsaCount          int    `json:"ospfv3IfLinkScopeLsaCount"`
		OSPFv3IfLinkLsaCksumSum            int    `json:"ospfv3IfLinkLsaCksumSum"`
		OSPFv3IfDemandNbrProbe             string `json:"ospfv3IfDemandNbrProbe"`
		OSPFv3IfDemandNbrProbeRetransLimit int    `json:"ospfv3IfDemandNbrProbeRetransLimit"`
		OSPFv3IfDemandNbrProbeInterval     int    `json:"ospfv3IfDemandNbrProbeInterval"`
		OSPFv3IfTEDisabled                 string `json:"ospfv3IfTEDisabled"`
		OSPFv3IfLinkLSASuppression         string `json:"ospfv3IfLinkLSASuppression"`
		ContextName                        string `json:"context_name"`
	}

	// OSPFv3PortsResponse represents a response containing OSPFv3 ports
	OSPFv3PortsResponse struct {
		BaseResponse
		OSPFv3Ports []OSPFv3Port `json:"ospfv3_ports"`
	}

	// VRF represents a VRF in LibreNMS
	VRF struct {
		VRFID                        string `json:"vrf_id"`
		VRFOID                       string `json:"vrf_oid"`
		VRFName                      string `json:"vrf_name"`
		MPLSVPNVRFRouteDistinguisher string `json:"mplsVpnVrfRouteDistinguisher"`
		MPLSVPNVRFDescription        string `json:"mplsVpnVrfDescription"`
		DeviceID                     string `json:"device_id"`
	}

	// VRFResponse represents a response containing VRFs
	VRFResponse struct {
		BaseResponse
		VRFs []VRF `json:"vrfs"`
	}

	// VRFQuery represents the query parameters for filtering VRFs
	VRFQuery struct {
		Hostname string `url:"hostname,omitempty"`
		VRFName  string `url:"vrfname,omitempty"`
	}

	// MP LSService represents an MPLS service in LibreNMS
	MPLSService struct {
		SvcID                int    `json:"svc_id"`
		SvcOID               int    `json:"svc_oid"`
		DeviceID             int    `json:"device_id"`
		SvcRowStatus         string `json:"svcRowStatus"`
		SvcType              string `json:"svcType"`
		SvcCustID            int    `json:"svcCustId"`
		SvcAdminStatus       string `json:"svcAdminStatus"`
		SvcOperStatus        string `json:"svcOperStatus"`
		SvcDescription       string `json:"svcDescription"`
		SvcMtu               int    `json:"svcMtu"`
		SvcNumSaps           int    `json:"svcNumSaps"`
		SvcNumSdps           int    `json:"svcNumSdps"`
		SvcLastMgmtChange    int    `json:"svcLastMgmtChange"`
		SvcLastStatusChange  int    `json:"svcLastStatusChange"`
		SvcVRouterId         int    `json:"svcVRouterId"`
		SvcTlsMacLearning    string `json:"svcTlsMacLearning"`
		SvcTlsStpAdminStatus string `json:"svcTlsStpAdminStatus"`
		SvcTlsStpOperStatus  string `json:"svcTlsStpOperStatus"`
		SvcTlsFdbTableSize   int    `json:"svcTlsFdbTableSize"`
		SvcTlsFdbNumEntries  int    `json:"svcTlsFdbNumEntries"`
		Hostname             string `json:"hostname"`
	}

	// MPLSServicesResponse represents a response containing MPLS services
	MPLSServicesResponse struct {
		BaseResponse
		MPLSServices []MPLSService `json:"mpls_services"`
	}

	// MPLSSAP represents an MPLS SAP in LibreNMS
	MPLSSAP struct {
		SapID               int    `json:"sap_id"`
		SvcID               int    `json:"svc_id"`
		SvcOID              int    `json:"svc_oid"`
		SapPortID           int64  `json:"sapPortId"`
		IfName              string `json:"ifName"`
		DeviceID            int    `json:"device_id"`
		SapEncapValue       int    `json:"sapEncapValue"`
		SapRowStatus        string `json:"sapRowStatus"`
		SapType             string `json:"sapType"`
		SapDescription      string `json:"sapDescription"`
		SapAdminStatus      string `json:"sapAdminStatus"`
		SapOperStatus       string `json:"sapOperStatus"`
		SapLastMgmtChange   int    `json:"sapLastMgmtChange"`
		SapLastStatusChange int    `json:"sapLastStatusChange"`
		Hostname            string `json:"hostname"`
	}

	// MPLSSAPsResponse represents a response containing MPLS SAPs
	MPLSSAPsResponse struct {
		BaseResponse
		SAPs []MPLSSAP `json:"saps"`
	}

	// IPSecTunnel represents an IPSec tunnel in LibreNMS
	IPSecTunnel struct {
		TunnelID     string `json:"tunnel_id"`
		DeviceID     string `json:"device_id"`
		PeerPort     string `json:"peer_port"`
		PeerAddr     string `json:"peer_addr"`
		LocalAddr    string `json:"local_addr"`
		LocalPort    string `json:"local_port"`
		TunnelName   string `json:"tunnel_name"`
		TunnelStatus string `json:"tunnel_status"`
	}

	// IPSecResponse represents a response containing IPSec tunnels
	IPSecResponse struct {
		BaseResponse
		IPSec []IPSecTunnel `json:"ipsec"`
	}

	// IPAddress represents an IP address in LibreNMS
	IPAddress struct {
		IPv4AddressID string `json:"ipv4_address_id"`
		IPv4Address   string `json:"ipv4_address"`
		IPv4Prefixlen string `json:"ipv4_prefixlen"`
		IPv4NetworkID string `json:"ipv4_network_id"`
		PortID        string `json:"port_id"`
		ContextName   string `json:"context_name"`
	}

	// IPAddressesResponse represents a response containing IP addresses
	IPAddressesResponse struct {
		BaseResponse
		IPAddresses []IPAddress `json:"ip_addresses"`
	}

	// IPNetwork represents an IP network in LibreNMS
	IPNetwork struct {
		IPv4NetworkID string `json:"ipv4_network_id"`
		IPv4Network   string `json:"ipv4_network"`
		ContextName   string `json:"context_name"`
	}

	// IPNetworksResponse represents a response containing IP networks
	IPNetworksResponse struct {
		BaseResponse
		IPNetworks []IPNetwork `json:"ip_networks"`
	}
)

// ListBGP retrieves a list of BGP sessions from the LibreNMS API
func (r *RoutingAPI) ListBGP(query *BGPQuery) (*BGPResponse, error) {
	c := r.client
	params, err := parseParams(query)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodGet, bgpEndpoint, nil, params)
	if err != nil {
		return nil, err
	}

	bgpResp := new(BGPResponse)
	return bgpResp, c.do(req, bgpResp)
}

// GetBGP retrieves a BGP session by ID from the LibreNMS API
func (r *RoutingAPI) GetBGP(id string) (*BGPSessionResponse, error) {
	c := r.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s", bgpEndpoint, id), nil, nil)
	if err != nil {
		return nil, err
	}

	bgpResp := new(BGPSessionResponse)
	return bgpResp, c.do(req, bgpResp)
}

// UpdateBGPDescription updates the description of a BGP session
func (r *RoutingAPI) UpdateBGPDescription(id string, payload *BGPDescriptionUpdate) (*BaseResponse, error) {
	c := r.client
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/%s", bgpEndpoint, id), payload, nil)
	if err != nil {
		return nil, err
	}

	resp := new(BaseResponse)
	return resp, c.do(req, resp)
}

// ListBGPCounters retrieves a list of BGP counters from the LibreNMS API
func (r *RoutingAPI) ListBGPCounters(hostname string) (*BGPCountersResponse, error) {
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

	countersResp := new(BGPCountersResponse)
	return countersResp, c.do(req, countersResp)
}

// ListIPAddresses retrieves a list of IP addresses from the LibreNMS API
func (r *RoutingAPI) ListIPAddresses(addressFamily string) (*IPAddressesResponse, error) {
	c := r.client
	endpoint := ipAddressesEndpoint
	if addressFamily != "" {
		endpoint = fmt.Sprintf("%s/%s", ipAddressesEndpoint, addressFamily)
	}

	req, err := c.newRequest(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	addressesResp := new(IPAddressesResponse)
	return addressesResp, c.do(req, addressesResp)
}

// GetNetworkIPAddresses retrieves IP addresses for a specific network
func (r *RoutingAPI) GetNetworkIPAddresses(networkID string) (*IPAddressesResponse, error) {
	c := r.client
	endpoint := fmt.Sprintf("%s/%s/ip", ipNetworkAddressesEndpoint, networkID)

	req, err := c.newRequest(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	addressesResp := new(IPAddressesResponse)
	return addressesResp, c.do(req, addressesResp)
}

// ListIPNetworks retrieves a list of IP networks from the LibreNMS API
func (r *RoutingAPI) ListIPNetworks(addressFamily string) (*IPNetworksResponse, error) {
	c := r.client
	endpoint := ipNetworksEndpoint
	if addressFamily != "" {
		endpoint = fmt.Sprintf("%s/%s", ipNetworksEndpoint, addressFamily)
	}

	req, err := c.newRequest(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	networksResp := new(IPNetworksResponse)
	return networksResp, c.do(req, networksResp)
}

// ListIPSec retrieves a list of IPSec tunnels from the LibreNMS API
func (r *RoutingAPI) ListIPSec(hostname string) (*IPSecResponse, error) {
	c := r.client
	endpoint := fmt.Sprintf("%s/%s", ipsecEndpoint, hostname)

	req, err := c.newRequest(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	ipsecResp := new(IPSecResponse)
	return ipsecResp, c.do(req, ipsecResp)
}

// ListOSPF retrieves a list of OSPF neighbors from the LibreNMS API
func (r *RoutingAPI) ListOSPF(hostname string) (*OSPFResponse, error) {
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

	ospfResp := new(OSPFResponse)
	return ospfResp, c.do(req, ospfResp)
}

// ListOSPFPorts retrieves a list of OSPF ports from the LibreNMS API
func (r *RoutingAPI) ListOSPFPorts() (*OSPFPortsResponse, error) {
	c := r.client
	req, err := c.newRequest(http.MethodGet, ospfPortsEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	portsResp := new(OSPFPortsResponse)
	return portsResp, c.do(req, portsResp)
}

// ListOSPFv3 retrieves a list of OSPFv3 neighbors from the LibreNMS API
func (r *RoutingAPI) ListOSPFv3(hostname string) (*OSPFv3Response, error) {
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

	ospfv3Resp := new(OSPFv3Response)
	return ospfv3Resp, c.do(req, ospfv3Resp)
}

// ListOSPFv3Ports retrieves a list of OSPFv3 ports from the LibreNMS API
func (r *RoutingAPI) ListOSPFv3Ports() (*OSPFv3PortsResponse, error) {
	c := r.client
	req, err := c.newRequest(http.MethodGet, ospfv3PortsEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	portsResp := new(OSPFv3PortsResponse)
	return portsResp, c.do(req, portsResp)
}

// ListVRF retrieves a list of VRFs from the LibreNMS API
func (r *RoutingAPI) ListVRF(query *VRFQuery) (*VRFResponse, error) {
	c := r.client
	params, err := parseParams(query)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodGet, vrfEndpoint, nil, params)
	if err != nil {
		return nil, err
	}

	vrfResp := new(VRFResponse)
	return vrfResp, c.do(req, vrfResp)
}

// GetVRF retrieves a VRF by ID from the LibreNMS API
func (r *RoutingAPI) GetVRF(id string) (*VRFResponse, error) {
	c := r.client
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("%s/%s", vrfEndpoint, id), nil, nil)
	if err != nil {
		return nil, err
	}

	vrfResp := new(VRFResponse)
	return vrfResp, c.do(req, vrfResp)
}

// ListMPLSServices retrieves a list of MPLS services from the LibreNMS API
func (r *RoutingAPI) ListMPLSServices(hostname string) (*MPLSServicesResponse, error) {
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

	servicesResp := new(MPLSServicesResponse)
	return servicesResp, c.do(req, servicesResp)
}

// ListMPLSSAPs retrieves a list of MPLS SAPs from the LibreNMS API
func (r *RoutingAPI) ListMPLSSAPs(hostname string) (*MPLSSAPsResponse, error) {
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

	sapsResp := new(MPLSSAPsResponse)
	return sapsResp, c.do(req, sapsResp)
}
