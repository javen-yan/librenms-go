package types

type (

	// SystemInfo represents system information from LibreNMS
	SystemInfo struct {
		LocalVer    string `json:"local_ver,omitempty"`
		LocalSha    string `json:"local_sha,omitempty"`
		LocalDate   string `json:"local_date,omitempty"`
		LocalBranch string `json:"local_branch,omitempty"`
		DBSchema    string `json:"db_schema,omitempty"`
		PHPVer      string `json:"php_ver,omitempty"`
		PythonVer   string `json:"python_ver,omitempty"`
		DatabaseVer string `json:"database_ver,omitempty"`
		RRDToolVer  string `json:"rrdtool_ver,omitempty"`
		NetSNMPVer  string `json:"netsnmp_ver,omitempty"`
	}

	// SystemResponse represents the response from the system API endpoint
	SystemResponse struct {
		BaseResponse
		System []SystemInfo `json:"system"`
		Count  int          `json:"count,omitempty"`
	}
)
