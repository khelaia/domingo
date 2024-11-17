package methods

import (
	"encoding/xml"
	"fmt"
	"github.com/khelaia/domingo/pkg/domingo"
	"github.com/khelaia/domingo/pkg/domingo/xmltypes"
	"time"
)

type Status struct {
	Status string
}

// AuthInfo represents the authentication information for a domain.
type AuthInfo struct {
	Password string
}

type DomainInfoType struct {
	Name             string
	RegistryObjectID string
	Statuses         []Status
	NameServers      []string
	Hosts            []string
	ClientID         string
	CreatorID        string
	CreationDate     time.Time
	UpdaterID        string
	UpdateDate       time.Time
	ExpirationDate   time.Time
	AuthInfo         AuthInfo
}

// DomainInfo get all info about domain
func DomainInfo(client *domingo.Client, domainName string) (*DomainInfoType, error) {
	req := VerisignEPPWrapperWithDefaults()
	req.Command.Info = &xmltypes.InfoCommand{
		DomainInfo: &xmltypes.DomainInfo{
			Name: &xmltypes.DomainName{
				Text:  domainName,
				Hosts: "all",
			},
		},
	}

	response, err := sendEPPRequest(client, req)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	var eppResponse xmltypes.EPPDomainInfoResponse
	err = xml.Unmarshal([]byte(response), &eppResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	if eppResponse.Response.Result.Code != "1000" {
		return nil, fmt.Errorf("%s", err)
	}
	infData := eppResponse.Response.ResData.InfData
	var resp *DomainInfoType
	resp, err = ConvertToDomainInfo(infData)
	return resp, nil
}

func ConvertToDomainInfo(infData *xmltypes.InfData) (*DomainInfoType, error) {
	// Parse dates to time.Time
	crDate, err := time.Parse(time.RFC3339, infData.CrDate)
	if err != nil {
		return &DomainInfoType{}, err
	}

	upDate, err := time.Parse(time.RFC3339, infData.UpDate)
	if err != nil {
		return &DomainInfoType{}, err
	}

	exDate, err := time.Parse(time.RFC3339, infData.ExDate)
	if err != nil {
		return &DomainInfoType{}, err
	}

	// Map statuses
	var statuses []Status
	for _, s := range infData.Status {
		statuses = append(statuses, Status{Status: s.S})
	}

	// Map to DomainInfo struct
	domainInfo := DomainInfoType{
		Name:             infData.Name,
		RegistryObjectID: infData.Roid,
		Statuses:         statuses,
		NameServers:      infData.Ns.HostObj,
		Hosts:            infData.Host,
		ClientID:         infData.ClID,
		CreatorID:        infData.CrID,
		CreationDate:     crDate,
		UpdaterID:        infData.UpID,
		UpdateDate:       upDate,
		ExpirationDate:   exDate,
		AuthInfo: AuthInfo{
			Password: infData.AuthInfo.Pw,
		},
	}

	return &domainInfo, nil
}
