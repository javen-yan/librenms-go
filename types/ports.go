package types

type (
	Port struct {
		PortID                  int    `json:"port_id,omitempty"`
		DeviceID                int    `json:"device_id,omitempty"`
		PortDescrType           string `json:"port_descr_type,omitempty"`
		PortDescrDescr          string `json:"port_descr_descr,omitempty"`
		PortDescrCircuit        string `json:"port_descr_circuit,omitempty"`
		PortDescrSpeed          string `json:"port_descr_speed,omitempty"`
		PortDescrNotes          string `json:"port_descr_notes,omitempty"`
		IfDescr                 string `json:"ifDescr,omitempty"`
		IfName                  string `json:"ifName,omitempty"`
		PortName                string `json:"portName,omitempty"`
		IfIndex                 string `json:"ifIndex,omitempty"`
		IfSpeed                 string `json:"ifSpeed,omitempty"`
		IfConnectorPresent      string `json:"ifConnectorPresent,omitempty"`
		IfPromiscuousMode       string `json:"ifPromiscuousMode,omitempty"`
		IfHighSpeed             string `json:"ifHighSpeed,omitempty"`
		IfOperStatus            string `json:"ifOperStatus,omitempty"`
		IfOperStatusPrev        string `json:"ifOperStatus_prev,omitempty"`
		IfAdminStatus           string `json:"ifAdminStatus,omitempty"`
		IfAdminStatusPrev       string `json:"ifAdminStatus_prev,omitempty"`
		IfDuplex                string `json:"ifDuplex,omitempty"`
		IfMtu                   string `json:"ifMtu,omitempty"`
		IfType                  string `json:"ifType,omitempty"`
		IfAlias                 string `json:"ifAlias,omitempty"`
		IfPhysAddress           string `json:"ifPhysAddress,omitempty"`
		IfHardType              string `json:"ifHardType,omitempty"`
		IfLastChange            string `json:"ifLastChange,omitempty"`
		IfVlan                  string `json:"ifVlan,omitempty"`
		IfTrunk                 string `json:"ifTrunk,omitempty"`
		IfVrf                   string `json:"ifVrf,omitempty"`
		CounterIn               string `json:"counter_in,omitempty"`
		CounterOut              string `json:"counter_out,omitempty"`
		Ignore                  string `json:"ignore,omitempty"`
		Disabled                string `json:"disabled,omitempty"`
		Detailed                string `json:"detailed,omitempty"`
		Deleted                 string `json:"deleted,omitempty"`
		PagpOperationMode       string `json:"pagpOperationMode,omitempty"`
		PagpPortState           string `json:"pagpPortState,omitempty"`
		PagpPartnerDeviceId     string `json:"pagpPartnerDeviceId,omitempty"`
		PagpPartnerLearnMethod  string `json:"pagpPartnerLearnMethod,omitempty"`
		PagpPartnerIfIndex      string `json:"pagpPartnerIfIndex,omitempty"`
		PagpPartnerGroupIfIndex string `json:"pagpPartnerGroupIfIndex,omitempty"`
		PagpPartnerDeviceName   string `json:"pagpPartnerDeviceName,omitempty"`
		PagpEthcOperationMode   string `json:"pagpEthcOperationMode,omitempty"`
		PagpDeviceId            string `json:"pagpDeviceId,omitempty"`
		PagpGroupIfIndex        string `json:"pagpGroupIfIndex,omitempty"`
		IfInUcastPkts           string `json:"ifInUcastPkts,omitempty"`
		IfInUcastPktsPrev       string `json:"ifInUcastPkts_prev,omitempty"`
		IfInUcastPktsDelta      string `json:"ifInUcastPkts_delta,omitempty"`
		IfInUcastPktsRate       string `json:"ifInUcastPkts_rate,omitempty"`
		IfOutUcastPkts          string `json:"ifOutUcastPkts,omitempty"`
		IfOutUcastPktsPrev      string `json:"ifOutUcastPkts_prev,omitempty"`
		IfOutUcastPktsDelta     string `json:"ifOutUcastPkts_delta,omitempty"`
		IfOutUcastPktsRate      string `json:"ifOutUcastPkts_rate,omitempty"`
		IfInErrors              string `json:"ifInErrors,omitempty"`
		IfInErrorsPrev          string `json:"ifInErrors_prev,omitempty"`
		IfInErrorsDelta         string `json:"ifInErrors_delta,omitempty"`
		IfInErrorsRate          string `json:"ifInErrors_rate,omitempty"`
		IfOutErrors             string `json:"ifOutErrors,omitempty"`
		IfOutErrorsPrev         string `json:"ifOutErrors_prev,omitempty"`
		IfOutErrorsDelta        string `json:"ifOutErrors_delta,omitempty"`
		IfOutErrorsRate         string `json:"ifOutErrors_rate,omitempty"`
		IfInOctets              string `json:"ifInOctets,omitempty"`
		IfInOctetsPrev          string `json:"ifInOctets_prev,omitempty"`
		IfInOctetsDelta         string `json:"ifInOctets_delta,omitempty"`
		IfInOctetsRate          string `json:"ifInOctets_rate,omitempty"`
		IfOutOctets             string `json:"ifOutOctets,omitempty"`
		IfOutOctetsPrev         string `json:"ifOutOctets_prev,omitempty"`
		IfOutOctetsDelta        string `json:"ifOutOctets_delta,omitempty"`
		IfOutOctetsRate         string `json:"ifOutOctets_rate,omitempty"`
		PollTime                string `json:"poll_time,omitempty"`
		PollPrev                string `json:"poll_prev,omitempty"`
		PollPeriod              string `json:"poll_period,omitempty"`
	}

	PortTransceiver struct {
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

	PortDescriptionUpdateRequest struct {
		Description string `json:"description,omitempty"`
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
		Addresses []IPAddress `json:"addresses"`
	}

	PortTransceiverResponse struct {
		BaseResponse
		Transceivers []PortTransceiver `json:"transceivers"`
	}

	PortDescriptionResponse struct {
		BaseResponse
		PortDescription string `json:"port_description,omitempty"`
	}

	PortsQueryParams struct {
		Columns string `url:"columns,omitempty"`
		Filter  string `url:"filter,omitempty"`
	}
)
