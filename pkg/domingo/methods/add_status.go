package methods

import (
	"encoding/xml"
	"fmt"
	"github.com/khelaia/domingo/pkg/domingo"
	"github.com/khelaia/domingo/pkg/domingo/constants"
	"github.com/khelaia/domingo/pkg/domingo/xmltypes"
)

// AddStatuses Add Statuses to Domain
func AddStatuses(client *domingo.Client, domainName string, statuses []constants.ClientStatus) (*string, error) {
	req := VerisignEPPWrapperWithDefaults()

	var addStatuses []xmltypes.AddStatus
	for _, status := range statuses {
		addStatuses = append(addStatuses, xmltypes.AddStatus{S: status})
	}
	req.Command.Update = &xmltypes.UpdateCommand{
		UpdateDomain: &xmltypes.UpdateDomain{
			Name: domainName,
			Add: &xmltypes.AddToDomain{
				Statuses: &addStatuses,
			},
		},
	}

	response, err := sendEPPRequest(client, req)
	if err != nil {
		return nil, fmt.Errorf("failed add domain statuses: %s", err)
	}

	var eppResponse xmltypes.EPPAddStatusResponse
	err = xml.Unmarshal([]byte(response), &eppResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	if eppResponse.Response.Result.Code != "1000" {
		return nil, fmt.Errorf(eppResponse.Response.Result.ExtValue.Reason)
	}

	return &eppResponse.Response.Result.Msg, nil
}
