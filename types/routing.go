package types

type (
	// BGP Session represents a BGP session in LibreNMS
	BGPSession struct {
		BGPPeerID                  int    `json:"bgpPeer_id,omitempty"`
		DeviceID                   int    `json:"device_id,omitempty"`
		VRFID                      int    `json:"vrf_id,omitempty"`
		ASText                     string `json:"astext,omitempty"`
		BGPPeerIdentifier          string `json:"bgpPeerIdentifier,omitempty"`
		BGPPeerRemoteAS            int    `json:"bgpPeerRemoteAs,omitempty"`
		BGPPeerState               string `json:"bgpPeerState,omitempty"`
		BGPPeerAdminStatus         string `json:"bgpPeerAdminStatus,omitempty"`
		BGPPeerLastErrorCode       int    `json:"bgpPeerLastErrorCode,omitempty"`
		BGPPeerLastErrorSubCode    int    `json:"bgpPeerLastErrorSubCode,omitempty"`
		BGPPeerLastErrorText       string `json:"bgpPeerLastErrorText,omitempty"`
		BGPPeerIface               int    `json:"bgpPeerIface,omitempty"`
		BGPLocalAddr               string `json:"bgpLocalAddr,omitempty"`
		BGPPeerRemoteAddr          string `json:"bgpPeerRemoteAddr,omitempty"`
		BGPPeerDescr               string `json:"bgpPeerDescr,omitempty"`
		BGPPeerInUpdates           int64  `json:"bgpPeerInUpdates,omitempty"`
		BGPPeerOutUpdates          int64  `json:"bgpPeerOutUpdates,omitempty"`
		BGPPeerInTotalMessages     int64  `json:"bgpPeerInTotalMessages,omitempty"`
		BGPPeerOutTotalMessages    int64  `json:"bgpPeerOutTotalMessages,omitempty"`
		BGPPeerFsmEstablishedTime  int64  `json:"bgpPeerFsmEstablishedTime,omitempty"`
		BGPPeerInUpdateElapsedTime int64  `json:"bgpPeerInUpdateElapsedTime,omitempty"`
		ContextName                string `json:"context_name,omitempty"`
	}

	// BGPQuery represents the query parameters for filtering BGP sessions
	BGPQuery struct {
		Hostname      string `form:"hostname,omitempty"`
		ASN           int    `form:"asn,omitempty"`
		RemoteASN     int    `form:"remote_asn,omitempty"`
		RemoteAddress string `form:"remote_address,omitempty"`
		LocalAddress  string `form:"local_address,omitempty"`
		BGPDescr      string `form:"bgp_descr,omitempty"`
		BGPState      string `form:"bgp_state,omitempty"`
		BGPAdminState string `form:"bgp_adminstate,omitempty"`
		BGPFamily     int    `form:"bgp_family,omitempty"`
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
		BGPDescr string `json:"bgp_descr,omitempty"`
	}

	// BGPCounters represents BGP counters in LibreNMS
	BGPCounters struct {
		DeviceID                string `json:"device_id,omitempty"`
		BGPPeerIdentifier       string `json:"bgpPeerIdentifier"`
		AFI                     string `json:"afi,omitempty"`
		SAFI                    string `json:"safi,omitempty"`
		AcceptedPrefixes        string `json:"AcceptedPrefixes,omitempty"`
		DeniedPrefixes          string `json:"DeniedPrefixes,omitempty"`
		PrefixAdminLimit        string `json:"PrefixAdminLimit,omitempty"`
		PrefixThreshold         string `json:"PrefixThreshold,omitempty"`
		PrefixClearThreshold    string `json:"PrefixClearThreshold,omitempty"`
		AdvertisedPrefixes      string `json:"AdvertisedPrefixes,omitempty"`
		SuppressedPrefixes      string `json:"SuppressedPrefixes,omitempty"`
		WithdrawnPrefixes       string `json:"WithdrawnPrefixes,omitempty"`
		AcceptedPrefixesDelta   string `json:"AcceptedPrefixes_delta,omitempty"`
		AcceptedPrefixesPrev    string `json:"AcceptedPrefixes_prev,omitempty"`
		DeniedPrefixesDelta     string `json:"DeniedPrefixes_delta,omitempty"`
		DeniedPrefixesPrev      string `json:"DeniedPrefixes_prev,omitempty"`
		AdvertisedPrefixesDelta string `json:"AdvertisedPrefixes_delta,omitempty"`
		AdvertisedPrefixesPrev  string `json:"AdvertisedPrefixes_prev,omitempty"`
		SuppressedPrefixesDelta string `json:"SuppressedPrefixes_delta,omitempty"`
		SuppressedPrefixesPrev  string `json:"SuppressedPrefixes_prev,omitempty"`
		WithdrawnPrefixesDelta  string `json:"WithdrawnPrefixes_delta,omitempty"`
		WithdrawnPrefixesPrev   string `json:"WithdrawnPrefixes_prev,omitempty"`
		ContextName             string `json:"context_name,omitempty"`
	}

	// BGPCountersResponse represents a response containing BGP counters
	BGPCountersResponse struct {
		BaseResponse
		BGPCounters []BGPCounters `json:"bgp_counters"`
	}

	// OSPFNeighbor represents an OSPF neighbor in LibreNMS
	OSPFNeighbor struct {
		DeviceID                string `json:"device_id,omitempty"`
		PortID                  string `json:"port_id,omitempty"`
		OSPFNbrID               string `json:"ospf_nbr_id,omitempty"`
		OSPFNbrIPAddr           string `json:"ospfNbrIpAddr,omitempty"`
		OSPFNbrAddressLessIndex string `json:"ospfNbrAddressLessIndex,omitempty"`
		OSPFNbrRtrID            string `json:"ospfNbrRtrId,omitempty"`
		OSPFNbrOptions          string `json:"ospfNbrOptions,omitempty"`
		OSPFNbrPriority         string `json:"ospfNbrPriority,omitempty"`
		OSPFNbrState            string `json:"ospfNbrState,omitempty"`
		OSPFNbrEvents           string `json:"ospfNbrEvents,omitempty"`
		OSPFNbrLsRetransQLen    string `json:"ospfNbrLsRetransQLen,omitempty"`
		OSPFNbmaNbrStatus       string `json:"ospfNbmaNbrStatus,omitempty"`
		OSPFNbmaNbrPermanence   string `json:"ospfNbmaNbrPermanence,omitempty"`
		OSPFNbrHelloSuppressed  string `json:"ospfNbrHelloSuppressed,omitempty"`
		ContextName             string `json:"context_name,omitempty"`
	}

	// OSPFResponse represents a response containing OSPF neighbors
	OSPFResponse struct {
		BaseResponse
		OSPFNeighbors []OSPFNeighbor `json:"ospf_neighbours"`
	}

	// OSPFPort represents an OSPF port in LibreNMS
	OSPFPort struct {
		ID                           int    `json:"id,omitempty"`
		DeviceID                     int    `json:"device_id,omitempty"`
		PortID                       int    `json:"port_id,omitempty"`
		OSPFPortID                   string `json:"ospf_port_id,omitempty"`
		OSPFIfIPAddress              string `json:"ospfIfIpAddress,omitempty"`
		OSPFAddressLessIf            int    `json:"ospfAddressLessIf,omitempty"`
		OSPFIfAreaID                 string `json:"ospfIfAreaId,omitempty"`
		OSPFIfType                   string `json:"ospfIfType,omitempty"`
		OSPFIfAdminStat              string `json:"ospfIfAdminStat,omitempty"`
		OSPFIfRtrPriority            int    `json:"ospfIfRtrPriority,omitempty"`
		OSPFIfTransitDelay           int    `json:"ospfIfTransitDelay,omitempty"`
		OSPFIfRetransInterval        int    `json:"ospfIfRetransInterval,omitempty"`
		OSPFIfHelloInterval          int    `json:"ospfIfHelloInterval,omitempty"`
		OSPFIfRtrDeadInterval        int    `json:"ospfIfRtrDeadInterval,omitempty"`
		OSPFIfPollInterval           int    `json:"ospfIfPollInterval,omitempty"`
		OSPFIfState                  string `json:"ospfIfState,omitempty"`
		OSPFIfDesignatedRouter       string `json:"ospfIfDesignatedRouter,omitempty"`
		OSPFIfBackupDesignatedRouter string `json:"ospfIfBackupDesignatedRouter,omitempty"`
		OSPFIfEvents                 int    `json:"ospfIfEvents,omitempty"`
		OSPFIfAuthKey                string `json:"ospfIfAuthKey,omitempty"`
		OSPFIfStatus                 string `json:"ospfIfStatus,omitempty"`
		OSPFIfMulticastForwarding    string `json:"ospfIfMulticastForwarding,omitempty"`
		OSPFIfDemand                 string `json:"ospfIfDemand,omitempty"`
		OSPFIfAuthType               int    `json:"ospfIfAuthType,omitempty"`
		OSPFIfMetricIPAddress        string `json:"ospfIfMetricIpAddress,omitempty"`
		OSPFIfMetricAddressLessIf    int    `json:"ospfIfMetricAddressLessIf,omitempty"`
		OSPFIfMetricTOS              int    `json:"ospfIfMetricTOS,omitempty"`
		OSPFIfMetricValue            int    `json:"ospfIfMetricValue,omitempty"`
		OSPFIfMetricStatus           string `json:"ospfIfMetricStatus,omitempty"`
		ContextName                  string `json:"context_name,omitempty"`
	}

	// OSPFPortsResponse represents a response containing OSPF ports
	OSPFPortsResponse struct {
		BaseResponse
		OSPFPorts []OSPFPort `json:"ospf_ports"`
	}

	// OSPFv3Neighbor represents an OSPFv3 neighbor in LibreNMS
	OSPFv3Neighbor struct {
		ID                               int    `json:"id,omitempty"`
		DeviceID                         int    `json:"device_id,omitempty"`
		OSPFv3InstanceID                 int    `json:"ospfv3_instance_id,omitempty"`
		PortID                           int    `json:"port_id,omitempty"`
		RouterID                         string `json:"router_id,omitempty"`
		OSPFv3NbrIfIndex                 int    `json:"ospfv3NbrIfIndex,omitempty"`
		OSPFv3NbrIfInstID                int    `json:"ospfv3NbrIfInstId,omitempty"`
		OSPFv3NbrRtrID                   int    `json:"ospfv3NbrRtrId,omitempty"`
		OSPFv3NbrAddressType             string `json:"ospfv3NbrAddressType,omitempty"`
		OSPFv3NbrAddress                 string `json:"ospfv3NbrAddress,omitempty"`
		OSPFv3NbrOptions                 int    `json:"ospfv3NbrOptions,omitempty"`
		OSPFv3NbrPriority                int    `json:"ospfv3NbrPriority,omitempty"`
		OSPFv3NbrState                   string `json:"ospfv3NbrState,omitempty"`
		OSPFv3NbrEvents                  int    `json:"ospfv3NbrEvents,omitempty"`
		OSPFv3NbrLsRetransQLen           int    `json:"ospfv3NbrLsRetransQLen,omitempty"`
		OSPFv3NbrHelloSuppressed         string `json:"ospfv3NbrHelloSuppressed,omitempty"`
		OSPFv3NbrIfID                    int    `json:"ospfv3NbrIfId,omitempty"`
		OSPFv3NbrRestartHelperStatus     string `json:"ospfv3NbrRestartHelperStatus,omitempty"`
		OSPFv3NbrRestartHelperAge        int    `json:"ospfv3NbrRestartHelperAge,omitempty"`
		OSPFv3NbrRestartHelperExitReason string `json:"ospfv3NbrRestartHelperExitReason,omitempty"`
		ContextName                      string `json:"context_name,omitempty"`
	}

	// OSPFv3Response represents a response containing OSPFv3 neighbors
	OSPFv3Response struct {
		BaseResponse
		OSPFv3Neighbors []OSPFv3Neighbor `json:"ospfv3_neighbours"`
	}

	// OSPFv3Port represents an OSPFv3 port in LibreNMS
	OSPFv3Port struct {
		ID                                 int    `json:"id,omitempty"`
		DeviceID                           int    `json:"device_id,omitempty"`
		OSPFv3InstanceID                   int    `json:"ospfv3_instance_id,omitempty"`
		OSPFv3AreaID                       int    `json:"ospfv3_area_id,omitempty"`
		PortID                             int    `json:"port_id,omitempty"`
		OSPFv3IfIndex                      int    `json:"ospfv3IfIndex,omitempty"`
		OSPFv3IfInstID                     int    `json:"ospfv3IfInstId,omitempty"`
		OSPFv3IfAreaID                     int    `json:"ospfv3IfAreaId,omitempty"`
		OSPFv3IfType                       string `json:"ospfv3IfType,omitempty"`
		OSPFv3IfAdminStatus                string `json:"ospfv3IfAdminStatus,omitempty"`
		OSPFv3IfRtrPriority                int    `json:"ospfv3IfRtrPriority,omitempty"`
		OSPFv3IfTransitDelay               int    `json:"ospfv3IfTransitDelay,omitempty"`
		OSPFv3IfRetransInterval            int    `json:"ospfv3IfRetransInterval,omitempty"`
		OSPFv3IfHelloInterval              int    `json:"ospfv3IfHelloInterval,omitempty"`
		OSPFv3IfRtrDeadInterval            int    `json:"ospfv3IfRtrDeadInterval,omitempty"`
		OSPFv3IfPollInterval               int    `json:"ospfv3IfPollInterval,omitempty"`
		OSPFv3IfState                      string `json:"ospfv3IfState,omitempty"`
		OSPFv3IfDesignatedRouter           string `json:"ospfv3IfDesignatedRouter,omitempty"`
		OSPFv3IfBackupDesignatedRouter     string `json:"ospfv3IfBackupDesignatedRouter,omitempty"`
		OSPFv3IfEvents                     int    `json:"ospfv3IfEvents,omitempty"`
		OSPFv3IfDemand                     string `json:"ospfv3IfDemand,omitempty"`
		OSPFv3IfMetricValue                int    `json:"ospfv3IfMetricValue,omitempty"`
		OSPFv3IfLinkScopeLsaCount          int    `json:"ospfv3IfLinkScopeLsaCount,omitempty"`
		OSPFv3IfLinkLsaCksumSum            int    `json:"ospfv3IfLinkLsaCksumSum,omitempty"`
		OSPFv3IfDemandNbrProbe             string `json:"ospfv3IfDemandNbrProbe,omitempty"`
		OSPFv3IfDemandNbrProbeRetransLimit int    `json:"ospfv3IfDemandNbrProbeRetransLimit,omitempty"`
		OSPFv3IfDemandNbrProbeInterval     int    `json:"ospfv3IfDemandNbrProbeInterval,omitempty"`
		OSPFv3IfTEDisabled                 string `json:"ospfv3IfTEDisabled,omitempty"`
		OSPFv3IfLinkLSASuppression         string `json:"ospfv3IfLinkLSASuppression,omitempty"`
		ContextName                        string `json:"context_name,omitempty"`
	}

	// OSPFv3PortsResponse represents a response containing OSPFv3 ports
	OSPFv3PortsResponse struct {
		BaseResponse
		OSPFv3Ports []OSPFv3Port `json:"ospfv3_ports"`
	}

	// VRF represents a VRF in LibreNMS
	VRF struct {
		VRFID                        string `json:"vrf_id,omitempty"`
		VRFOID                       string `json:"vrf_oid,omitempty"`
		VRFName                      string `json:"vrf_name,omitempty"`
		MPLSVPNVRFRouteDistinguisher string `json:"mplsVpnVrfRouteDistinguisher,omitempty"`
		MPLSVPNVRFDescription        string `json:"mplsVpnVrfDescription,omitempty"`
		DeviceID                     string `json:"device_id,omitempty"`
	}

	// VRFResponse represents a response containing VRFs
	VRFResponse struct {
		BaseResponse
		VRFs []VRF `json:"vrfs"`
	}

	// VRFQuery represents the query parameters for filtering VRFs
	VRFQuery struct {
		Hostname string `form:"hostname,omitempty"`
		VRFName  string `form:"vrfname,omitempty"`
	}

	// MP LSService represents an MPLS service in LibreNMS
	MPLSService struct {
		SvcID                int    `json:"svc_id,omitempty"`
		SvcOID               int    `json:"svc_oid,omitempty"`
		DeviceID             int    `json:"device_id,omitempty"`
		SvcRowStatus         string `json:"svcRowStatus,omitempty"`
		SvcType              string `json:"svcType,omitempty"`
		SvcCustID            int    `json:"svcCustId,omitempty"`
		SvcAdminStatus       string `json:"svcAdminStatus,omitempty"`
		SvcOperStatus        string `json:"svcOperStatus,omitempty"`
		SvcDescription       string `json:"svcDescription,omitempty"`
		SvcMtu               int    `json:"svcMtu,omitempty"`
		SvcNumSaps           int    `json:"svcNumSaps,omitempty"`
		SvcNumSdps           int    `json:"svcNumSdps,omitempty"`
		SvcLastMgmtChange    int    `json:"svcLastMgmtChange,omitempty"`
		SvcLastStatusChange  int    `json:"svcLastStatusChange,omitempty"`
		SvcVRouterId         int    `json:"svcVRouterId,omitempty"`
		SvcTlsMacLearning    string `json:"svcTlsMacLearning,omitempty"`
		SvcTlsStpAdminStatus string `json:"svcTlsStpAdminStatus,omitempty"`
		SvcTlsStpOperStatus  string `json:"svcTlsStpOperStatus,omitempty"`
		SvcTlsFdbTableSize   int    `json:"svcTlsFdbTableSize,omitempty"`
		SvcTlsFdbNumEntries  int    `json:"svcTlsFdbNumEntries,omitempty"`
		Hostname             string `json:"hostname,omitempty"`
	}

	// MPLSServicesResponse represents a response containing MPLS services
	MPLSServicesResponse struct {
		BaseResponse
		MPLSServices []MPLSService `json:"mpls_services"`
	}

	// MPLSSAP represents an MPLS SAP in LibreNMS
	MPLSSAP struct {
		SapID               int    `json:"sap_id,omitempty"`
		SvcID               int    `json:"svc_id,omitempty"`
		SvcOID              int    `json:"svc_oid,omitempty"`
		SapPortID           int64  `json:"sapPortId,omitempty"`
		IfName              string `json:"ifName,omitempty"`
		DeviceID            int    `json:"device_id,omitempty"`
		SapEncapValue       int    `json:"sapEncapValue,omitempty"`
		SapRowStatus        string `json:"sapRowStatus,omitempty"`
		SapType             string `json:"sapType,omitempty"`
		SapDescription      string `json:"sapDescription,omitempty"`
		SapAdminStatus      string `json:"sapAdminStatus,omitempty"`
		SapOperStatus       string `json:"sapOperStatus,omitempty"`
		SapLastMgmtChange   int    `json:"sapLastMgmtChange,omitempty"`
		SapLastStatusChange int    `json:"sapLastStatusChange,omitempty"`
		Hostname            string `json:"hostname,omitempty"`
	}

	// MPLSSAPsResponse represents a response containing MPLS SAPs
	MPLSSAPsResponse struct {
		BaseResponse
		SAPs []MPLSSAP `json:"saps"`
	}

	// IPSecTunnel represents an IPSec tunnel in LibreNMS
	IPSecTunnel struct {
		TunnelID     string `json:"tunnel_id,omitempty"`
		DeviceID     string `json:"device_id,omitempty"`
		PeerPort     string `json:"peer_port,omitempty"`
		PeerAddr     string `json:"peer_addr,omitempty"`
		LocalAddr    string `json:"local_addr,omitempty"`
		LocalPort    string `json:"local_port,omitempty"`
		TunnelName   string `json:"tunnel_name,omitempty"`
		TunnelStatus string `json:"tunnel_status,omitempty"`
	}

	// IPSecResponse represents a response containing IPSec tunnels
	IPSecResponse struct {
		BaseResponse
		IPSec []IPSecTunnel `json:"ipsec"`
	}

	// IPAddress represents an IP address in LibreNMS
	IPAddress struct {
		IPv4AddressID int    `json:"ipv4_address_id,omitempty"`
		IPv4Address   string `json:"ipv4_address,omitempty"`
		IPv4Prefixlen int    `json:"ipv4_prefixlen,omitempty"`
		IPv4NetworkID int    `json:"ipv4_network_id,omitempty"`
		PortID        int    `json:"port_id,omitempty"`
		ContextName   string `json:"context_name,omitempty"`
	}

	// IPAddressesResponse represents a response containing IP addresses
	IPAddressesResponse struct {
		BaseResponse
		IPAddresses []IPAddress `json:"ip_addresses"`
	}

	// IPNetwork represents an IP network in LibreNMS
	IPNetwork struct {
		IPv4NetworkID int    `json:"ipv4_network_id,omitempty"`
		IPv4Network   string `json:"ipv4_network,omitempty"`
		ContextName   string `json:"context_name,omitempty"`
	}

	// IPNetworksResponse represents a response containing IP networks
	IPNetworksResponse struct {
		BaseResponse
		IPNetworks []IPNetwork `json:"ip_networks"`
	}
)
