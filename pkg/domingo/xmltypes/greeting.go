package xmltypes

import "encoding/xml"

type EPPGreeting struct {
	XMLName  xml.Name `xml:"epp"`
	Greeting struct {
		SvID    string  `xml:"svID"`
		SvDate  string  `xml:"svDate"`
		SvcMenu SvcMenu `xml:"svcMenu"`
		DCP     DCP     `xml:"dcp"`
	} `xml:"greeting"`
}

type SvcMenu struct {
	Version      string       `xml:"version"`
	Lang         string       `xml:"lang"`
	ObjURI       []string     `xml:"objURI"`
	SvcExtension SvcExtension `xml:"svcExtension"`
}

type SvcExtension struct {
	ExtURI []string `xml:"extURI"`
}

type DCP struct {
	Access    Access    `xml:"access"`
	Statement Statement `xml:"statement"`
}

type Access struct {
	All string `xml:"all"`
}

type Statement struct {
	Purpose   Purpose   `xml:"purpose"`
	Recipient Recipient `xml:"recipient"`
	Retention Retention `xml:"retention"`
}

type Purpose struct {
	Admin string `xml:"admin"`
	Other string `xml:"other"`
	Prov  string `xml:"prov"`
}

type Recipient struct {
	Ours      string `xml:"ours"`
	Public    string `xml:"public"`
	Unrelated string `xml:"unrelated"`
}

type Retention struct {
	Indefinite string `xml:"indefinite"`
}
