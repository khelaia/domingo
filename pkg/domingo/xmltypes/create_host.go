package xmltypes

import "encoding/xml"

type CreateHostAddr struct {
	Text string `xml:",chardata"`
	Ip   string `xml:"ip,attr"`
}
type CreateHost struct {
	Text string          `xml:",chardata"`
	Name string          `xml:"host:name"`
	Addr *CreateHostAddr `xml:"host:addr"`
}

type EPPCreateHostResponse struct {
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
				Host   string `xml:"xmlns:host,attr"`
				Name   string `xml:"name"`
				CrDate string `xml:"crDate"`
			} `xml:"creData"`
		} `xml:"resData"`
		TrID struct {
			Text   string `xml:",chardata"`
			ClTRID string `xml:"clTRID"`
			SvTRID string `xml:"svTRID"`
		} `xml:"trID"`
	} `xml:"response"`
}
