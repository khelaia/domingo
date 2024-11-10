package xmltypes

import "encoding/xml"

type UpdateDomainNs struct {
	Text    string   `xml:",chardata"`
	HostObj []string `xml:"domain:hostObj"`
}
type AddHostToDomain struct {
	Text string          `xml:",chardata"`
	Ns   *UpdateDomainNs `xml:"domain:ns"`
}
type UpdateDomainHosts struct {
	Text string           `xml:",chardata"`
	Name string           `xml:"domain:name"`
	Add  *AddHostToDomain `xml:"domain:add"`
}

type EPPAttachNameserversResponse struct {
	XMLName  xml.Name `xml:"epp"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	Response struct {
		Text   string `xml:",chardata"`
		Result struct {
			Text     string `xml:",chardata"`
			Code     string `xml:"code,attr"`
			Msg      string `xml:"msg"`
			ExtValue struct {
				Reason string `xml:"reason"`
			} `xml:"extValue,omitempty"`
		} `xml:"result"`
		TrID struct {
			Text   string `xml:",chardata"`
			ClTRID string `xml:"clTRID"`
			SvTRID string `xml:"svTRID"`
		} `xml:"trID"`
	} `xml:"response"`
}
