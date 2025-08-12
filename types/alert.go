package types

import (
	"net/url"
	"strconv"
)

type (
	// Alert represents a LibreNMS alert.
	//
	// Pointers are used for fields that may be null.
	// A custom type Bool is used to represent booleans that may be defined as 0/1 by the API.
	Alert struct {
		ID           int    `json:"id,omitempty"`
		Alerted      Bool   `json:"alerted,omitempty"`
		DeviceID     int    `json:"device_id,omitempty"`
		Hostname     string `json:"hostname,omitempty"`
		Info         string `json:"info,omitempty"`
		Name         string `json:"name,omitempty"`
		Note         string `json:"note,omitempty"`
		Notes        string `json:"notes,omitempty"`
		Open         Bool   `json:"open,omitempty"`
		ProcedureURL string `json:"proc,omitempty"`
		RuleID       int    `json:"rule_id,omitempty"`
		Severity     string `json:"severity,omitempty"` // "ok", "warning", "critical"
		State        int    `json:"state,omitempty"`    // 0 = ok, 1 = alert, 2 = ack
		Timestamp    string `json:"timestamp,omitempty"`
	}

	// AlertAckRequest represents the request payload for acknowledging an alert.
	AlertAckRequest struct {
		Note       string `json:"note,omitempty"`
		UntilClear bool   `json:"until_clear"` // if set to false, the alert will re-alert if it gets worse/better or changes
	}

	// AlertsQuery represents the query parameters for GetAlerts().
	//
	// Documentation: https://docs.librenms.org/API/Alerts/#list_alerts
	AlertsQuery struct {
		Order    string `url:"order"`
		RuleID   int    `url:"alert_rule"`
		Severity string `url:"severity"` // "ok", "warning", "critical"
		State    int    `url:"state"`    // 0 = ok, 1 = alert, 2 = ack
	}

	// AlertsResponse represents the response from the alerts API endpoint.
	AlertsResponse struct {
		BaseResponse
		Alerts []Alert `json:"alerts"`
	}
)

// NewAlertsQuery creates a new AlertsQuery with default values.
func NewAlertsQuery() *AlertsQuery {
	return &AlertsQuery{}
}

// SetOrder sets the order for the AlertsQuery.
func (q *AlertsQuery) SetOrder(order string) *AlertsQuery {
	q.Order = order
	return q
}

// SetRuleID sets the rule ID for the AlertsQuery.
func (q *AlertsQuery) SetRuleID(ruleID int) *AlertsQuery {
	q.RuleID = ruleID
	return q
}

// SetSeverity sets the severity for the AlertsQuery.
func (q *AlertsQuery) SetSeverity(severity string) *AlertsQuery {
	q.Severity = severity
	return q
}

// SetState sets the state for the AlertsQuery.
func (q *AlertsQuery) SetState(state int) *AlertsQuery {
	q.State = state
	return q
}

// values generates the actual query payload for the request,
// only including fields that are not nil.
//
// This will allow us to send a partial list of fields and still
// be able to send 'empty' values such as 0=ok for state (avoids
// `omitempty` issues ).
func (q *AlertsQuery) Values() *url.Values {
	v := &url.Values{}
	if q.Order != "" {
		v.Set("order", q.Order)
	}
	if q.RuleID != 0 {
		v.Set("alert_rule", strconv.Itoa(q.RuleID))
	}
	if q.Severity != "" {
		v.Set("severity", q.Severity)
	}
	if q.State != 0 {
		v.Set("state", strconv.Itoa(q.State))
	}

	return v
}
