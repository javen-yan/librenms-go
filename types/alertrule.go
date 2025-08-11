package types

type (
	// AlertRule represents an alert rule in LibreNMS.
	//
	// See https://docs.librenms.org/API/Alerts/#add_rule for field descriptions.
	AlertRule struct {
		ID           int     `json:"id"`
		Builder      string  `json:"builder"`
		Devices      []int   `json:"devices"`
		Disabled     Bool    `json:"disabled"`
		Extra        string  `json:"extra"`
		Groups       []int   `json:"groups"`
		InvertMap    Bool    `json:"invert_map"`
		Locations    []int   `json:"locations"`
		Name         string  `json:"name"`
		Notes        *string `json:"notes"`
		ProcedureURL *string `json:"proc"`
		Query        string  `json:"query"`
		Rule         string  `json:"rule"`
		Severity     string  `json:"severity"`
	}

	// AlertRuleCreateRequest is the request structure for creating an alert rule.
	//
	// See https://docs.librenms.org/API/Alerts/#add_rule for field descriptions.
	AlertRuleCreateRequest struct {
		Builder      string `json:"builder"`         // encoded JSON
		Count        int    `json:"count,omitempty"` // Max Alerts in the UI
		Delay        string `json:"delay,omitempty"`
		Devices      []int  `json:"devices"`
		Disabled     Bool   `json:"disabled,omitempty"`
		Groups       []int  `json:"groups"`
		Interval     string `json:"interval,omitempty"`
		Locations    []int  `json:"locations"`
		Mute         bool   `json:"mute,omitempty"`
		Name         string `json:"name"`
		Notes        string `json:"notes,omitempty"`
		ProcedureURL string `json:"proc,omitempty"`
		Query        string `json:"query,omitempty"`
		Rule         string `json:"rule,omitempty"`
		Severity     string `json:"severity"` // ok, warning, critical
	}

	// AlertRuleUpdateRequest is the request structure for updating an alert rule.
	AlertRuleUpdateRequest struct {
		AlertRuleCreateRequest
		ID int `json:"rule_id"`
	}

	// AlertRuleResponse is the response structure for alert rules.
	AlertRuleResponse struct {
		BaseResponse
		Rules []AlertRule `json:"rules"`
	}
)
