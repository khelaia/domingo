package methods

import (
	"encoding/xml"
	"fmt"
	"github.com/khelaia/domingo/pkg/domingo"
	"github.com/khelaia/domingo/pkg/domingo/xmltypes"
)

type CreateHostResponse struct {
	Message      string
	HostName     string
	CreationDate string
}

func CreateHost(client *domingo.Client, hostName string, ipAddress string) (*CreateHostResponse, error) {
	req := VerisignEPPWrapperWithDefaults()
	req.Command.Create = &xmltypes.CreateCommand{
		CreateHost: &xmltypes.CreateHost{
			Name: hostName,
			Addr: &xmltypes.CreateHostAddr{
				Ip:   "v4",
				Text: ipAddress,
			},
		},
	}

	response, err := sendEPPRequest(client, req)
	if err != nil {
		return nil, fmt.Errorf("create host failed: %s", err)
	}

	var eppResponse xmltypes.EPPCreateHostResponse

	err = xml.Unmarshal([]byte(response), &eppResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %s", err)
	}

	if eppResponse.Response.Result.Code != "1000" {
		return nil, fmt.Errorf(eppResponse.Response.Result.Msg)
	}
	msg := "Host Created"
	return &CreateHostResponse{
		Message:      msg,
		HostName:     eppResponse.Response.ResData.CreData.Name,
		CreationDate: eppResponse.Response.ResData.CreData.CrDate,
	}, nil
}
