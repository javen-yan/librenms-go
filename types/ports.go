package types

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

	PortsQueryParams struct {
		Columns *string `url:"columns,omitempty"`
		Filter  *string `url:"filter,omitempty"`
	}
)
