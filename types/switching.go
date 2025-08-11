package types

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
		ID             int     `json:"id"`
		LocalPortID    int     `json:"local_port_id"`
		LocalDeviceID  int     `json:"local_device_id"`
		RemotePortID   int     `json:"remote_port_id"`
		Active         int     `json:"active"`
		Protocol       string  `json:"protocol"`
		RemoteHostname string  `json:"remote_hostname"`
		RemoteDeviceID int     `json:"remote_device_id"`
		RemotePort     string  `json:"remote_port"`
		RemotePlatform *string `json:"remote_platform"`
		RemoteVersion  string  `json:"remote_version"`
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
		PortsNACID  string  `json:"ports_nac_id"`
		AuthID      string  `json:"auth_id"`
		DeviceID    int     `json:"device_id"`
		PortID      int     `json:"port_id"`
		Domain      string  `json:"domain"`
		Username    string  `json:"username"`
		MACAddress  string  `json:"mac_address"`
		IPAddress   string  `json:"ip_address"`
		HostMode    string  `json:"host_mode"`
		AuthzStatus string  `json:"authz_status"`
		AuthzBy     string  `json:"authz_by"`
		AuthcStatus string  `json:"authc_status"`
		Method      string  `json:"method"`
		Timeout     string  `json:"timeout"`
		TimeLeft    string  `json:"time_left"`
		VLAN        int     `json:"vlan"`
		TimeElapsed *string `json:"time_elapsed"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
		Historical  int     `json:"historical"`
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
		MAC      string          `json:"mac"`
		MACOUI   string          `json:"mac_oui"`
		PortsFDB []PortFDBDetail `json:"ports_fdb"`
	}

	PortNACResponse struct {
		BaseResponse
		PortsNAC []PortNAC `json:"ports_nac"`
	}

	SwitchingQueryParams struct {
		Columns *string `url:"columns,omitempty"`
		Filter  *string `url:"filter,omitempty"`
	}
)
