package methods

import (
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"github.com/khelaia/domingo/pkg/domingo"
	"github.com/khelaia/domingo/pkg/domingo/xmltypes"
	"strings"
	"time"
)

// generateClTRID generates a unique client transaction ID
func generateClTRID() string {
	return fmt.Sprintf("client-%d", time.Now().UnixNano())
}

// prepareMessage adds an XML header and EPP length prefix
func prepareMessage(data []byte) ([]byte, error) {
	xmlHeader := []byte(`<?xml version="1.0" encoding="UTF-8"?>`)
	dataWithHeader := append(xmlHeader, data...)

	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(dataWithHeader)+4))
	return append(length, dataWithHeader...), nil
}

func VerisignEPPWrapperWithDefaults() *xmltypes.EPPWrapper {
	return &xmltypes.EPPWrapper{
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
			Extension: &xmltypes.Extension{
				NamestoreExt: &xmltypes.NamestoreExtension{
					SubProduct:        "COM",
					XmlnsNamestoreExt: "http://www.verisign-grs.com/epp/namestoreExt-1.1",
				},
			},
			ClTRID: generateClTRID(),
		},
	}
}

// sendEPPRequest marshals an EPP request, prepares it for sending, and handles the response
func sendEPPRequest(client *domingo.Client, request *xmltypes.EPPWrapper) (string, error) {
	request.Command.ClTRID = generateClTRID() // Set unique transaction ID

	data, err := xml.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	message, err := prepareMessage(data)

	if err != nil {
		return "", fmt.Errorf("failed to prepare message: %w", err)
	}

	if err := client.Send(message); err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	responseXML, err := client.Read()
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	startIdx := strings.Index(responseXML, "<?xml")
	if startIdx == -1 {
		return "", fmt.Errorf("no valid XML found in response")
	}

	cleanedResponse := responseXML[startIdx:]

	return cleanedResponse, nil
}
