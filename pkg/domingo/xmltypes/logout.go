package xmltypes

import "encoding/xml"

type EPPLogoutResponse struct {
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
		TrID struct {
			Text   string `xml:",chardata"`
			ClTRID string `xml:"clTRID"`
			SvTRID string `xml:"svTRID"`
		} `xml:"trID"`
	} `xml:"response"`
}
