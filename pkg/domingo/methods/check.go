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

	checkReq := &xmltypes.EPPWrapper{
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
			Check: (*xmltypes.CheckCommand)(&xmltypes.DomainCheck{
				DomainCheck: xmltypes.DomainCheckType{
					Names: []string{domainName},
				},
			}),
			Extension: &xmltypes.Extension{
				NamestoreExt: &xmltypes.NamestoreExtension{
					SubProduct:        "COM",
					XmlnsNamestoreExt: "http://www.verisign-grs.com/epp/namestoreExt-1.1",
				},
			},
			ClTRID: generateClTRID(),
		},
	}

	response, err := sendEPPRequest(client, checkReq)
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
