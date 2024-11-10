package xmltypes

import "encoding/xml"

// EPPWrapper represents the root structure for EPP requests and responses
type EPPWrapper struct {
	XMLName               xml.Name `xml:"epp"`
	Xmlns                 string   `xml:"xmlns,attr"`
	XmlnsDomain           string   `xml:"xmlns:domain,attr,omitempty"`
	XmlnsContact          string   `xml:"xmlns:contact,attr,omitempty"`
	XmlnsHost             string   `xml:"xmlns:host,attr,omitempty"`
	XmlnsRegistry         string   `xml:"xmlns:registry,attr,omitempty"`
	XmlnsRGPPoll          string   `xml:"xmlns:rgp-poll,attr,omitempty"`
	XmlnsRGP              string   `xml:"xmlns:rgp,attr,omitempty"`
	XmlnsNamestore        string   `xml:"xmlns:namestoreExt,attr,omitempty"`
	XmlnsVerificationCode string   `xml:"xmlns:verificationCode,attr,omitempty"`
	XmlnsChangePoll       string   `xml:"xmlns:changePoll,attr,omitempty"`
	XmlnsSecDNS           string   `xml:"xmlns:secDNS,attr,omitempty"`
	Command               *Command `xml:"command,omitempty"`
}

// Command represents the command structure in an EPP request
type Command struct {
	Login     *LoginCommand  `xml:"login,omitempty"`
	Check     *CheckCommand  `xml:"check,omitempty"`
	Create    *CreateCommand `xml:"create,omitempty"`
	Update    *UpdateCommand `xml:"update,omitempty"`
	Logout    string         `xml:"logout,omitempty"`
	Extension *Extension     `xml:"extension,omitempty"`
	ClTRID    string         `xml:"clTRID"`
}

type CreateCommand struct {
	Text         string                `xml:",chardata"`
	CreateDomain *RegisterDomainStruct `xml:"domain:create,omitempty"`
	CreateHost   *CreateHost           `xml:"host:create"`
}

type UpdateCommand struct {
	Text              string             `xml:",chardata"`
	UpdateDomainHosts *UpdateDomainHosts `xml:"domain:update"`
}

// Extension represents the common extension data
type Extension struct {
	NamestoreExt *NamestoreExtension `xml:"namestoreExt:namestoreExt"`
}

// NamestoreExtension represents the Verisign-specific namestore extension
type NamestoreExtension struct {
	XMLName           xml.Name `xml:"namestoreExt:namestoreExt"`
	SubProduct        string   `xml:"namestoreExt:subProduct"`
	XmlnsNamestoreExt string   `xml:"xmlns:namestoreExt,attr,omitempty"`
}

type Options struct {
	Version string `xml:"version"`
	Lang    string `xml:"lang"`
}

type Svcs struct {
	ObjURI []string `xml:"objURI"`
}
