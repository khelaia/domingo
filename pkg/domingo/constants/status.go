package constants

type ClientStatus string

const (
	StatusClientDeleteProhibited   ClientStatus = "clientDeleteProhibited"
	StatusClientUpdateProhibited   ClientStatus = "clientUpdateProhibited"
	StatusClientRenewProhibited    ClientStatus = "clientRenewProhibited"
	StatusClientTransferProhibited ClientStatus = "clientTransferProhibited"
	StatusClientHold               ClientStatus = "clientHold"
)
