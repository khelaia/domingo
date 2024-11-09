package xmltypes

import "encoding/xml"

type RegisterDomainPeriod struct {
	Text string `xml:",chardata"`
	Unit string `xml:"unit,attr"`
}

type RegisterDomainAuthInfo struct {
	Text string `xml:",chardata"`
	Pw   string `xml:"domain:pw"`
}
type RegisterDomainStruct struct {
	Text     string                  `xml:",chardata"`
	Name     string                  `xml:"domain:name"`
	Period   *RegisterDomainPeriod   `xml:"domain:period"`
	AuthInfo *RegisterDomainAuthInfo `xml:"domain:authInfo"`
}
type RegisterCommand struct {
	Text   string                `xml:",chardata"`
	Create *RegisterDomainStruct `xml:"domain:create"`
}

type EPPRegisterDomainResponse struct {
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
			Text    string `xml:",chardata"`
			CreData struct {
				Text   string `xml:",chardata"`
				Domain string `xml:"domain,attr"`
				Name   string `xml:"name"`
				CrDate string `xml:"crDate"`
				ExDate string `xml:"exDate"`
			} `xml:"creData"`
		} `xml:"resData"`
		TrID struct {
			Text   string `xml:",chardata"`
			ClTRID string `xml:"clTRID"`
			SvTRID string `xml:"svTRID"`
		} `xml:"trID"`
	} `xml:"response"`
}
