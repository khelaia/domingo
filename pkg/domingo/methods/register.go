package methods

import (
	"encoding/xml"
	"fmt"
	"github.com/khelaia/domingo/pkg/domingo"
	"github.com/khelaia/domingo/pkg/domingo/xmltypes"
)

type RegisterDomainType struct {
	CreationDate   string
	ExpirationDate string
	Name           string
}

// RegisterDomain is method to register domain in Registrar system
func RegisterDomain(client *domingo.Client, domainName string, authCode string, unit string, period string) (*RegisterDomainType, error) {
	req := VerisignEPPWrapperWithDefaults()
	req.Command.Create = &xmltypes.CreateCommand{
		CreateDomain: &xmltypes.RegisterDomainStruct{
			Name: domainName,
			Period: &xmltypes.RegisterDomainPeriod{
				Unit: unit,
				Text: period,
			},
			AuthInfo: &xmltypes.RegisterDomainAuthInfo{
				Pw: authCode,
			},
		},
	}

	response, err := sendEPPRequest(client, req)
	if err != nil {
		return nil, fmt.Errorf("domain register failed: %w", err)
	}

	var eppResponse xmltypes.EPPRegisterDomainResponse
	err = xml.Unmarshal([]byte(response), &eppResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	if eppResponse.Response.Result.Code != "1000" {
		return nil, fmt.Errorf(eppResponse.Response.Result.Msg)
	}

	return &RegisterDomainType{
		Name:           eppResponse.Response.ResData.CreData.Name,
		CreationDate:   eppResponse.Response.ResData.CreData.CrDate,
		ExpirationDate: eppResponse.Response.ResData.CreData.ExDate,
	}, nil
}
