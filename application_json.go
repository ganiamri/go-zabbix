package zabbix

import (
	"fmt"
	"strconv"
)

// jApplication is a private map for the Zabbix API Host object.
// See: https://www.zabbix.com/documentation/4.0/manual/api/reference/application/get
type jApplication struct {
	HostID              string   `json:"hostid,omitempty"`
	ApplicationID       string   `json:"applicationid"`
	ApplicationName     string   `json:"name"`
	ApplicationFlag     string   `json:"flag"`
	ApplicationTemplate []string `json:"templateids"`
	ApplicationItems    []jItem  `json:"items,omitempty"`
}

// Application returns a native Go Application struct mapped from the given JSON Application data.
func (c *jApplication) Application() (*Application, error) {
	var err error
	application := &Application{}
	application.HostID, err = strconv.Atoi(c.HostID)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Host ID: %v", err)
	}
	application.ApplicationID, err = strconv.Atoi(c.ApplicationID)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Application ID: %v", err)
	}
	application.Name = c.ApplicationName
	application.Flags = c.ApplicationFlag
	application.Templateids = c.ApplicationTemplate
	application.Items = c.ApplicationItems

	return application, err
}

// jApplications is a slice of jApplications structs.
type jApplications []jApplication

// Applications returns a native Go slice of applications mapped from the given JSON ApplicationS
// data.
func (c jApplications) Applications() ([]Application, error) {
	if c != nil {
		applications := make([]Application, len(c))
		for i, jApplication := range c {
			application, err := jApplication.Application()
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling application %d in JSON data: %v", i, err)
			}
			applications[i] = *application
		}

		return applications, nil
	}

	return nil, nil
}
