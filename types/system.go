package types

type (

	// SystemInfo represents system information from LibreNMS
	SystemInfo struct {
		LocalVer    string `json:"local_ver"`
		LocalSha    string `json:"local_sha"`
		LocalDate   string `json:"local_date"`
		LocalBranch string `json:"local_branch"`
		DBSchema    string `json:"db_schema"`
		PHPVer      string `json:"php_ver"`
		PythonVer   string `json:"python_ver"`
		DatabaseVer string `json:"database_ver"`
		RRDToolVer  string `json:"rrdtool_ver"`
		NetSNMPVer  string `json:"netsnmp_ver"`
	}

	// SystemResponse represents the response from the system API endpoint
	SystemResponse struct {
		BaseResponse
		System []SystemInfo `json:"system"`
		Count  int          `json:"count"`
	}
)
