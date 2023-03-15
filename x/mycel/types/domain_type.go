package types

type DomainType int

const (
	TLD DomainType = iota
	SLD
	SubDomain
)
