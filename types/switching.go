package types

type (
	VLAN struct {
		VLANID     string `json:"vlan_id,omitempty"`
		DeviceID   string `json:"device_id,omitempty"`
		VLANVLAN   string `json:"vlan_vlan,omitempty"`
		VLANDomain string `json:"vlan_domain,omitempty"`
		VLANName   string `json:"vlan_name,omitempty"`
		VLANType   string `json:"vlan_type,omitempty"`
		VLANState  int    `json:"vlan_state,omitempty"`
	}

	Link struct {
		ID             int    `json:"id,omitempty"`
		LocalPortID    int    `json:"local_port_id,omitempty"`
		LocalDeviceID  int    `json:"local_device_id,omitempty"`
		RemotePortID   int    `json:"remote_port_id,omitempty"`
		Active         int    `json:"active,omitempty"`
		Protocol       string `json:"protocol,omitempty"`
		RemoteHostname string `json:"remote_hostname,omitempty"`
		RemoteDeviceID int    `json:"remote_device_id,omitempty"`
		RemotePort     string `json:"remote_port,omitempty"`
		RemotePlatform string `json:"remote_platform,omitempty"`
		RemoteVersion  string `json:"remote_version,omitempty"`
	}

	PortFDB struct {
		PortsFDBID int    `json:"ports_fdb_id,omitempty"`
		PortID     int    `json:"port_id,omitempty"`
		MACAddress string `json:"mac_address,omitempty"`
		VLANID     int    `json:"vlan_id,omitempty"`
		DeviceID   int    `json:"device_id,omitempty"`
		CreatedAt  string `json:"created_at,omitempty"`
		UpdatedAt  string `json:"updated_at,omitempty"`
	}

	PortFDBDetail struct {
		Hostname  string `json:"hostname,omitempty"`
		SysName   string `json:"sysName,omitempty"`
		IfName    string `json:"ifName,omitempty"`
		IfAlias   string `json:"ifAlias,omitempty"`
		IfDescr   string `json:"ifDescr,omitempty"`
		LastSeen  string `json:"last_seen,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
	}

	PortNAC struct {
		PortsNACID  string `json:"ports_nac_id,omitempty"`
		AuthID      string `json:"auth_id,omitempty"`
		DeviceID    int    `json:"device_id,omitempty"`
		PortID      int    `json:"port_id,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Username    string `json:"username,omitempty"`
		MACAddress  string `json:"mac_address,omitempty"`
		IPAddress   string `json:"ip_address,omitempty"`
		HostMode    string `json:"host_mode,omitempty"`
		AuthzStatus string `json:"authz_status,omitempty"`
		AuthzBy     string `json:"authz_by,omitempty"`
		AuthcStatus string `json:"authc_status,omitempty"`
		Method      string `json:"method,omitempty"`
		Timeout     string `json:"timeout,omitempty"`
		TimeLeft    string `json:"time_left,omitempty"`
		VLAN        int    `json:"vlan,omitempty"`
		TimeElapsed string `json:"time_elapsed,omitempty"`
		CreatedAt   string `json:"created_at,omitempty"`
		UpdatedAt   string `json:"updated_at,omitempty"`
		Historical  int    `json:"historical,omitempty"`
	}

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
		MAC      string          `json:"mac,omitempty"`
		MACOUI   string          `json:"mac_oui,omitempty"`
		PortsFDB []PortFDBDetail `json:"ports_fdb"`
	}

	PortNACResponse struct {
		BaseResponse
		PortsNAC []PortNAC `json:"ports_nac"`
	}

	SwitchingQueryParams struct {
		Columns string `form:"columns,omitempty"`
		Filter  string `form:"filter,omitempty"`
	}
)
