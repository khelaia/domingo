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

	createHostReq := &xmltypes.EPPWrapper{
		Xmlns:                 "urn:ietf:params:xml:ns:epp-1.0",
		XmlnsDomain:           "urn:ietf:params:xml:ns:domain-1.0",
		XmlnsContact:          "urn:ietf:params:xml:ns:contact-1.0",
		XmlnsHost:             "urn:ietf:params:xml:ns:host-1.0",
		XmlnsRegistry:         "http://www.verisign.com/epp/registry-1.0",
		XmlnsRGPPoll:          "http://www.verisign.com/epp/rgp-poll-1.0",
		XmlnsRGP:              "urn:ietf:params:xml:ns:rgp-1.0",
		XmlnsNamestore:        "http://www.verisign-grs.com/epp/namestoreExt-1.1",
		XmlnsVerificationCode: "urn:ietf:params:xml:ns:verificationCode-1.0",
		XmlnsChangePoll:       "urn:ietf:params:xml:ns:changePoll-1.0",
		XmlnsSecDNS:           "urn:ietf:params:xml:ns:secDNS-1.1",

		Command: &xmltypes.Command{
			Create: &xmltypes.CreateCommand{
				CreateHost: &xmltypes.CreateHost{
					Name: hostName,
					Addr: &xmltypes.CreateHostAddr{
						Ip:   "v4",
						Text: ipAddress,
					},
				},
			},
			Extension: &xmltypes.Extension{
				NamestoreExt: &xmltypes.NamestoreExtension{
					SubProduct:        "COM",
					XmlnsNamestoreExt: "http://www.verisign-grs.com/epp/namestoreExt-1.1",
				},
			},
			ClTRID: generateClTRID(),
		},
	}

	response, err := sendEPPRequest(client, createHostReq)
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
