package xmltypes

import "encoding/xml"

type CheckCommand struct {
	XMLName     xml.Name        `xml:"check"`
	DomainCheck DomainCheckType `xml:"domain:check"`
	Extension   *Extension      `xml:"extension,omitempty"`
}

type DomainCheckType struct {
	XMLName xml.Name `xml:"domain:check"`
	Names   []string `xml:"domain:name"`
}

type DomainCheck struct {
	XMLName     xml.Name        `xml:"check"`
	DomainCheck DomainCheckType `xml:"domain:check"`
	Extension   *Extension      `xml:"extension,omitempty"`
}

type EPPCheckResponse struct {
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
			ChkData struct {
				Text   string `xml:",chardata"`
				Domain string `xml:"domain,attr"`
				Cd     struct {
					Text string `xml:",chardata"`
					Name struct {
						Text  string `xml:",chardata"`
						Avail string `xml:"avail,attr"`
					} `xml:"name"`
					Reason string `xml:"reason"`
				} `xml:"cd"`
			} `xml:"chkData"`
		} `xml:"resData"`
		TrID struct {
			Text   string `xml:",chardata"`
			ClTRID string `xml:"clTRID"`
			SvTRID string `xml:"svTRID"`
		} `xml:"trID"`
	} `xml:"response"`
}
