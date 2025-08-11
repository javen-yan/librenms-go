package types

type (
	// Device represents a device in LibreNMS.
	//
	// Pointers are used for fields that may be null.
	// A custom type Bool is used to represent booleans that may be defined as 0/1 by the API.
	Device struct {
		DeviceID int `json:"device_id"`

		AgentUptime             int      `json:"agent_uptime"`
		AuthAlgorithm           *string  `json:"authalgo"`
		AuthLevel               *string  `json:"authlevel"`
		AuthName                *string  `json:"authname"`
		AuthPass                *string  `json:"authpass"`
		BGPLocalAS              *int     `json:"bgpLocalAs"`
		Community               *string  `json:"community"`
		CryptoAlgorithm         *string  `json:"cryptoalgo"`
		CryptoPass              *string  `json:"cryptopass"`
		DisableNotify           Bool     `json:"disable_notify"`
		Disabled                Bool     `json:"disabled"`
		Display                 *string  `json:"display"`
		Features                *string  `json:"features"`
		Hardware                string   `json:"hardware"`
		Hostname                string   `json:"hostname"`
		Icon                    string   `json:"icon"`
		Ignore                  Bool     `json:"ignore"`
		IgnoreStatus            Bool     `json:"ignore_status"`
		Inserted                string   `json:"inserted"`
		IP                      string   `json:"ip"`
		LastDiscovered          *string  `json:"last_discovered"`
		LastDiscoveredTimeTaken float64  `json:"last_discovered_timetaken"`
		LastPing                *string  `json:"last_ping"`
		LastPingTimeTaken       float64  `json:"last_ping_timetaken"`
		LastPollAttempted       *string  `json:"last_poll_attempted"`
		LastPolled              *string  `json:"last_pulled"`
		LastPolledTimeTaken     float64  `json:"last_polled_timetaken"`
		Latitude                *Float64 `json:"lat"`
		Longitude               *Float64 `json:"lng"`
		Location                *string  `json:"location"`
		LocationID              *int     `json:"location_id"`
		MaxDepth                *int     `json:"max_depth"`
		Notes                   *string  `json:"notes"`
		OS                      string   `json:"os"`
		OverrideSysLocation     Bool     `json:"override_sysLocation"`
		OverwriteIP             string   `json:"overwrite_ip"`
		PollerGroup             int      `json:"poller_group"`
		Port                    int      `json:"port"`
		PortAssociationMode     int      `json:"port_association_mode"`
		Purpose                 *string  `json:"purpose"`
		Retries                 *int     `json:"retries"`
		Serial                  *string  `json:"serial"`
		SNMPDisable             Bool     `json:"snmp_disable"`
		SNMPVersion             string   `json:"snmpver"`
		Status                  Bool     `json:"status"` // /devices returns 0/1, and /devices/:id returns true/false
		StatusReason            string   `json:"status_reason"`
		SysContact              *string  `json:"sysContact"`
		SysDescr                *string  `json:"sysDescr"`
		SysName                 string   `json:"sysName"`
		SysObjectID             *string  `json:"sysObjectID"`
		Timeout                 *int     `json:"timeout"`
		Transport               string   `json:"transport"`
		Type                    string   `json:"type"`
		Uptime                  *int64   `json:"uptime"`
		Version                 *string  `json:"version"`
	}

	// DeviceCreateRequest represents the request body for creating a new device in LibreNMS.
	DeviceCreateRequest struct {
		Hostname            string `json:"hostname"`
		Display             string `json:"display,omitempty"`
		ForceAdd            bool   `json:"force_add,omitempty"`
		Hardware            string `json:"hardware,omitempty"`
		Location            string `json:"location,omitempty"`
		LocationID          int    `json:"location_id,omitempty"`
		OS                  string `json:"os,omitempty"`
		OverrideSysLocation bool   `json:"override_sysLocation,omitempty"`
		PingFallback        bool   `json:"ping_fallback,omitempty"`
		PollerGroup         int    `json:"poller_group,omitempty"`
		Port                int    `json:"port,omitempty"`
		PortAssocMode       int    `json:"port_association_mode,omitempty"` // ifIndex(1), ifName(2), ifDescr(3), ifAlias(4)
		SNMPAuthAlgo        string `json:"authalgo,omitempty"`              // MD5, SHA, SHA-224, SHA-256, SHA384, SHA-512
		SNMPAuthLevel       string `json:"authlevel,omitempty"`             // noAuthNoPriv, authNoPriv, authPriv
		SNMPAuthName        string `json:"authname,omitempty"`
		SNMPAuthPass        string `json:"authpass,omitempty"`
		SNMPCrytoAlgo       string `json:"cryptoalgo,omitempty"` // DES, AES, AES-192, AES-256, AES-256-C
		SNMPCryptoPass      string `json:"cryptopass,omitempty"`
		SNMPCommunity       string `json:"community,omitempty"`
		SNMPDisable         bool   `json:"snmp_disable,omitempty"`
		SNMPVersion         string `json:"snmpver,omitempty"` // v1, v2c, v3
		SysName             string `json:"sysName,omitempty"`
		Transport           string `json:"transport,omitempty"`
	}

	// DeviceUpdateRequest represents the request body for updating a device in LibreNMS.
	//
	// The `Field` slice contains the names of the field(s) to update,
	// and `Data` contains the corresponding values. Only specify the fields you want to update.
	DeviceUpdateRequest struct {
		Field []string `json:"field"`
		Data  []any    `json:"data"`
	}

	// DeviceResponse represents a response containing a list of devices from the LibreNMS API.
	DeviceResponse struct {
		BaseResponse
		Devices []Device `json:"devices"`
	}

	// DeviceAvailability represents availability information for a device.
	DeviceAvailability struct {
		Duration         int     `json:"duration"`
		AvailabilityPerc Float64 `json:"availability_perc"`
	}

	// DeviceAvailabilityResponse represents a response containing availability information for a device.
	DeviceAvailabilityResponse struct {
		BaseResponse
		Availability []DeviceAvailability `json:"availability"`
	}

	// DeviceOutage represents an outage for a device.
	DeviceOutage struct {
		GoingDown int64 `json:"going_down"`
		UpAgain   int64 `json:"up_again"`
	}

	// DeviceOutagesResponse represents a response containing outages information for a device.
	DeviceOutagesResponse struct {
		BaseResponse
		Outages []DeviceOutage `json:"outages"`
	}

	// DeviceGraph represents a graph for a device.
	DeviceGraph struct {
		Desc string `json:"desc"`
		Name string `json:"name"`
	}

	// DeviceGraphsResponse represents a response containing graphs information for a device.
	DeviceGraphsResponse struct {
		BaseResponse
		Graphs []DeviceGraph `json:"graphs"`
	}

	// DevicePort represents a port for a device.
	DevicePort struct {
		IfName string `json:"ifName"`
	}

	// DevicePortsResponse represents a response containing ports information for a device.
	DevicePortsResponse struct {
		BaseResponse
		Ports []DevicePort `json:"ports"`
	}

	// DeviceFDB represents a FDB entry for a device.
	DeviceFDB struct {
		PortFDBID  int    `json:"ports_fdb_id"`
		PortID     int    `json:"port_id"`
		MacAddress string `json:"mac_address"`
		VlanID     int    `json:"vlan_id"`
		DeviceID   int    `json:"device_id"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	}

	// DeviceFDBResponse represents a response containing FDB information for a device.
	DeviceFDBResponse struct {
		BaseResponse
		FDB []DeviceFDB `json:"fdb"`
	}

	// DeviceNAC represents a NAC entry for a device.
	DeviceNAC struct {
		PortNACID   int     `json:"ports_nac_id"`
		AuthID      string  `json:"auth_id"`
		DeviceID    int     `json:"device_id"`
		PortID      int     `json:"port_id"`
		Domain      string  `json:"domain"`
		Username    string  `json:"username"`
		MacAddress  string  `json:"mac_address"`
		IPAddress   string  `json:"ip_address"`
		HostMode    string  `json:"host_mode"`
		AuthzStatus string  `json:"authz_status"`
		AuthzBy     string  `json:"authz_by"`
		AuthcStatus string  `json:"authc_status"`
		Method      string  `json:"method"`
		Timeout     string  `json:"timeout"`
		TimeLeft    string  `json:"time_left"`
		Vlan        int     `json:"vlan"`
		TimeElapsed *string `json:"time_elapsed"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
		Historical  int     `json:"historical"`
	}
	// DeviceNACResponse represents a response containing NAC information for a device.
	DeviceNACResponse struct {
		BaseResponse
		NAC []DeviceNAC `json:"nac"`
	}

	// DeviceIPAddress represents an IP address for a device.
	DeviceIPAddress struct {
		IPv4AddressID int    `json:"ipv4_address_id"`
		IPv4Address   string `json:"ipv4_address"`
		IPv4Prefixlen int    `json:"ipv4_prefixlen"`
		IPv4NetworkID int    `json:"ipv4_network_id"`
		PortID        int    `json:"port_id"`
		ContextName   string `json:"context_name"`
	}

	// DeviceIPAddressesResponse represents a response containing IP addresses information for a device.
	DeviceIPAddressesResponse struct {
		BaseResponse
		IPAddresses []DeviceIPAddress `json:"addresses"`
	}

	// DevicePortStack represents a port stack for a device.
	DevicePortStack struct {
		DeviceID      int    `json:"device_id"`
		PortIDHigh    int    `json:"port_id_high"`
		PortIDLow     int    `json:"port_id_low"`
		IfStackStatus string `json:"ifStackStatus"`
	}

	// DevicePortStackResponse represents a response containing port stack information for a device.
	DevicePortStackResponse struct {
		BaseResponse
		PortStack []DevicePortStack `json:"mappings"`
	}

	// DeviceTransceiver represents a transceiver for a device.
	DeviceTransceiver struct {
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

	// DeviceTransceiversResponse represents a response containing transceivers information for a device.
	DeviceTransceiversResponse struct {
		BaseResponse
		Transceivers []DeviceTransceiver `json:"transceivers"`
	}

	// ComponentsQuery represents the query parameters for filtering GetComponents().
	ComponentsQuery struct {
		Type     string `url:"type,omitempty"`
		ID       int    `url:"id,omitempty"`
		Label    string `url:"label,omitempty"`
		Status   string `url:"status,omitempty"`
		Disabled bool   `url:"disabled,omitempty"`
		Ignore   bool   `url:"ignore,omitempty"`
	}

	// DeviceComponent represents a component for a device.
	DeviceComponent struct {
		TestAttribute1 string `json:"TestAttribute-1"`
		TestAttribute2 string `json:"TestAttribute-2"`
		TestAttribute3 string `json:"TestAttribute-3"`
		Type           string `json:"type"`
		Label          string `json:"label"`
		Status         string `json:"status"`
		Ignore         bool   `json:"ignore"`
		Disabled       bool   `json:"disabled"`
	}

	// DeviceComponentsResponse represents a response containing components information for a device.
	DeviceComponentsResponse struct {
		BaseResponse
		Components map[string]DeviceComponent `json:"components"`
	}

	// PortStats represents port stats for a device.
	PortStats struct {
		PortID     int `json:"port_id"`
		DeviceID   int `json:"device_id"`
		PollPrev   int `json:"poll_prev"`
		PollPeriod int `json:"poll_period"`
	}

	// DevicePortStatsResponse represents a response containing port stats information for a device.
	DevicePortStatsResponse struct {
		BaseResponse
		Port PortStats `json:"port"`
	}

	// DeviceMaintenanceResponse represents a response containing maintenance information for a device.
	DeviceMaintenanceResponse struct {
		BaseResponse
		IsUnderMaintenance bool `json:"is_under_maintenance"`
	}

	// DeviceMaintenanceRequest represents a request to set maintenance for a device.
	DeviceMaintenanceRequest struct {
		Title    string `json:"title,omitempty"`
		Notes    string `json:"notes,omitempty"`
		Start    string `json:"start,omitempty"`
		Duration string `json:"duration"`
	}

	// DeviceGroupsResponse represents a response containing groups information for a device.
	DeviceGroupsResponse struct {
		BaseResponse
		Groups []DeviceGroup `json:"groups"`
	}

	// DevicesQuery represents the query parameters for filtering GetDevices().
	DevicesQuery struct {
		DeviceID   int    `url:"device_id,omitempty"`
		Display    string `url:"display,omitempty"`
		Hostname   string `url:"hostname,omitempty"`
		IPv4       string `url:"ipv4,omitempty"`
		IPv6       string `url:"ipv6,omitempty"`
		Location   string `url:"location,omitempty"`
		LocationID int    `url:"location_id,omitempty"`
		MACAddress string `url:"mac,omitempty"`
		Order      string `url:"order,omitempty"`
		OS         string `url:"os,omitempty"`
		SysName    string `url:"sysName,omitempty"`
		Type       string `url:"type,omitempty"`
	}
)
