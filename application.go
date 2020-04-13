package zabbix

import (
	"fmt"
)

// Application represents a Zabbix Application returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/4.0/manual/api/reference/application/object
type Application struct {
	// HostID is the unique ID of the Host.
	HostID int

	// ApplicationID is the unique ID of the Application.
	ApplicationID int

	// Applicationname is the technical name of the Application.
	Name string

	// ApplicationFlag is the description of the Application.
	Flags string

	// LastValueType is the type of LastValue
	// 0 - float; 1 - text; 3 - int;
	Templateids []string

	// Items contains all item assigned to the host
	Items []jItem
}

type ApplicationGetParams struct {
	GetParameters

	// ItemIDs filters search results to applications with the given Item ID's.
	ItemIDs []string `json:"itemids,omitempty"`

	// GroupIDs filters search results to Application belong to the hosts
	// of the given Group ID's.
	GroupIDs []string `json:"groupids,omitempty"`

	// TemplateIDs filters search results to Application belong to the
	// given templates of the given Template ID's.
	TemplateIDs []string `json:"templateids,omitempty"`

	// HostIDs filters search results to Application belong to the
	// given Host ID's.
	HostIDs []string `json:"hostids,omitempty"`

	// ProxyIDs filters search results to Application that are
	// monitored by the given Proxy ID's.
	ProxyIDs []string `json:"proxyids,omitempty"`

	// InterfaceIDs filters search results to Application that use
	// the given host Interface ID's.
	InterfaceIDs []string `json:"interfaceids,omitempty"`

	// GraphIDs filters search results to Application that are used
	// in the given graph ID's.
	GraphIDs []string `json:"graphids,omitempty"`

	// TriggerIDs filters search results to Application that are used
	// in the given Trigger ID's.
	TriggerIDs []string `json:"triggerids,omitempty"`

	// ApplicationIDs filters search results to Application that
	// belong to the given Applications ID's.
	ApplicationIDs []string `json:"applicationids,omitempty"`

	// WebApplication flag includes web Application in the result.
	WebApplication bool `json:"webApplication,omitempty"`

	// Inherited flag return only Application inherited from a template
	// if set to 'true'.
	Inherited bool `json:"inherited,omitempty"`

	// Templated flag return only Application that belong to templates
	// if set to 'true'.
	Templated bool `json:"templated,omitempty"`

	// Monitored flag return only enabled Application that belong to
	// monitored hosts if set to 'true'.
	Monitored bool `json:"monitored,omitempty"`

	// Group filters search results to Application belong to a group
	// with the given name.
	Group string `json:"group,omitempty"`

	// Host filters search results to Application that belong to a host
	// with the given name.
	Host string `json:"host,omitempty"`

	// Application filters search results to Application that belong to
	// an application with the given name.
	Application string `json:"application,omitempty"`

	// WithTriggers flag return only Application that are used in triggers
	WithTriggers bool `json:"with_triggers,omitempty"`

	SelectItems SelectQuery `json:"selectItems,omitempty"`
}

// GetApplications queries the Zabbix API for Application matching the given search
// parameters.
//
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetApplications(params ApplicationGetParams) ([]Application, error) {
	applications := make([]jApplication, 0)
	err := c.Get("application.get", params, &applications)
	if err != nil {
		return nil, err
	}
	if len(applications) == 0 {
		return nil, ErrNotFound
	}
	// map JSON Events to Go Events
	out := make([]Application, len(applications))
	for i, japplication := range applications {
		application, err := japplication.Application()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Application %d in response: %v", i, err)
		}
		out[i] = *application
	}

	return out, nil
}
