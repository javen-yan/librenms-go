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
		ID           int     `json:"id"`
		Alerted      Bool    `json:"alerted"`
		DeviceID     int     `json:"device_id"`
		Hostname     string  `json:"hostname"`
		Info         string  `json:"info"`
		Name         string  `json:"name"`
		Note         *string `json:"note"`
		Notes        *string `json:"notes"`
		Open         Bool    `json:"open"`
		ProcedureURL *string `json:"proc"`
		RuleID       int     `json:"rule_id"`
		Severity     string  `json:"severity"` // "ok", "warning", "critical"
		State        int     `json:"state"`    // 0 = ok, 1 = alert, 2 = ack
		Timestamp    string  `json:"timestamp"`
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
		Order    *string `url:"order"`
		RuleID   *int    `url:"alert_rule"`
		Severity *string `url:"severity"` // "ok", "warning", "critical"
		State    *int    `url:"state"`    // 0 = ok, 1 = alert, 2 = ack
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
	q.Order = &order
	return q
}

// SetRuleID sets the rule ID for the AlertsQuery.
func (q *AlertsQuery) SetRuleID(ruleID int) *AlertsQuery {
	q.RuleID = &ruleID
	return q
}

// SetSeverity sets the severity for the AlertsQuery.
func (q *AlertsQuery) SetSeverity(severity string) *AlertsQuery {
	q.Severity = &severity
	return q
}

// SetState sets the state for the AlertsQuery.
func (q *AlertsQuery) SetState(state int) *AlertsQuery {
	q.State = &state
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
	if q.Order != nil {
		v.Set("order", *q.Order)
	}
	if q.RuleID != nil {
		v.Set("alert_rule", strconv.Itoa(*q.RuleID))
	}
	if q.Severity != nil {
		v.Set("severity", *q.Severity)
	}
	if q.State != nil {
		v.Set("state", strconv.Itoa(*q.State))
	}

	return v
}
