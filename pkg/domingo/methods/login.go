package methods

import (
	"encoding/xml"
	"fmt"
	"github.com/khelaia/domingo/pkg/domingo"
	"github.com/khelaia/domingo/pkg/domingo/xmltypes"
)

func Login(client *domingo.Client) error {
	loginReq := &xmltypes.EPPWrapper{
		Xmlns: "urn:ietf:params:xml:ns:epp-1.0",
		Command: &xmltypes.Command{
			Login: &xmltypes.LoginCommand{
				ClientID: client.Credentials().UserID,
				Password: client.Credentials().Password,
				Options: xmltypes.Options{
					Version: "1.0",
					Lang:    "en",
				},
				Svcs: xmltypes.Svcs{
					ObjURI: []string{
						"urn:ietf:params:xml:ns:domain-1.0",
						"urn:ietf:params:xml:ns:contact-1.0",
						"urn:ietf:params:xml:ns:host-1.0",
					},
				},
			},
		},
	}

	response, err := sendEPPRequest(client, loginReq)
	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	var eppResponse xmltypes.EppLoginFullResponse
	err = xml.Unmarshal([]byte(response), &eppResponse)
	if err != nil {
		return nil
	}

	if eppResponse.Response.Result.Code == "1000" {
		fmt.Println("Login successful!")
		return nil
	}

	return fmt.Errorf("login unsuccessful, response: %+v", response)
}
