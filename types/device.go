package types

type (
	// Device represents a device in LibreNMS.
	//
	// Pointers are used for fields that may be null.
	// A custom type Bool is used to represent booleans that may be defined as 0/1 by the API.
	Device struct {
		DeviceID int `json:"device_id,omitempty"`

		AgentUptime             int      `json:"agent_uptime,omitempty"`
		AuthAlgorithm           string   `json:"authalgo,omitempty"`
		AuthLevel               string   `json:"authlevel,omitempty"`
		AuthName                string   `json:"authname,omitempty"`
		AuthPass                string   `json:"authpass,omitempty"`
		BGPLocalAS              int      `json:"bgpLocalAs,omitempty"`
		Community               string   `json:"community,omitempty"`
		CryptoAlgorithm         string   `json:"cryptoalgo,omitempty"`
		CryptoPass              string   `json:"cryptopass,omitempty"`
		DisableNotify           Bool     `json:"disable_notify,omitempty"`
		Disabled                Bool     `json:"disabled,omitempty"`
		Display                 string   `json:"display,omitempty"`
		Features                string   `json:"features,omitempty"`
		Hardware                string   `json:"hardware,omitempty"`
		Hostname                string   `json:"hostname,omitempty"`
		Icon                    string   `json:"icon,omitempty"`
		Ignore                  Bool     `json:"ignore,omitempty"`
		IgnoreStatus            Bool     `json:"ignore_status,omitempty"`
		Inserted                string   `json:"inserted,omitempty"`
		IP                      string   `json:"ip,omitempty"`
		LastDiscovered          string   `json:"last_discovered"`
		LastDiscoveredTimeTaken float64  `json:"last_discovered_timetaken,omitempty"`
		LastPing                string   `json:"last_ping,omitempty"`
		LastPingTimeTaken       float64  `json:"last_ping_timetaken,omitempty"`
		LastPollAttempted       string   `json:"last_poll_attempted,omitempty"`
		LastPolled              string   `json:"last_pulled,omitempty"`
		LastPolledTimeTaken     float64  `json:"last_polled_timetaken,omitempty"`
		Latitude                *Float64 `json:"lat,omitempty"`
		Longitude               *Float64 `json:"lng,omitempty"`
		Location                string   `json:"location,omitempty"`
		LocationID              int      `json:"location_id,omitempty"`
		MaxDepth                int      `json:"max_depth,omitempty"`
		Notes                   string   `json:"notes,omitempty"`
		OS                      string   `json:"os,omitempty"`
		OverrideSysLocation     Bool     `json:"override_sysLocation,omitempty"`
		OverwriteIP             string   `json:"overwrite_ip,omitempty"`
		PollerGroup             int      `json:"poller_group,omitempty"`
		Port                    int      `json:"port,omitempty"`
		PortAssociationMode     int      `json:"port_association_mode,omitempty"`
		Purpose                 string   `json:"purpose,omitempty"`
		Retries                 int      `json:"retries,omitempty"`
		Serial                  string   `json:"serial,omitempty"`
		SNMPDisable             Bool     `json:"snmp_disable,omitempty"`
		SNMPVersion             string   `json:"snmpver,omitempty"`
		Status                  Bool     `json:"status,omitempty"` // /devices returns 0/1, and /devices/:id returns true/false
		StatusReason            string   `json:"status_reason,omitempty"`
		SysContact              string   `json:"sysContact,omitempty"`
		SysDescr                string   `json:"sysDescr,omitempty"`
		SysName                 string   `json:"sysName,omitempty"`
		SysObjectID             string   `json:"sysObjectID,omitempty"`
		Timeout                 int      `json:"timeout,omitempty"`
		Transport               string   `json:"transport,omitempty"`
		Type                    string   `json:"type,omitempty"`
		Uptime                  int64    `json:"uptime,omitempty"`
		Version                 string   `json:"version,omitempty"`
	}

	// DeviceCreateRequest represents the request body for creating a new device in LibreNMS.
	DeviceCreateRequest struct {
		Hostname            string `json:"hostname,omitempty"`
		Display             string `json:"display,omitempty"`
		ForceAdd            bool   `json:"force_add,omitempty"`
		Hardware            string `json:"hardware,omitempty"`
		Location            string `json:"location,omitempty"`
		LocationID          int    `json:"location_id,omitempty"`
		OS                  string `json:"os"`
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
		Field []string `json:"field,omitempty"`
		Data  []any    `json:"data,omitempty"`
	}

	// DeviceResponse represents a response containing a list of devices from the LibreNMS API.
	DeviceResponse struct {
		BaseResponse
		Devices []Device `json:"devices"`
	}

	// DeviceAvailability represents availability information for a device.
	DeviceAvailability struct {
		Duration         int     `json:"duration,omitempty"`
		AvailabilityPerc Float64 `json:"availability_perc,omitempty"`
	}

	// DeviceAvailabilityResponse represents a response containing availability information for a device.
	DeviceAvailabilityResponse struct {
		BaseResponse
		Availability []DeviceAvailability `json:"availability"`
	}

	// DeviceOutage represents an outage for a device.
	DeviceOutage struct {
		GoingDown int64 `json:"going_down,omitempty"`
		UpAgain   int64 `json:"up_again,omitempty"`
	}

	// DeviceOutagesResponse represents a response containing outages information for a device.
	DeviceOutagesResponse struct {
		BaseResponse
		Outages []DeviceOutage `json:"outages"`
	}

	// DeviceGraph represents a graph for a device.
	DeviceGraph struct {
		Desc string `json:"desc,omitempty"`
		Name string `json:"name,omitempty"`
	}

	// DeviceGraphsResponse represents a response containing graphs information for a device.
	DeviceGraphsResponse struct {
		BaseResponse
		Graphs []DeviceGraph `json:"graphs"`
	}

	// DevicePort represents a port for a device.
	DevicePort struct {
		IfName string `json:"ifName,omitempty"`
	}

	// DevicePortsResponse represents a response containing ports information for a device.
	DevicePortsResponse struct {
		BaseResponse
		Ports []DevicePort `json:"ports"`
	}

	// DeviceFDB represents a FDB entry for a device.
	DeviceFDB struct {
		PortFDBID  int    `json:"ports_fdb_id,omitempty"`
		PortID     int    `json:"port_id,omitempty"`
		MacAddress string `json:"mac_address,omitempty"`
		VlanID     int    `json:"vlan_id,omitempty"`
		DeviceID   int    `json:"device_id,omitempty"`
		CreatedAt  string `json:"created_at,omitempty"`
		UpdatedAt  string `json:"updated_at,omitempty"`
	}

	// DeviceFDBResponse represents a response containing FDB information for a device.
	DeviceFDBResponse struct {
		BaseResponse
		FDB []DeviceFDB `json:"fdb"`
	}

	// DeviceNAC represents a NAC entry for a device.
	DeviceNAC struct {
		PortNACID   int    `json:"ports_nac_id,omitempty"`
		AuthID      string `json:"auth_id,omitempty"`
		DeviceID    int    `json:"device_id,omitempty"`
		PortID      int    `json:"port_id,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Username    string `json:"username,omitempty"`
		MacAddress  string `json:"mac_address,omitempty"`
		IPAddress   string `json:"ip_address,omitempty"`
		HostMode    string `json:"host_mode,omitempty"`
		AuthzStatus string `json:"authz_status,omitempty"`
		AuthzBy     string `json:"authz_by,omitempty"`
		AuthcStatus string `json:"authc_status,omitempty"`
		Method      string `json:"method,omitempty"`
		Timeout     string `json:"timeout,omitempty"`
		TimeLeft    string `json:"time_left,omitempty"`
		Vlan        int    `json:"vlan,omitempty"`
		TimeElapsed string `json:"time_elapsed,omitempty"`
		CreatedAt   string `json:"created_at,omitempty"`
		UpdatedAt   string `json:"updated_at,omitempty"`
		Historical  int    `json:"historical,omitempty"`
	}
	// DeviceNACResponse represents a response containing NAC information for a device.
	DeviceNACResponse struct {
		BaseResponse
		NAC []DeviceNAC `json:"nac"`
	}

	// DeviceIPAddressesResponse represents a response containing IP addresses information for a device.
	DeviceIPAddressesResponse struct {
		BaseResponse
		IPAddresses []IPAddress `json:"addresses"`
	}

	// DevicePortStack represents a port stack for a device.
	DevicePortStack struct {
		DeviceID      int    `json:"device_id,omitempty"`
		PortIDHigh    int    `json:"port_id_high,omitempty"`
		PortIDLow     int    `json:"port_id_low,omitempty"`
		IfStackStatus string `json:"ifStackStatus,omitempty"`
	}

	// DevicePortStackResponse represents a response containing port stack information for a device.
	DevicePortStackResponse struct {
		BaseResponse
		PortStack []DevicePortStack `json:"mappings"`
	}

	// DeviceTransceiver represents a transceiver for a device.
	DeviceTransceiver struct {
		ID         int    `json:"id,omitempty"`
		CreatedAt  string `json:"created_at,omitempty"`
		UpdatedAt  string `json:"updated_at,omitempty"`
		DeviceID   int    `json:"device_id,omitempty"`
		PortID     int    `json:"port_id,omitempty"`
		Index      string `json:"index,omitempty"`
		Type       string `json:"type,omitempty"`
		Vendor     string `json:"vendor,omitempty"`
		OUI        string `json:"oui,omitempty"`
		Model      string `json:"model,omitempty"`
		Revision   string `json:"revision,omitempty"`
		Serial     string `json:"serial,omitempty"`
		Date       string `json:"date,omitempty"`
		DDM        bool   `json:"ddm,omitempty"`
		Encoding   string `json:"encoding,omitempty"`
		Cable      string `json:"cable,omitempty"`
		Distance   int    `json:"distance,omitempty"`
		Wavelength int    `json:"wavelength,omitempty"`
		Connector  string `json:"connector,omitempty"`
		Channels   int    `json:"channels,omitempty"`
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
		TestAttribute1 string `json:"TestAttribute-1,omitempty"`
		TestAttribute2 string `json:"TestAttribute-2,omitempty"`
		TestAttribute3 string `json:"TestAttribute-3,omitempty"`
		Type           string `json:"type,omitempty"`
		Label          string `json:"label,omitempty"`
		Status         string `json:"status,omitempty"`
		Ignore         bool   `json:"ignore,omitempty"`
		Disabled       bool   `json:"disabled,omitempty"`
	}

	// DeviceComponentsResponse represents a response containing components information for a device.
	DeviceComponentsResponse struct {
		BaseResponse
		Components map[string]DeviceComponent `json:"components,omitempty"`
	}

	// PortStats represents port stats for a device.
	PortStats struct {
		PortID     int `json:"port_id,omitempty"`
		DeviceID   int `json:"device_id,omitempty"`
		PollPrev   int `json:"poll_prev,omitempty"`
		PollPeriod int `json:"poll_period,omitempty"`
	}

	// DevicePortStatsResponse represents a response containing port stats information for a device.
	DevicePortStatsResponse struct {
		BaseResponse
		Port PortStats `json:"port,omitempty"`
	}

	// DeviceMaintenanceResponse represents a response containing maintenance information for a device.
	DeviceMaintenanceResponse struct {
		BaseResponse
		IsUnderMaintenance bool `json:"is_under_maintenance,omitempty"`
	}

	// DeviceMaintenanceRequest represents a request to set maintenance for a device.
	DeviceMaintenanceRequest struct {
		Title    string `json:"title,omitempty"`
		Notes    string `json:"notes,omitempty"`
		Start    string `json:"start,omitempty"`
		Duration string `json:"duration,omitempty"`
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
