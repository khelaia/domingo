package methods

import (
	"encoding/xml"
	"fmt"
	"github.com/khelaia/domingo/pkg/domingo"
	"github.com/khelaia/domingo/pkg/domingo/xmltypes"
)

func AttachNameservers(client *domingo.Client, domainName string, hosts []string) (*string, error) {
	req := VerisignEPPWrapperWithDefaults()
	req.Command.Update = &xmltypes.UpdateCommand{
		UpdateDomainHosts: &xmltypes.UpdateDomainHosts{
			Name: domainName,
			Add: &xmltypes.AddHostToDomain{
				Ns: &xmltypes.UpdateDomainNs{
					HostObj: hosts,
				},
			},
		},
	}

	response, err := sendEPPRequest(client, req)
	if err != nil {
		return nil, fmt.Errorf("attach nameservers failed: %s", err)
	}

	var eppResponse xmltypes.EPPAttachNameserversResponse

	err = xml.Unmarshal([]byte(response), &eppResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %s", err)
	}

	if eppResponse.Response.Result.Code != "1000" {
		return nil, fmt.Errorf(eppResponse.Response.Result.ExtValue.Reason)
	}
	msg := "Nameservers added to domain"
	return &msg, nil
}
