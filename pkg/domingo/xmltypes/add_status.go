package xmltypes

import (
	"encoding/xml"
	"github.com/khelaia/domingo/pkg/domingo/constants"
)

type AddStatus struct {
	Text string                 `xml:",chardata"`
	S    constants.ClientStatus `xml:"s,attr"`
}

type EPPAddStatusResponse struct {
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
