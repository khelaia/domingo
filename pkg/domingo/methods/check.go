package methods

import (
	"encoding/xml"
	"fmt"
	"github.com/khelaia/domingo/pkg/domingo"
	"github.com/khelaia/domingo/pkg/domingo/xmltypes"
)

type CheckDomainType struct {
	Name        string
	IsAvailable bool
	Reason      string
}

// CheckDomain sends a domain check request and returns the availability status
func CheckDomain(client *domingo.Client, domainName string) (*CheckDomainType, error) {
	req := VerisignEPPWrapperWithDefaults()
	req.Command.Check = (*xmltypes.CheckCommand)(&xmltypes.DomainCheck{
		DomainCheck: xmltypes.DomainCheckType{
			Names: []string{domainName},
		},
	})

	response, err := sendEPPRequest(client, req)
	if err != nil {
		return nil, fmt.Errorf("domain check failed: %w", err)
	}

	var eppResponse xmltypes.EPPCheckResponse
	err = xml.Unmarshal([]byte(response), &eppResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	if eppResponse.Response.Result.Code != "1000" {
		return nil, fmt.Errorf("domain check failed: %w", err)
	}

	return &CheckDomainType{
		Name:        eppResponse.Response.ResData.ChkData.Cd.Name.Text,
		IsAvailable: eppResponse.Response.ResData.ChkData.Cd.Name.Avail == "1",
		Reason:      eppResponse.Response.ResData.ChkData.Cd.Reason,
	}, nil
}
