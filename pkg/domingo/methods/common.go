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
