package xmltypes

import "encoding/xml"

// LoginCommand represents the login command in an EPP request
type LoginCommand struct {
	ClientID string  `xml:"clID"`
	Password string  `xml:"pw"`
	Options  Options `xml:"options"`
	Svcs     Svcs    `xml:"svcs"`
}

type EPPLoginRequest struct {
	XMLName xml.Name `xml:"epp"`
	Xmlns   string   `xml:"xmlns,attr"`
	Command struct {
		Login struct {
			ClientID string `xml:"clID"`
			Password string `xml:"pw"`
			Options  struct {
				Version string `xml:"version"`
				Lang    string `xml:"lang"`
			} `xml:"options"`
			Svcs struct {
				ObjURI []string `xml:"objURI"`
			} `xml:"svcs"`
		} `xml:"login"`
		ClTRID string `xml:"clTRID"`
	} `xml:"command"`
}

type EppLoginFullResponse struct {
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
