package methods

import (
	"encoding/xml"
	"fmt"
	"github.com/khelaia/domingo/pkg/domingo"
	"github.com/khelaia/domingo/pkg/domingo/xmltypes"
)

func Logout(client *domingo.Client) error {
	logoutReq := &xmltypes.EPPWrapper{
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
			Logout: "\n",
		},
	}

	response, err := sendEPPRequest(client, logoutReq)
	if err != nil {
		return fmt.Errorf("logout failed: %w", err)
	}

	var eppResponse xmltypes.EPPLogoutResponse
	err = xml.Unmarshal([]byte(response), &eppResponse)
	if err != nil {
		return nil
	}

	if eppResponse.Response.Result.Code == "1500" {
		fmt.Println("Logout successful!")
		return nil
	}

	return fmt.Errorf("logout unsuccessful, response: %+v", response)
}
