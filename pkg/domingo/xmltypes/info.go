package xmltypes

import "encoding/xml"

type DomainName struct {
	Text  string `xml:",chardata"`
	Hosts string `xml:"hosts,attr"`
}
type DomainInfo struct {
	Text string      `xml:",chardata"`
	Name *DomainName `xml:"domain:name"`
}
type InfoCommand struct {
	Text       string      `xml:",chardata"`
	DomainInfo *DomainInfo `xml:"domain:info"`
}

type EPPDomainInfoResponse struct {
	XMLName  xml.Name `xml:"epp"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	Response struct {
		Text   string `xml:",chardata"`
		Result struct {
			Text string `xml:",chardata"`
			Code string `xml:"code,attr"`
			Msg  string `xml:"msg"`
		} `xml:"result"`
		ResData struct {
			Text    string   `xml:",chardata"`
			InfData *InfData `xml:"infData"`
		} `xml:"resData"`
		Extension struct {
			Text    string `xml:",chardata"`
			InfData struct {
				Text      string `xml:",chardata"`
				Rgp       string `xml:"rgp,attr"`
				RgpStatus struct {
					Text string `xml:",chardata"`
					S    string `xml:"s,attr"`
				} `xml:"rgpStatus"`
			} `xml:"infData"`
		} `xml:"extension"`
		TrID struct {
			Text   string `xml:",chardata"`
			ClTRID string `xml:"clTRID"`
			SvTRID string `xml:"svTRID"`
		} `xml:"trID"`
	} `xml:"response"`
}

type InfData struct {
	Text   string `xml:",chardata"`
	Domain string `xml:"domain,attr"`
	Name   string `xml:"name"`
	Roid   string `xml:"roid"`
	Status []struct {
		Text string `xml:",chardata"`
		S    string `xml:"s,attr"`
	} `xml:"status"`
	Ns struct {
		Text    string   `xml:",chardata"`
		HostObj []string `xml:"hostObj"`
	} `xml:"ns"`
	Host     []string `xml:"host"`
	ClID     string   `xml:"clID"`
	CrID     string   `xml:"crID"`
	CrDate   string   `xml:"crDate"`
	UpID     string   `xml:"upID"`
	UpDate   string   `xml:"upDate"`
	ExDate   string   `xml:"exDate"`
	AuthInfo struct {
		Text string `xml:",chardata"`
		Pw   string `xml:"pw"`
	} `xml:"authInfo"`
}
